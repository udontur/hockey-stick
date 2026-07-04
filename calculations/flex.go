package calculations

import (
	"math"
)

func GetFlex(weightKg int) int {
	const flexFactor float64 = 1.0523

	var exactFlex float64 = float64(weightKg) * flexFactor
	var roundedFlex int = int(math.Round(exactFlex))
	return roundedFlex
}
