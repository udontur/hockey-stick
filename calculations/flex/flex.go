package flex

import (
	"math"
)

func Calculate(weightKg int) int {
	const flexFactor float64=1.1023

	var exactFlex float64=float64(weightKg)*flexFactor
	var roundedFlex int=int(math.Round(exactFlex))
	return roundedFlex
}
