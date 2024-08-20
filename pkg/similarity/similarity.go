package similarity

import (
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
)

/*
Infra similarity calculation methods.
Methods for comparing similar infrastructure information using text similarity.
When direct comparison is difficult, predefined maps are used to assess similarity.
*/

// Predefined architecture map
var archMap = map[string]string{
	"x86_64":  "amd64",
	"x64":     "amd64",
	"amd64":   "amd64",
	"x86":     "i386",
	"i386":    "i386",
	"i686":    "i386",
	"arm64":   "aarch64",
	"armv8":   "aarch64",
	"aarch64": "aarch64",
	"armv7":   "arm",
	"arm":     "arm",
	"ppc64":   "ppc64",
	"ppc64le": "ppc64le",
	"power8":  "ppc64",
	"power9":  "ppc64",
}

// Accessor to ensure archMap immutability
func GetArch(key string) (string, bool) {
	value, ok := archMap[key]
	return value, ok
}

// Predefined architecture bit map
var archBitMap = map[string]string{
	"64bit": "64",
	"32bit": "32",
}

// Accessor to ensure archBitMap immutability
func GetArchBit(key string) (string, bool) {
	value, ok := archBitMap[key]
	return value, ok
}

// ActivateByThresholdReLU applies a ReLU function that activates if the similarity is greater than a threshold
func ActivateByThresholdReLU(similarity, threshold float64) float64 {
	if similarity > threshold {
		return similarity
	}
	return 0.0
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
			similarity := CalculateSimilarityByMapAndSequenceMatcher(word1, word2)
			if similarity > bestMatch {
				bestMatch = similarity
				bestMatchWord = word2

			}
		}

		activatedValue := ActivateByThresholdReLU(bestMatch, 0.5)
		if activatedValue > 0 {
			log.Trace().Msgf("Best match for '%s': '%s' (similarity: %.2f)", word1, bestMatchWord, bestMatch)
			totalSimilarity += activatedValue
		}
	}

	// Normalize by the number of words
	return totalSimilarity / float64(len(words1))
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

func CalculateSimilarityByMapAndSequenceMatcher(word1, word2 string) float64 {
	// Check if the words are mapped in the predefined architecture map
	if mapped, ok := GetArch(word1); ok && mapped == word2 {
		return 1.0
	}
	if mapped, ok := GetArch(word2); ok && mapped == word1 {
		return 1.0
	}

	// if not mapped, calculate the similarity based on SequenceMatcher
	return SequenceMatcher(word1, word2)
}

// CalculateSimilarityBySequenceMatcher calculates the similarity between two words based on SequenceMatcher
func CalculateSimilarityBySequenceMatcher(word1, word2 string) float64 {
	return SequenceMatcher(word1, word2)
}

/*
Text similarity calculation methods.
Methods for measuring how similar two pieces of text are.
Useful for tasks like matching, searching, and deduplication.
*/

// SequenceMatcher similarity by finding the longest substring that two texts have in common.
// It has relatively high complexity. Comparing large strings can take a lot of time.
// It is not suitable for measuring similarity of non-text data (e.g. numbers).
func SequenceMatcher(text1, text2 string) float64 {
	lcs := longestCommonSubstring(text1, text2)
	return 2.0 * float64(len(lcs)) / float64(len(text1)+len(text2))
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

// CalculateSimilarityByLevenshteinDistance calculates the similarity between two words based on Levenshtein distance
func CalculateSimilarityByLevenshteinDistance(word1, word2 string) float64 {
	maxLen := float64(max(len(word1), len(word2)))
	if maxLen == 0 {
		return 1.0
	}
	return 1.0 - float64(LevenshteinDistance(word1, word2))/maxLen
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// LevenshteinDistance measures the similarity between two texts based on
// the minimum number of times one text (s1) is converted to another text (s2).
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

// JaccardSimilarity measures similarity using the size of the intersection and union between two text sets.
// Changing the text order does not affect similarity measurements.
// The result is the same even if the order is changed.
func JaccardSimilarity(text1, delimiter1, text2, delimiter2 string) float64 {

	// Convert a string into a set of words (e.g., "hello world" -> {"hello", "world"})
	setA := toSet(text1, delimiter1)
	setB := toSet(text2, delimiter2)

	// Calculate the Jaccard similarity
	intersectionSize := len(intersection(setA, setB))
	unionSize := len(union(setA, setB))

	if unionSize == 0 {
		return 0
	}

	return float64(intersectionSize) / float64(unionSize)
}

func intersection(setA, setB map[string]struct{}) map[string]struct{} {
	intersection := make(map[string]struct{})
	for item := range setA {
		if _, found := setB[item]; found {
			intersection[item] = struct{}{}
		}
	}
	return intersection
}

func union(setA, setB map[string]struct{}) map[string]struct{} {
	union := make(map[string]struct{})
	for item := range setA {
		union[item] = struct{}{}
	}
	for item := range setB {
		union[item] = struct{}{}
	}
	return union
}

func toSet(text, delimiter string) map[string]struct{} {

	if delimiter == "" {
		log.Warn().Msg("delimiter is empty. Set it to a space (' ')")
		delimiter = " "
	}

	// Convert to lowercase
	text = strings.ToLower(text)

	// Split text by delimiter
	arr := strings.Split(text, delimiter)

	set := make(map[string]struct{})
	for _, item := range arr {
		set[item] = struct{}{}
	}
	return set
}
