package calculations

import (
	"log"
)

func GetKick(position string) (kick string) {
	switch position {
	case "Winger":
		kick = "Low"
	case "Center":
		kick = "Hybrid"
	case "Defenseman":
		kick = "Mid"
	default:
		log.Fatal("Invalid hockey position")
	}
	return kick
}
