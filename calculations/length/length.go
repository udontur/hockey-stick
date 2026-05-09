package length

import (
	"math"
)

func Calculate(heightCm int, hockeyType string, position string) int {
	var heightInch float64=float64(heightCm)*0.3937

	var exactLength float64
	switch hockeyType{
		case "Ice", "Roller":
			switch position{
				case "Winger", "Center":
					exactLength=heightInch-8
				case "Defenseman":
					exactLength=heightInch-6
				default:
					exactLength=-1
			}
		case "Ball":
			switch position{
				case "Winger", "Center":
					exactLength=heightInch-12
				case "Defenseman":
					exactLength=heightInch-10
				default:
					exactLength=-1
			}
		default:
			exactLength=-1
	}

	var length int=int(math.Round(exactLength))
	return length
}
