package curve

import (
	"log"
)

func Get(hockeyType string) (curve string){
	switch hockeyType{
		case "Ice", "Roller":
			curve="P92"
		case "Ball":
			curve="P88"
		default:
			log.Fatal("Invalid type of hockey")
	}
	return curve
}
