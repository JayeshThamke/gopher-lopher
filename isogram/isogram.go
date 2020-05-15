package isogram

import (
	"strings"
)

// IsIsogram returns true if phrase is isogram else false
func IsIsogram(phrase string) bool {
	// '-' and ' ' are not required to check if phrase is isogram
	sanitizedPhrase := strings.ReplaceAll(strings.ReplaceAll(phrase, "-", ""), " ", "")
	// case doesn't matter here 
	phraseUpper := strings.ToUpper(sanitizedPhrase)
	
	if phraseUpper == "" {
		return true
	}
	freqMap := make(map[int32]int32)
	
	// every char in string is represented as int32 val
	for _, char := range phraseUpper {
		if _, ok := freqMap[char]; ok {
			// any char repeates then phrase isn't isogram
			return false
		}
		freqMap[char] = freqMap[char] + 1
	}

	return true
}
