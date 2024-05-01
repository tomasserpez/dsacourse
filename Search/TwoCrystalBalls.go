package Search

import (
	"math"
)

func TwiCrystalBalls(breaks []bool) int {
	temp := float64(len(breaks))
	jmpAmount := int(math.Floor(math.Sqrt(temp)))

	i := jmpAmount
	for ; i < len(breaks); i += jmpAmount {
		if breaks[i] {
			break
		}
	}

	i -= jmpAmount
	for j := 0; j < jmpAmount && i < len(breaks); j, i = j+1, i+1 {

	}
}
