package curve

func Get(hockeyType string) (curve string){
	switch hockeyType{
		case "Ice", "Roller":
			curve="P92"
		case "Ball":
			curve="P88"
	}
	return curve
}
