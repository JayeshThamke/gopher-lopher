package isogram

import (
	"strings"
)

// IsIsogram -
func IsIsogram(phrase string) bool {
	// '-' and ' ' are not required to check if phrase is isogram
	sanitizedPhrase := strings.ReplaceAll(strings.ReplaceAll(phrase, "-", ""), " ", "")
	phraseUpper := strings.ToUpper(sanitizedPhrase)
	if phraseUpper == "" {
		return true
	}
	freqMap := make(map[int32]int32)

	for _, char := range phraseUpper {
		if _, ok := freqMap[char]; ok {
			return false
		}
		freqMap[char] = freqMap[char] + 1
	}

	return true
}
