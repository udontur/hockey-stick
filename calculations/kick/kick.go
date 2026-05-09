package kick

import (
	"log"
)

func Get(position string) (kick string) {
	switch position{
		case "Winger":
			kick="Low"
		case "Center":
			kick="Hybrid"
		case "Defenseman":
			kick="Mid"
		default:
			log.Fatal("Invalid hockey position")
	}
	return kick
}
