package main

import (
	"fmt"

	"github.com/cloud-barista/cm-beetle/pkg/similarity"
)

func main() {

	compareWordSet := []struct {
		str1 string
		str2 string
	}{
		{"22.04", "22.04.1"},
		{"22.04", "20.04"},
		{"20.04", "18.04"},
		{"x86_64", "x86_64"},
		{"amd64", "x86_64"},
		{"x64", "x86_64"},
		{"x86", "i386"},
		{"x86", "i686"},
		{"i686", "i386"},
		{"32bit", "i386"},
		{"amd64", "arm64"},
		{"arm64", "arm64"},
		{"aarch64", "arm64"},
		{"armv8", "arm64"},
		{"armv7", "armv7"},
		{"arm", "armv7"},
		{"amd32", "i386"},
		{"hvm-ssd", "ssd"},
		{"hvm-ssd", "hdd"},
	}

	for _, set := range compareWordSet {
		fmt.Printf("Comparing '%s' with '%s':\n", set.str1, set.str2)
		fmt.Printf(" - LevenshteinDistance, Similarity ratio: %.2f\n", similarity.NormalizedLevenshteinDistance(set.str1, set.str2))
		fmt.Printf(" - SequenceMatcher, Similarity ratio: %.2f\n", similarity.NormalizedSequenceMatcher(set.str1, set.str2))
		fmt.Println("--------------------------------------------------------")
	}

	requirementContent := "Ubuntu 22.04.4 LTS (Jammy Jellyfish) x86_64 HDD"

	/*
		Example: [AWS] evaluate similarity between keywords and VM OS images
	*/
	// VM OS images
	awsVmOSImages := []string{
		"ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20220609",
		"Ubuntu_20.04-x86_64-SQL_2022_Standard-2023.10.16",
		"ubuntu/images/hvm-ssd/ubuntu-bionic-18.04-amd64-server-20191002",
	}

	// Evaluate similarity between requirementContent and VM OS images
	awsRows := evaluateSimilarity(requirementContent, awsVmOSImages)

	fmt.Println("[AWS] Requirement Content, Target VM OS Image, Similarity Score")
	for _, row := range awsRows {
		fmt.Println(row)
	}

	/*
		Example: [Azure] evaluate similarity between keywords and VM OS images from
	*/

	// VM OS images
	azureVmOSImages := []string{
		"Debian:debian-10:10:0.20240204.1647",
		"Canonical:0001-com-ubuntu-server-jammy:22_04-lts:22.04.202404090",
		"Canonical:0001-com-ubuntu-server-focal:20_04-lts:20.04.202404080",
	}

	// Evaluate similarity between requirementContent and VM OS images
	azureRows := evaluateSimilarity(requirementContent, azureVmOSImages)

	fmt.Println("[Azure] Requirement Content, Target VM OS Image, Similarity Score")
	for _, row := range azureRows {
		fmt.Println(row)
	}

	/*
		Example: [GCP] evaluate similarity between keywords and VM OS images from
	*/

	// VM OS images
	gcpVmOSImages := []string{
		"debian-10-buster-v20221102",
		"ubuntu-2204-jammy-v20240319",
		"windows-server-2012-r2-dc-v20221014",
		"ubuntu-2004-focal-v20240307b",
		"tf-2-15-cu121-v20240417-debian-11",
	}

	// Evaluate similarity between requirementContent and VM OS images
	gcpRows := evaluateSimilarity(requirementContent, gcpVmOSImages)

	fmt.Println("[GCP] Requirement Content, Target VM OS Image, Similarity Score")
	for _, row := range gcpRows {
		fmt.Println(row)
	}

}

// VM OS 이미지와 키워드를 비교하는 함수
func evaluateSimilarity(requirementContent string, vmOSImages []string) [][]string {
	var rows [][]string
	delimiters1 := []string{" ", "-", "_", ",", "(", ")", "[", "]", "/"}
	delimiters2 := delimiters1

	for _, image := range vmOSImages {
		score := similarity.CalcResourceSimilarity(requirementContent, delimiters1, image, delimiters2)
		rows = append(rows, []string{requirementContent, image, fmt.Sprintf("%.2f", score)})
	}

	return rows
}
