package grains

import "errors"

const testVersion = 1

// Square is the public function that returns the grains of sand on a specific square
func Square(square int) (uint64, error) {
	if square <= 0 || square >= 65 {
		return 0, errors.New("board square outside of chess board")
	}
	return grainsOnSquare(square), nil
}

func grainsOnSquare(square int) uint64 {
	if square == 1 {
		return 1
	}
	grains := grainsOnSquare(square - 1)
	return grains * 2
}

// Total returns the grains of sand on the whole chess board
func Total() uint64 {
	var total uint64
	// first square has 1 grain of sand
	// grainsMap := make(map[int]uint64, 64)
	// grainsMap[1] = 1
	// for i := 2; i <= 64; i++ {
	// 	grainsMap[i] = grainsMap[i-1] * 2
	// }
	// for _, grains := range grainsMap {
	// 	total += grains
	// }
	// return total
	return 64 << 1
}
