package recommendation

import (
	"fmt"
	"sort"
	"strings"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/modelconv"
	"github.com/cloud-barista/cm-beetle/pkg/similarity"
	cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-model/infra/on-premise-model"
	"github.com/rs/zerolog/log"
)

// Recommendation limits constants for VM images
const (
	defaultImagesLimit = 20
)

// GetDefaultImagesLimit returns the default VM images recommendation limit
func GetDefaultImagesLimit() int {
	return defaultImagesLimit
}

// RecommendVmOsImage recommends an appropriate VM OS image (e.g., Ubuntu 22.04) for the given VM spec
func RecommendVmOsImage(csp string, region string, server onpremmodel.ServerProperty) (cloudmodel.ImageInfo, error) {

	var emptyRes cloudmodel.ImageInfo

	imageList, err := RecommendVmOsImages(csp, region, server, 20)
	if err != nil {
		log.Error().Err(err).Msg("Failed to recommend VM OS images")
		return emptyRes, err
	}

	// Set keywords and delimiters to calculate text similarity
	keywords, kwDelimiters, imgDelimiters := setKeywordsAndDelimeters(server)
	log.Debug().Msg("keywords for the VM OS image recommendation: " + keywords)

	// Find the best VM OS image
	bestVmOsImage := FindBestVmOsImage(keywords, kwDelimiters, imageList, imgDelimiters)

	log.Debug().Msgf("Best VM OS image found: %+v", bestVmOsImage)

	return bestVmOsImage, nil
}

// RecommendVmOsImageId recommends an appropriate VM OS image (e.g., Ubuntu 22.04) for the given VM spec
func RecommendVmOsImageId(csp string, region string, server onpremmodel.ServerProperty) (string, error) {

	imageList, err := RecommendVmOsImages(csp, region, server, 20)
	if err != nil {
		log.Error().Err(err).Msg("Failed to recommend VM OS images")
		return "", err
	}

	// Set keywords and delimiters to calculate text similarity
	keywords, kwDelimiters, imgDelimiters := setKeywordsAndDelimeters(server)
	log.Debug().Msg("keywords for the VM OS image recommendation: " + keywords)

	vmOsImageId := FindBestVmOsImageNameUsedInCsp(keywords, kwDelimiters, imageList, imgDelimiters)

	log.Debug().Msgf("Best VM OS image ID found: %s", vmOsImageId)

	return vmOsImageId, nil
}

// RecommendVmOsImages recommends an appropriate VM OS image (e.g., Ubuntu 22.04) for the given VM spec
func RecommendVmOsImages(csp string, region string, server onpremmodel.ServerProperty, limit int) ([]cloudmodel.ImageInfo, error) {

	var emptyRes = []cloudmodel.ImageInfo{}
	var vmOsImageInfoList = []cloudmodel.ImageInfo{}

	if limit <= 0 {
		err := fmt.Errorf("invalid 'limit' value: %d, set default: 5", limit)
		log.Warn().Msgf("%s", err.Error())
		limit = defaultImagesLimit
	}

	// Request body
	falseValue := false
	trueValue := true
	searchImageReq := tbmodel.SearchImageRequest{
		// DetailSearchKeys:       []string{},
		// IncludeDeprecatedImage: &falseValue,
		// IsKubernetesImage:      &falseValue, // The only image in the Azure (ubuntu 22.04) is both for K8s nodes and gerneral VMs.
		// IsRegisteredByAsset:    &falseValue,
		IncludeBasicImageOnly: &trueValue,
		MaxResults:            &limit,
		OSArchitecture:        tbmodel.OSArchitecture(server.CPU.Architecture),
		OSType:                server.OS.Name + " " + server.OS.VersionID,
		ProviderName:          csp,
		RegionName:            region,
	}

	// TODO: Add condition to check if searchImageReq.IsGPUImage is set, when GPU information is confirmed in the source model
	searchImageReq.IsGPUImage = &falseValue

	log.Debug().Msgf("searchImageReq: %+v", searchImageReq)

	// Call Tumblebug API to search VM OS images
	apiConfig := tbclient.ApiConfig{
		RestUrl:  config.Tumblebug.RestUrl,
		Username: config.Tumblebug.API.Username,
		Password: config.Tumblebug.API.Password,
	}
	tbCli := tbclient.NewClient(apiConfig)
	nsId := "system" // default

	resSearchImage, err := tbCli.SearchVmOsImage(nsId, searchImageReq)
	if err != nil {
		log.Error().Err(err).Msg("")
		return emptyRes, err
	}

	// Debug logging up to 3 images to avoid excessive output
	if len(resSearchImage.ImageList) > 3 {
		for i := range 3 {
			log.Debug().Msgf("[Response from Tumblebug] resSearchImage.ImageList[%d]: %+v", i, resSearchImage.ImageList[i])
		}
	} else {
		for i := range resSearchImage.ImageList {
			log.Debug().Msgf("[Response from Tumblebug] resSearchImage.ImageList[%d]: %+v", i, resSearchImage.ImageList[i])
		}
	}

	// Filter VM OS images to support stability
	var filteredImages []tbmodel.ImageInfo
	for _, img := range resSearchImage.ImageList {
		if strings.Contains(strings.ToLower(img.CspImageName), "uefi") {
			continue
		}
		// Add more filters as needed

		filteredImages = append(filteredImages, img)
	}

	// Convert model from '[]tbmodel.ImageInfo' to '[]cloudmodel.ImageInfo'
	imageList, err := modelconv.ConvertWithValidation[[]tbmodel.ImageInfo, []cloudmodel.ImageInfo](filteredImages)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert VM OS image list")
		return emptyRes, err
	}

	// Set keywords and delimiters to calculate text similarity
	keywords, kwDelimiters, imgDelimiters := setKeywordsAndDelimeters(server)
	log.Debug().Msg("keywords for the VM OS image recommendation: " + keywords)

	// Select VM OS image via LevenshteinDistance-based text similarity
	vmOsImageInfoList = FindAndSortVmOsImageInfoListBySimilarity(keywords, kwDelimiters, imageList, imgDelimiters)

	count := len(vmOsImageInfoList)
	if count == 0 {
		err := fmt.Errorf("no VM OS image recommended for the inserted PM/VM")
		log.Warn().Msgf("%s", err.Error())
		return emptyRes, err
	}

	// [Output]
	// Limit the number of VM specs
	if limit < count {
		log.Debug().Msgf("Limiting the number of recommended VM OS images to %d", limit)
		// * Note: If the number of recommended VM OS images is less than the limit, it will not be truncated.
		// * This is to ensure that the user can see all available images.
		vmOsImageInfoList = vmOsImageInfoList[:limit]
	}

	log.Debug().Msgf("Found %d VM OS images for the given server: %s", len(vmOsImageInfoList), server.MachineId)

	return vmOsImageInfoList, nil
}

func setKeywordsAndDelimeters(server onpremmodel.ServerProperty) (string, []string, []string) {
	keywords := fmt.Sprintf("%s %s %s %s %s",
		server.OS.Name,
		server.OS.VersionID,
		server.OS.VersionCodename,
		server.CPU.Architecture,
		server.RootDisk.Type)

	kwDelimiters := []string{" ", "-", ",", "(", ")", "[", "]", "/"}
	imgDelimiters := []string{" ", "-", ",", "(", ")", "[", "]", "/"}

	return keywords, kwDelimiters, imgDelimiters
}

// FindBestVmOsImage finds the best matching image based on the similarity scores
func FindBestVmOsImage(keywords string, kwDelimiters []string, vmImages []cloudmodel.ImageInfo, imgDelimiters []string) cloudmodel.ImageInfo {

	var bestVmOsImage cloudmodel.ImageInfo
	var highestScore float64 = 0.0

	for _, image := range vmImages {

		vmImgKeywords := fmt.Sprintf("%s %s %s %s",
			image.OSType,
			image.OSArchitecture,
			image.OSDiskType,
			image.OSDistribution,
		)

		score := similarity.CalcResourceSimilarity(keywords, kwDelimiters, vmImgKeywords, imgDelimiters)
		if score > highestScore {
			highestScore = score
			bestVmOsImage = image
		}
		// log.Debug().Msgf("VmImageName: %s, score: %f, description: %s", image.OSDistribution, score, image.Description)

	}
	log.Debug().Msgf("highestScore: %f, bestVmOsImage: %v", highestScore, bestVmOsImage)

	return bestVmOsImage
}

type VmOsImageInfoWithScore struct {
	VmOsImageInfo   cloudmodel.ImageInfo
	SimilarityScore float64
}

// FindAndSortVmOsImageInfoListBySimilarity finds VM OS images that match the keywords and sorts them by similarity score
func FindAndSortVmOsImageInfoListBySimilarity(keywords string, kwDelimiters []string, vmImages []cloudmodel.ImageInfo, imgDelimiters []string) []cloudmodel.ImageInfo {

	var imageInfoListForSorting []VmOsImageInfoWithScore
	var imageInfoList []cloudmodel.ImageInfo

	for _, image := range vmImages {

		vmImgKeywords := fmt.Sprintf("%s %s %s %s",
			image.OSType,
			image.OSArchitecture,
			image.OSDiskType,
			image.OSDistribution,
		)

		score := similarity.CalcResourceSimilarity(keywords, kwDelimiters, vmImgKeywords, imgDelimiters)
		imageInfo := VmOsImageInfoWithScore{
			VmOsImageInfo:   image,
			SimilarityScore: score,
		}
		imageInfoListForSorting = append(imageInfoListForSorting, imageInfo)

	}

	// Sort the imageInfoList by highestScore in descending order
	sort.Slice(imageInfoListForSorting, func(i, j int) bool {
		return imageInfoListForSorting[i].SimilarityScore > imageInfoListForSorting[j].SimilarityScore
	})

	// List the sorted images
	for _, imageInfo := range imageInfoListForSorting {
		log.Debug().Msgf("VmImageName: %s, score: %f, description: %s", imageInfo.VmOsImageInfo.Name, imageInfo.SimilarityScore, imageInfo.VmOsImageInfo.Description)
		imageInfoList = append(imageInfoList, imageInfo.VmOsImageInfo)
	}

	return imageInfoList
}

// FindBestVmOsImageNameUsedInCsp finds the best matching image based on the similarity scores
func FindBestVmOsImageNameUsedInCsp(keywords string, kwDelimiters []string, vmImages []cloudmodel.ImageInfo, imgDelimiters []string) string {

	var bestVmOsImageNameUsedInCsp string
	var highestScore float64 = 0.0

	for _, image := range vmImages {
		vmImgKeywords := fmt.Sprintf("%s %s %s %s",
			image.OSType,
			image.OSArchitecture,
			image.OSDiskType,
			image.OSDistribution,
		)

		score := similarity.CalcResourceSimilarity(keywords, kwDelimiters, vmImgKeywords, imgDelimiters)
		if score > highestScore {
			highestScore = score
			bestVmOsImageNameUsedInCsp = image.CspImageName
		}
		// log.Debug().Msgf("VmImageName: %s, score: %f, description: %s", image.OSDistribution, score, image.Description)

	}
	log.Debug().Msgf("bestVmOsImageID: %s, highestScore: %f", bestVmOsImageNameUsedInCsp, highestScore)

	return bestVmOsImageNameUsedInCsp
}
