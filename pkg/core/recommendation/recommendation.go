package recommendation

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/cloud-barista/cb-tumblebug/src/core/mcir"
	"github.com/cloud-barista/cb-tumblebug/src/core/mcis"
	cloudmodel "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/cloud/infra"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model/onprem/infra"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Recommend(srcInfra []infra.Infra) (cloudmodel.InfraMigrationReq, error) {

	// Initialize resty client with basic auth
	client := resty.New()
	apiUser := viper.GetString("api.username")
	apiPass := viper.GetString("api.password")
	client.SetBasicAuth(apiUser, apiPass)

	// set endpoint
	epTumblebug := common.TumblebugRestUrl

	// check readyz
	method := "GET"
	url := fmt.Sprintf("%s/readyz", epTumblebug)
	reqReadyz := common.NoBody
	resReadyz := new(common.SimpleMsg)

	err := common.ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		common.SetUseBody(reqReadyz),
		&reqReadyz,
		resReadyz,
		common.VeryShortDuration,
	)

	if err != nil {
		log.Err(err).Msg("")
		return cloudmodel.InfraMigrationReq{}, err
	}
	log.Debug().Msgf("resReadyz: %+v", resReadyz.Message)

	// Set a deployment plan to recommand virtual machines
	// Ref: https://github.com/cloud-barista/cb-tumblebug/discussions/1234
	planDocstring := `{
	"filter": {
		"policy": [
			{
				"condition": [
					{
						"operand": "%d",
						"operator": ">="
					},
					{
						"operand": "%d",
						"operator": "<="
					}
				],
				"metric": "vCPU"
			},
			{
				"condition": [
					{
						"operand": "%d",
						"operator": ">="
					},
					{
						"operand": "%d",
						"operator": "<="
					}
				],
				"metric": "memoryGiB"
			},
			{
				"condition": [
					{
						"operand": "%s"
					}
				],
				"metric": "providerName"
			},
			{
				"condition": [
					{
						"operand": "%s"
					}
				],
				"metric": "regionName"
			}
		]
	},
	"limit": "5",
	"priority": {
		"policy": [
			{
				"metric": "performance"
			}
		]
	}
}`

	// A target infrastructure by recommendation
	targetInfra := cloudmodel.InfraMigrationReq{
		Description:     "A cloud infra recommended by CM-Beetle",
		InstallMonAgent: "no",
		Label:           "rehosted-infra",
		Name:            "",
		SystemLabel:     "",
		Vm:              []cloudmodel.HostMigrationReq{},
	}

	// Recommand VMs
	for _, server := range srcInfra {

		// Extract server info from source computing infra info
		cores := server.Compute.ComputeResource.CPU.Cores
		memory := MBtoGiB(float64(server.Compute.ComputeResource.Memory.Size))

		coreUpperLimit := cores << 1
		var coreLowerLimit uint
		if cores > 1 {
			coreLowerLimit = cores >> 1
		} else {
			coreLowerLimit = 1
		}

		memoryUpperLimit := memory << 1
		var memoryLowerLimit uint32
		if memory > 1 {
			memoryLowerLimit = memory >> 1
		} else {
			memoryLowerLimit = 1
		}

		providerName := "aws"
		regionName := "ap-northeast-2"

		osVendor := server.Compute.OS.OS.Vendor
		osVersion := server.Compute.OS.OS.Release
		osNameWithVersion := strings.ToLower(osVendor + osVersion)

		log.Debug().
			Uint("coreUpperLimit", coreUpperLimit).
			Uint("coreLowerLimit", coreLowerLimit).
			Uint32("memoryUpperLimit (GiB)", memoryUpperLimit).
			Uint32("memoryLowerLimit (GiB)", memoryLowerLimit).
			Str("providerName", providerName).
			Str("regionName", regionName).
			Str("osNameWithVersion", osNameWithVersion).
			Msg("Source computing infrastructure info")

		// To search proper VMs with the server info, set a deployment plan
		planToSearchProperVm := fmt.Sprintf(planDocstring,
			coreLowerLimit,
			coreUpperLimit,
			memoryLowerLimit,
			memoryUpperLimit,
			providerName,
			regionName,
		)

		////////////////////////////////////////
		// Search and set a target VM spec
		method := "POST"
		url := fmt.Sprintf("%s/mcisRecommendVm", epTumblebug)

		// Request body
		reqRecommVm := new(mcis.DeploymentPlan)
		err := json.Unmarshal([]byte(planToSearchProperVm), reqRecommVm)
		if err != nil {
			log.Err(err).Msg("")
			return cloudmodel.InfraMigrationReq{}, err
		}
		log.Trace().Msgf("deployment plan for the VM recommendation: %+v", reqRecommVm)

		// Response body
		resRecommVmList := []mcir.TbSpecInfo{}

		err = common.ExecuteHttpRequest(
			client,
			method,
			url,
			nil,
			common.SetUseBody(*reqRecommVm),
			reqRecommVm,
			&resRecommVmList,
			common.VeryShortDuration,
		)

		if err != nil {
			log.Err(err).Msg("")
			return cloudmodel.InfraMigrationReq{}, err
		}

		numRecommenedVm := len(resRecommVmList)

		log.Debug().Msgf("the number of recommended VM specs: %d (for the inserted PM/VM with spec (cores: %d, memory (GiB): %d))", numRecommenedVm, cores, memory)
		log.Trace().Msgf("recommendedVmList for the inserted PM/VM with spec (cores: %d, memory (GiB): %d): %+v", cores, memory, resRecommVmList)

		if numRecommenedVm == 0 {
			log.Warn().Msgf("no VM spec recommended for the inserted PM/VM with spec (cores: %d, memory (GiB): %d)", cores, memory)
			continue
		}
		log.Debug().Msgf("select the 1st recommended virtual machine: %+v", resRecommVmList[0])
		recommendedSpec := resRecommVmList[0].Id

		// name := fmt.Sprintf("rehosted-%s-%s", server.Compute.OS.Node.Hostname, server.Compute.OS.Node.Machineid)
		name := fmt.Sprintf("rehosted-%s", server.Compute.OS.Node.Hostname)

		////////////////////////////////////////
		// Search and set target VM image (e.g. ubuntu22.04)
		method = "POST"
		url = fmt.Sprintf("%s/mcisDynamicCheckRequest", epTumblebug)

		// Request body
		reqMcisDynamicCheck := new(mcis.McisConnectionConfigCandidatesReq)
		reqMcisDynamicCheck.CommonSpecs = []string{recommendedSpec}

		// Response body
		resMcisDynamicCheck := new(mcis.CheckMcisDynamicReqInfo)

		err = common.ExecuteHttpRequest(
			client,
			method,
			url,
			nil,
			common.SetUseBody(*reqMcisDynamicCheck),
			reqMcisDynamicCheck,
			resMcisDynamicCheck,
			common.VeryShortDuration,
		)

		if err != nil {
			log.Err(err).Msg("")
			return cloudmodel.InfraMigrationReq{}, err
		}

		log.Trace().Msgf("resMcisDynamicCheck: %+v", resMcisDynamicCheck)

		if len(resMcisDynamicCheck.ReqCheck) == 0 {
			log.Warn().Msg("no VM OS image recommended for the inserted PM/VM")
			continue
		}

		keywords := fmt.Sprintf("%s %s %s %s",
			server.Compute.OS.OS.Vendor,
			server.Compute.OS.OS.Version,
			server.Compute.OS.OS.Architecture,
			server.Compute.ComputeResource.RootDisk.Type)
		log.Debug().Msg("keywords for the VM OS image recommendation: " + keywords)

		// Select VM OS image via LevenshteinDistance-based text similarity
		delimiters1 := []string{" ", "-", "_", ",", "(", ")", "[", "]", "/"}
		delimiters2 := delimiters1
		vmOsImageId := FindBestVmOsImage(keywords, delimiters1, resMcisDynamicCheck.ReqCheck[0].Image, delimiters2)

		// vmOsImage := fmt.Sprintf("%s+%s+%s", providerName, regionName, osNameWithVersion)

		vm := cloudmodel.HostMigrationReq{
			ConnectionName: "",
			CommonImage:    vmOsImageId,
			CommonSpec:     recommendedSpec,
			Description:    "a recommended virtual machine",
			Label:          "rehosted-vm",
			Name:           name,
			RootDiskSize:   "",
			RootDiskType:   "",
			SubGroupSize:   "",
			VmUserPassword: "",
		}

		targetInfra.Vm = append(targetInfra.Vm, vm)
	}

	log.Trace().Msgf("targetInfra: %+v", targetInfra)

	return targetInfra, nil
}

func MBtoGiB(mb float64) uint32 {
	const bytesInMB = 1000000.0
	const bytesInGiB = 1073741824.0
	gib := (mb * bytesInMB) / bytesInGiB
	return uint32(gib)
}

// FindBestVmOsImage finds the best matching image based on the similarity scores
func FindBestVmOsImage(keywords string, kwDelimiters []string, vmImages []mcir.TbImageInfo, imgDelimiters []string) string {

	var bestVmOsImageID string
	var highestScore float64

	for _, image := range vmImages {
		score := CalculateSimilarity(keywords, kwDelimiters, image.CspImageName, imgDelimiters)
		if score > highestScore {
			highestScore = score
			bestVmOsImageID = image.Id
		}
		log.Trace().Msgf("VmImageName: %s, score: %f", image.CspImageName, score)

	}
	log.Debug().Msgf("bestVmOsImageID: %s, highestScore: %f", bestVmOsImageID, highestScore)

	return bestVmOsImageID
}

// CalculateSimilarity calculates the similarity between two texts based on word similarities
func CalculateSimilarity(text1 string, delimiters1 []string, text2 string, delimiters2 []string) float64 {

	words1 := splitToArray(text1, delimiters1)
	words2 := splitToArray(text2, delimiters2)

	log.Trace().Msgf("From text 1: %s", text1)
	log.Trace().Msgf("To word array 1: %v", words1)
	log.Trace().Msgf("From text 2: %s", text2)
	log.Trace().Msgf("To word array 2: %v", words2)

	// Calculate the similarity between two texts based on word similarities
	totalSimilarity := 0.0
	for _, word1 := range words1 {
		bestMatch := 0.0
		bestMatchWord := ""
		for _, word2 := range words2 {
			// similarity := CalculateSimilarityByLevenshteinDistance(word1, word2)
			similarity := CalculateSimilarityBySequenceMatcher(word1, word2)
			if similarity > bestMatch {
				bestMatch = similarity
				bestMatchWord = word2

			}
		}
		log.Trace().Msgf("Best match for '%s': '%s' (similarity: %.2f)", word1, bestMatchWord, bestMatch)
		totalSimilarity += activateByReLU(bestMatch, 0.5)
	}

	// Normalize by the number of words
	return totalSimilarity // / float64(len(words1))
}

func splitToArray(text string, delimiters []string) []string {

	if len(delimiters) == 0 {
		log.Warn().Msg("warning: delimiters empty. delimiters are empty. Using space (' ') as default delimiter.")
		delimiters = []string{" "}
	}

	// Convert to lowercase
	text = strings.ToLower(text)

	// Create a regular expression pattern for the delimiters
	escapedDelimiters := make([]string, len(delimiters))
	for i, d := range delimiters {
		escapedDelimiters[i] = regexp.QuoteMeta(d)
	}
	pattern := strings.Join(escapedDelimiters, "|")
	re := regexp.MustCompile(pattern)

	// Split text by the delimiters
	arr := re.Split(text, -1)

	// Remove empty strings resulting from the split
	result := []string{}
	for _, str := range arr {
		if str != "" {
			result = append(result, str)
		}
	}

	return result
}

// CalculateSimilarityByLevenshteinDistance calculates the similarity between two words based on Levenshtein distance
func CalculateSimilarityByLevenshteinDistance(word1, word2 string) float64 {
	maxLen := float64(max(len(word1), len(word2)))
	if maxLen == 0 {
		return 1.0
	}
	return 1.0 - float64(LevenshteinDistance(word1, word2))/maxLen
}

// CalculateSimilarityBySequenceMatcher calculates the similarity between two words based on Levenshtein distance
func CalculateSimilarityBySequenceMatcher(word1, word2 string) float64 {
	return SequenceMatcher(word1, word2)
}

// activateByReLU applies a ReLU function that activates if the similarity is greater than a threshold
func activateByReLU(similarity, threshold float64) float64 {
	if similarity > threshold {
		return similarity
	}
	return 0.0
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// LevenshteinDistance calculates the Levenshtein distance between two strings
func LevenshteinDistance(text1, text2 string) int {
	text1Len, text2Len := len(text1), len(text2)
	if text1Len == 0 {
		return text2Len
	}
	if text2Len == 0 {
		return text1Len
	}
	matrix := make([][]int, text1Len+1)
	for i := range matrix {
		matrix[i] = make([]int, text2Len+1)
	}
	for i := 0; i <= text1Len; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= text2Len; j++ {
		matrix[0][j] = j
	}
	for i := 1; i <= text1Len; i++ {
		for j := 1; j <= text2Len; j++ {
			cost := 0
			if text1[i-1] != text2[j-1] {
				cost = 1
			}
			matrix[i][j] = min(matrix[i-1][j]+1, min(matrix[i][j-1]+1, matrix[i-1][j-1]+cost))
		}
	}
	return matrix[text1Len][text2Len]
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// longestCommonSubstring finds the longest common substring between two strings.
func longestCommonSubstring(s1, s2 string) string {
	l1, l2 := len(s1), len(s2)
	matrix := make([][]int, l1+1)
	for i := range matrix {
		matrix[i] = make([]int, l2+1)
	}

	longest := 0
	endIndex := l1
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if s1[i-1] == s2[j-1] {
				matrix[i][j] = matrix[i-1][j-1] + 1
				if matrix[i][j] > longest {
					longest = matrix[i][j]
					endIndex = i
				}
			}
		}
	}

	return s1[endIndex-longest : endIndex]
}

// SequenceMatcher calculates the similarity ratio between two strings.
func SequenceMatcher(text1, text2 string) float64 {
	lcs := longestCommonSubstring(text1, text2)
	return 2.0 * float64(len(lcs)) / float64(len(text1)+len(text2))
}

// // JaccardSimilarity calculates the Jaccard similarity between two strings
// func JaccardSimilarity(text1, delimiter1, text2, delimiter2 string) float64 {

// 	// Convert a string into a set of words (e.g., "hello world" -> {"hello", "world"})
// 	setA := toSet(text1, delimiter1)
// 	setB := toSet(text2, delimiter2)

// 	// Calculate the Jaccard similarity
// 	intersectionSize := len(intersection(setA, setB))
// 	unionSize := len(union(setA, setB))

// 	if unionSize == 0 {
// 		return 0
// 	}

// 	return float64(intersectionSize) / float64(unionSize)
// }

// func intersection(setA, setB map[string]struct{}) map[string]struct{} {
// 	intersection := make(map[string]struct{})
// 	for item := range setA {
// 		if _, found := setB[item]; found {
// 			intersection[item] = struct{}{}
// 		}
// 	}
// 	return intersection
// }

// func union(setA, setB map[string]struct{}) map[string]struct{} {
// 	union := make(map[string]struct{})
// 	for item := range setA {
// 		union[item] = struct{}{}
// 	}
// 	for item := range setB {
// 		union[item] = struct{}{}
// 	}
// 	return union
// }

// func toSet(text, delimiter string) map[string]struct{} {

// 	if delimiter == "" {
// 		log.Warn().Msg("delimiter is empty. Set it to a space (' ')")
// 		delimiter = " "
// 	}

// 	// Convert to lowercase
// 	text = strings.ToLower(text)

// 	// Split text by delimiter
// 	arr := strings.Split(text, delimiter)

// 	set := make(map[string]struct{})
// 	for _, item := range arr {
// 		set[item] = struct{}{}
// 	}
// 	return set
// }
