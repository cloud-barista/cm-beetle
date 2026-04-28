// priority.go provides image prioritization logic for different CSPs
package compat

import (
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/rs/zerolog/log"
)

// List of keywords for image prioritization
const (
	PriorityStandard   = 0
	PriorityMinimal    = 1
	PriorityK8s        = 2
	PriorityProMinimal = 3
	PriorityPro        = 4
	PriorityTest       = 5
	PriorityMarketplace = 10
)

// GetImagePriority returns a priority value for a given image (lower is better)
// It combines general keyword-based prioritization with CSP-specific rules.
func GetImagePriority(csp string, image cloudmodel.ImageInfo) int {
	cspLower := strings.ToLower(csp)
	dist := strings.ToLower(image.OSDistribution)

	// --- 1. General Prioritization (Keyword-based) ---
	// Replicates existing logic from resource-vm-image.go to ensure backward compatibility
	hasMinimal := strings.Contains(dist, "minimal")
	hasK8s := strings.Contains(dist, "k8s")
	hasPro := strings.Contains(dist, "pro")
	hasTest := strings.Contains(dist, "test")
	hasProMinimal := hasPro && hasMinimal

	priority := PriorityStandard
	if hasTest {
		priority = PriorityTest
	} else if hasProMinimal {
		priority = PriorityProMinimal
	} else if hasPro {
		priority = PriorityPro
	} else if hasK8s {
		priority = PriorityK8s
	} else if hasMinimal {
		priority = PriorityMinimal
	}

	// --- 2. CSP-Specific Prioritization ---
	switch cspLower {
	case "azure":
		// Deprioritize images that likely require Marketplace Plan information
		if isAzureMarketplaceImage(image) {
			log.Trace().Msgf("Azure: Deprioritizing Marketplace image: %s", image.CspImageName)
			// Marketplace images should have even lower priority than "test" images
			if priority < PriorityMarketplace {
				priority = PriorityMarketplace
			}
		}
	}

	return priority
}

// isAzureMarketplaceImage detects if an Azure image likely requires Marketplace Plan info
func isAzureMarketplaceImage(image cloudmodel.ImageInfo) bool {
	// 1. Check for well-known standard publishers (Canonical, RedHat, etc.)
	// Azure URN format: Publisher:Offer:Sku:Version
	name := strings.ToLower(image.CspImageName)
	parts := strings.Split(name, ":")
	if len(parts) >= 1 && parts[0] != "" {
		publisher := parts[0]
		standardPublishers := map[string]bool{
			"canonical":               true,
			"redhat":                  true,
			"suse":                    true,
			"openlogic":               true,
			"debian":                  true,
			"oracle":                  true,
			"microsoftwindowsserver":  true,
			"microsoftsqlserver":      true,
			"almalinux":               true,
			"rockylinux":              true,
			"center-for-internet-security-inc": true, // CIS images are often marketplace but sometimes well-known; we use keywords for CIS too
		}

		if !standardPublishers[publisher] {
			log.Trace().Msgf("Azure: Detected non-standard publisher: %s", publisher)
			return true
		}
	}

	// 2. Check for known Marketplace keywords in names and distribution
	dist := strings.ToLower(image.OSDistribution)
	
	marketplaceKeywords := []string{
		"hardened", "cis-", "checkpoint", "fortinet", "f5-", 
		"barracuda", "bitnami", "fips-", "byos",
	}

	for _, kw := range marketplaceKeywords {
		if strings.Contains(name, kw) || strings.Contains(dist, kw) {
			return true
		}
	}

	// 3. Check for explicit Plan information in details
	for _, kv := range image.Details {
		key := strings.ToLower(kv.Key)
		if strings.Contains(key, "plan") || strings.Contains(key, "purchase") {
			return true
		}
	}

	return false
}
