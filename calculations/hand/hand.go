package hand

import (
	"log"
)

func Get(broomTopHandPosition string) (hand string) {
	switch broomTopHandPosition {
		case "Left":
			hand="Right"
		case "Right":
			hand="Left"
		default:
			log.Fatal("Invalid broom top hand position. New mutation?")
	}
	return hand
}
