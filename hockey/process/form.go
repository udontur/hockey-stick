package process

import(
	"charm.land/huh/v2"
	"strconv"
	"errors"
	"log"

	"hockey-stick/hockey/types"
)

func QuestionForm() (player types.Player){
	var strHeight, strWeight string
	var form=huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What kind of hockey do you play?").
				Options(
					huh.NewOption("Ice hockey", "Ice"),
					huh.NewOption("Roller hockey", "Roller"),
					huh.NewOption("Ball hockey", "Ball"),
				).
				Value(&player.HockeyType),
			huh.NewSelect[string]().
				Title("What position do you play?").
				Options(
					huh.NewOption("Winger (Forward)", "Winger"),
					huh.NewOption("Center (Forward)", "Center"),
					huh.NewOption("Defenseman", "Defenseman"),
				).
				Value(&player.Position),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("How much do you weigh? (In kg)").
				Value(&strWeight).
				Validate(func(str string) error {
					if _, err:=strconv.Atoi(strWeight); err!=nil {
						return errors.New("Please enter an integer")
					}
					return nil
				}),
			huh.NewInput().
				Title("How tall are you? (In cm)").
				Value(&strHeight).
				Validate(func(str string) error {
					if _, err:=strconv.Atoi(strHeight); err!=nil {
						return errors.New("Please enter an integer")
					}
					return nil
				}),
			huh.NewSelect[string]().
				Title("How do you hold a broom? (Try it now!)").
				Options(
					huh.NewOption("Left hand on top", "Left"),
					huh.NewOption("Right hand on top", "Right"),
				).
				Value(&player.BroomTopHandPosition),
		),
	)

	err:=form.Run()
	if err!=nil{
		log.Fatal(err)
	}

	player.Weight, _=strconv.Atoi(strWeight)
	player.Height, _=strconv.Atoi(strHeight)
	return
}
