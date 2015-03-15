package queenattack

import (
	"errors"
	"math"
	"strconv"
)

func CanQueenAttack(white, black string) (attack bool, e error) {
	if white == black {
		return attack, errors.New("!")
	}

	W, err1 := BoardPosition(white)
	B, err2 := BoardPosition(black)
	if err1 != nil || err2 != nil {
		return attack, errors.New("!")
	}

	rank := W[1] == B[1] // 4 in b4
	file := W[0] == B[0] // b in b4
	slope := (float64(W[1]) - float64(B[1])) / (float64(W[0]) - float64(B[0]))

	if rank || file || math.Abs(slope) == 1.00 {
		attack = true
	}

	return attack, nil
}

func BoardPosition(s string) (coords []int, err error) {
	long := int(s[0] - 97)
	lat, _ := strconv.ParseInt(string(s[1]), 10, 0)
	coords = []int{long, int(lat)}

	inbounds := (0 <= long && long <= 7) && (0 <= lat && lat <= 7)
	if !inbounds {
		err = errors.New("!")
	}
	return coords, err
}
