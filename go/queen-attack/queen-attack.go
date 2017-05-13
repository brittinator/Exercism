package queenattack

import "errors"

const testVersion = 2

// CanQueenAttack returns if the Queens given their locations can attack each other on a chess board
func CanQueenAttack(w, b string) (bool, error) {
	// guard
	if w == b {
		return false, errors.New("cannot have pieces in the same location")
	}
	if checkPositionValidity(w) && checkPositionValidity(b) {
		// on same horizontal
		if w[1] == b[1] {
			return true, nil
		}
		// on same vertical
		if w[0] == b[0] {
			return true, nil
		}
		// on diagonal if differences between letter and digit are the same
		var horizontal uint8
		var vertical uint8
		if w[1] > b[1] {
			horizontal = w[1] - b[1]
		} else {
			horizontal = b[1] - w[1]
		}
		if w[0] > b[0] {
			vertical = w[0] - b[0]
		} else {
			vertical = b[0] - w[0]
		}
		if horizontal == vertical {
			return true, nil
		}
	} else {
		return false, errors.New("piece(s) are not on the board")
	}
	return false, nil
}

func checkPositionValidity(pos string) bool {
	if pos[0] >= 'a' && pos[0] <= 'h' && pos[1] >= '1' && pos[1] <= '8' {
		return true
	}
	return false
}
