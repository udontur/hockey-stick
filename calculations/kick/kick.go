package kick

func Get(position string) (kick string) {
	switch position{
		case "Winger":
			kick="Low"
		case "Center":
			kick="Hybrid"
		case "Defenseman":
			kick="Mid"
		default:
			kick="Invalid position"
	}
	return kick
}
