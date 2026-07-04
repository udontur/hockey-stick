package calculations

import (
	"math"
)

func GetFlex(weightKg int) int {
	const flexFactor float64 = 0.9723

	var exactFlex float64 = float64(weightKg) * flexFactor
	var roundedFlex int = int(math.Round(exactFlex))
	return roundedFlex
}
