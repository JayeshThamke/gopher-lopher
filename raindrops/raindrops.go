package raindrops

import (
	"strconv"
)

// Convert converts given number to sound of raindrops
func Convert(num int) string {
	primes := []int{3, 5, 7}
	var sound string

	for _, prime := range primes {
		if num%prime != 0 {
			continue
		}

		switch prime {
		case 3:
			sound = sound + "Pling"
		case 5:
			sound = sound + "Plang"
		case 7:
			sound = sound + "Plong"
		}
	}
	if sound == "" {
		sound = strconv.Itoa(num)
	}

	return sound
}
