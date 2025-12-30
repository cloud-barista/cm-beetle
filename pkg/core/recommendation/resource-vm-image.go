package recommendation

import (
	"fmt"
	"sort"
	"strings"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/compat"
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

// RecommendVmOsImagesForSpec recommends an appropriate VM OS images (e.g., Ubuntu 22.04) for the given VM spec
func RecommendVmOsImagesForSpec(csp string, region string, server onpremmodel.ServerProperty, limit int, spec cloudmodel.SpecInfo) ([]cloudmodel.ImageInfo, error) {

	var emptyRes = []cloudmodel.ImageInfo{}

	if limit <= 0 {
		err := fmt.Errorf("invalid 'limit' value: %d, set default: %d", limit, defaultImagesLimit)
		log.Warn().Msgf("%s", err.Error())
		limit = defaultImagesLimit
	}

	// Get all recommended OS images for the server
	imageList, err := RecommendVmOsImages(csp, region, server, limit)
	if err != nil {
		log.Warn().Err(err).Msg("failed to recommend VM OS images")
		return emptyRes, err
	}

	// Filter images that are compatible with the given VM spec
	compatibleImages := make([]cloudmodel.ImageInfo, 0, len(imageList))

	for _, image := range imageList {
		// Check compatibility between spec and image using the existing compatibility checker
		// This follows the same pattern as RecommendVmSpecsForImage but in reverse
		if isCompatible := compat.CheckCompatibility(strings.ToLower(csp), spec, image); isCompatible {
			compatibleImages = append(compatibleImages, image)
		} else {
			log.Debug().Msgf("Filtered incompatible image: %s for spec: %s on CSP: %s",
				image.CspImageName, spec.CspSpecName, csp)
		}
	}

	if len(compatibleImages) == 0 {
		log.Warn().Msgf("No compatible images found for spec %s on CSP %s, returning original list",
			spec.CspSpecName, csp)
		return imageList, nil
	}

	log.Info().Msgf("Filtered %d images to %d compatible images for spec %s on CSP %s",
		len(imageList), len(compatibleImages), spec.CspSpecName, csp)

	return compatibleImages, nil
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
	keywords, kwDelimiters, imgDelimiters := SetKeywordsAndDelimeters(server)
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
	keywords, kwDelimiters, imgDelimiters := SetKeywordsAndDelimeters(server)
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

	// Call Tumblebug API to search VM OS images
	nsId := "system" // default

	var filteredImages []tbmodel.ImageInfo

	// Note - (sample) the extracted OS information in case of Ubuntu 22.04
	// "os": {
	// 	"prettyName": "Ubuntu 22.04.3 LTS",
	// 	"version": "22.04.3 LTS (Jammy Jellyfish)",
	// 	"name": "Ubuntu",
	// 	"versionId": "22.04",
	// 	"versionCodename": "jammy",
	// 	"id": "ubuntu",
	// 	"idLike": "debian"
	// }

	// Note - (sample) the extracted OS information in case of Debian 13
	// "os": {
	// 		"id": "debian",
	// 		"name": "Debian GNU/Linux",
	// 		"prettyName": "Debian GNU/Linux 13 (trixie)",
	// 		"version": "13 (trixie)",
	// 		"versionCodename": "trixie",
	// 		"versionId": "13"
	// }

	// Request body
	falseValue := false
	trueValue := true

	osType := server.OS.ID + " " + server.OS.VersionID
	// if server.OS.ID == "debian" {
	// 	log.Warn().Msg("Tumblebug currently does not support 'versionID' for debian images; using only 'debian' as OSType")
	// 	osType = server.OS.ID // TODO: Check Tumblebug API to append 'versionID' for debian when debian images are available in CSPs
	// }

	// Try first search with OS ID + Version ID
	searchImageReq := tbmodel.SearchImageRequest{
		// DetailSearchKeys:       []string{},
		// IncludeDeprecatedImage: &falseValue,
		// IsRegisteredByAsset:    &falseValue,
		// IsKubernetesImage:      &falseValue, // The only image in the Azure (ubuntu 22.04) is both for K8s nodes and gerneral VMs.
		IncludeBasicImageOnly: &trueValue,
		MaxResults:            &limit,
		OSArchitecture:        tbmodel.OSArchitecture(server.CPU.Architecture),
		OSType:                osType,
		ProviderName:          csp,
		RegionName:            region,
	}

	// TODO: Add condition to check if searchImageReq.IsGPUImage is set, when GPU information is confirmed in the source model
	searchImageReq.IsGPUImage = &falseValue

	log.Debug().Msgf("searchImageReq: %+v", searchImageReq)

	resSearchImage, err := tbclient.NewSession().SearchVmOsImage(nsId, searchImageReq)
	if err != nil {
		log.Error().Err(err).Msg("")
		return emptyRes, err
	}

	// Filter VM OS images to support stability
	for _, img := range resSearchImage.ImageList {
		if strings.Contains(strings.ToLower(img.CspImageName), "uefi") {
			continue
		}
		// Add more filters as needed

		filteredImages = append(filteredImages, img)
	}

	// If no images found after filtering, retry with OS ID only (without version)
	if len(filteredImages) == 0 {
		log.Warn().Msgf("No images found with osType '%s', retrying with OS ID only: '%s'", osType, server.OS.ID)

		// Update search request to use OS ID only
		searchImageReq.OSType = server.OS.ID

		log.Debug().Msgf("Retry searchImageReq: %+v", searchImageReq)

		resSearchImage, err = tbclient.NewSession().SearchVmOsImage(nsId, searchImageReq)
		if err != nil {
			log.Error().Err(err).Msg("Failed to retry image search with OS ID only")
			return emptyRes, err
		}

		// Filter VM OS images again
		filteredImages = []tbmodel.ImageInfo{} // Reset filtered images
		for _, img := range resSearchImage.ImageList {
			if strings.Contains(strings.ToLower(img.CspImageName), "uefi") {
				continue
			}
			// Add more filters as needed

			filteredImages = append(filteredImages, img)
		}

		if len(filteredImages) > 0 {
			log.Info().Msgf("Found %d images after retrying with OS ID only: '%s'", len(filteredImages), server.OS.ID)
		}
	}

	if len(filteredImages) == 0 {
		err := fmt.Errorf("no VM OS images found for the given server even though retrying with OS ID only")
		return emptyRes, err
	}

	// Debug logging up to 3 images to avoid excessive output
	if len(filteredImages) > 3 {
		for i := range 3 {
			log.Debug().Msgf("Searched and filtered images[%d]: %+v", i, filteredImages[i])
		}
	} else {
		for i := range filteredImages {
			log.Debug().Msgf("Searched and filtered images[%d]: %+v", i, filteredImages[i])
		}
	}

	// Convert model from '[]tbmodel.ImageInfo' to '[]cloudmodel.ImageInfo'
	imageList, err := modelconv.ConvertWithValidation[[]tbmodel.ImageInfo, []cloudmodel.ImageInfo](filteredImages)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert VM OS image list")
		return emptyRes, err
	}

	// Set keywords and delimiters to calculate text similarity
	keywords, kwDelimiters, imgDelimiters := SetKeywordsAndDelimeters(server)
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

func SetKeywordsAndDelimeters(server onpremmodel.ServerProperty) (string, []string, []string) {
	keywords := fmt.Sprintf("%s %s %s %s %s",
		server.OS.ID,
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
	// If scores are equal, deprioritize images containing "minimal", "k8s", "pro", or "test" in OSDistribution
	sort.Slice(imageInfoListForSorting, func(i, j int) bool {
		scoreI := imageInfoListForSorting[i].SimilarityScore
		scoreJ := imageInfoListForSorting[j].SimilarityScore

		// Primary sort: by similarity score (descending)
		if scoreI != scoreJ {
			return scoreI > scoreJ
		}

		// Secondary sort: prioritize standard images, then "minimal", then "k8s", then "pro-minimal", then "pro", then "test"
		distI := strings.ToLower(imageInfoListForSorting[i].VmOsImageInfo.OSDistribution)
		distJ := strings.ToLower(imageInfoListForSorting[j].VmOsImageInfo.OSDistribution)

		// TODO: Modify when the kubernetes images are supported normally
		hasMinimalI := strings.Contains(distI, "minimal")
		hasK8sI := strings.Contains(distI, "k8s")
		hasProI := strings.Contains(distI, "pro")
		hasTestI := strings.Contains(distI, "test")
		hasMinimalJ := strings.Contains(distJ, "minimal")
		hasK8sJ := strings.Contains(distJ, "k8s")
		hasProJ := strings.Contains(distJ, "pro")
		hasTestJ := strings.Contains(distJ, "test")

		// Check for pro-minimal (contains both "pro" and "minimal")
		hasProMinimalI := hasProI && hasMinimalI
		hasProMinimalJ := hasProJ && hasMinimalJ

		// Calculate priority: 0 (standard) > 1 (minimal) > 2 (k8s) > 3 (pro-minimal) > 4 (pro) > 5 (test)
		priorityI := 0
		if hasTestI {
			priorityI = 5
		} else if hasProMinimalI {
			priorityI = 3
		} else if hasProI {
			priorityI = 4
		} else if hasK8sI {
			priorityI = 2
		} else if hasMinimalI {
			priorityI = 1
		}

		priorityJ := 0
		if hasTestJ {
			priorityJ = 5
		} else if hasProMinimalJ {
			priorityJ = 3
		} else if hasProJ {
			priorityJ = 4
		} else if hasK8sJ {
			priorityJ = 2
		} else if hasMinimalJ {
			priorityJ = 1
		}

		// Lower priority value comes first
		if priorityI != priorityJ {
			return priorityI < priorityJ
		}

		// Otherwise, maintain original order (stable sort)
		return false
	})

	// List the sorted images
	for _, imageInfo := range imageInfoListForSorting {
		log.Debug().Msgf("VmImageName: %s, score: %f, description: %s, osDist: %s", imageInfo.VmOsImageInfo.Name, imageInfo.SimilarityScore, imageInfo.VmOsImageInfo.Description, imageInfo.VmOsImageInfo.OSDistribution)
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
