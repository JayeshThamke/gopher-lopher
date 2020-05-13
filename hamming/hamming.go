package hamming

import (
	"fmt"
)

// Distance returns hamming distance between two equal length DNA stands
func Distance(a, b string) (int, error) {
	var hamming int
	if len(a) != len(b) {
		hamming = 0
		return hamming, fmt.Errorf("Unequal DNA stands with length %v and %v", len(a), len(b))
	}

	if len(a) == 0 || len(b) == 0 {
		hamming = 0
		return hamming, nil
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hamming++
		}
	}

	return hamming, nil
}
