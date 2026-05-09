package hand

func Get(broomTopHand string) (hand string) {
	switch broomTopHand {
		case "Left":
			hand="Right"
		case "Right":
			hand="Left"
	}
	return hand
}
