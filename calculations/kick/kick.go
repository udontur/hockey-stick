package kick

func get(position string) string {
	var kick string
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
