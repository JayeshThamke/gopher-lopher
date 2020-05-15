package luhn

import (
	"strconv"
	"strings"
)

// Valid returns a true if string is valid creditcard according to luhn algorithm else false
func Valid(creditcard string) bool {
	// stripping intermediate, leading and trailing whitespaces, and -
	sanitizedCC := strings.ReplaceAll(strings.TrimSpace(creditcard), " ", "")
	splitCC := strings.Split(sanitizedCC, "")

	// validate for corner case
	if len(splitCC) <= 1 {
		return false
	}

	var intCC = make([]int, len(splitCC))
	var accumulator int
	var secondPosIndicator int

	// for array with even and odd len. the second position from right side
	// is different and array reversal invalidates algorithm
	// hence by looking at the size of array we should first decide how
	// we are going to calculte second pos. from right (end of array)
	if len(splitCC)%2 != 0 {
		secondPosIndicator = 1
		// otherwise default val is 0
	}

	for i := len(splitCC) - 1; i >= 0; i-- {
		// convert str to int32 and modify every second digit from right
		num, err := strconv.Atoi(splitCC[i])

		if err != nil {
			return false
		}
		intCC[i] = num

		if i%2 == secondPosIndicator {
			num *= 2
			if num >= 9 {
				num = num - 9
			}
			intCC[i] = num
		}

		// aggregate sum of array
		accumulator = accumulator + intCC[i]
	}

	if accumulator%10 == 0 {
		return true
	}

	return false
}
