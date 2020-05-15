package scrabble

import "strings"

// Score returns the scabble score of given word.
func Score(word string) int {
	// local cache helps optmizing repeated operations
	cache := make(map[string]int)

	// early return
	if word == "" {
		return 0
	}

	// convert to all UPPERCASE as letter case does not matter
	wordCAPS := strings.ToUpper(word)

	// check first in cache
	for _, byteChar := range wordCAPS {
		// this is golang
		char := string(byteChar)
		if val, ok := cache[char]; ok {
			cache[char] = val + val
			continue
		}

		// if not found in cache then perform lookup
		// this is easy lookup but it may be performance
		// expensive operation
		cache[char] = scoreSingleChar(char)
	}

	// aggregate score
	var scrabbleScore int
	for _, val := range cache {
		scrabbleScore = scrabbleScore + val
	}

	return scrabbleScore
}

func scoreSingleChar(char string) int {
	switch char {
	case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
		return 1
	case "D", "G":
		return 2
	case "B", "C", "M", "P":
		return 3
	case "F", "H", "V", "W", "Y":
		return 4
	case "K":
		return 5
	case "J", "X":
		return 8
	case "Q", "Z":
		return 10
	}
	return 0
}
