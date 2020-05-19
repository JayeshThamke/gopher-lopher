package grains

import (
	"fmt"
)

// Square returns number of grains on the chessboard position
func Square(pos int) (uint64, error) {

	if pos <= 0 || pos > 64 {
		return 0, fmt.Errorf("OutofBound chessboard square position(s) : %v", pos)
	}

	chessBoardSquarePos := pos - 1

	var init uint64 = 1

	// bitwise left shift operator squares the original number (bit)
	// at pos 1 - 0001 = d(1)
	// at pos 2 - 0001 << 1 -> 0010 = d(2)
	// at pos 3 - 0010 << 1 -> 0100 = d(4)
	// at pos 3 - 0100 << 1 -> 1000 = d(8) and so on ...
	return init << chessBoardSquarePos, nil
}

// Total returns total sum of grains on chessboard
func Total() uint64 {
	var accumulator uint64
	for i := 1; i <= 64; i++ {
		sq, _ := Square(i)
		accumulator += sq
	}
	return accumulator
}
