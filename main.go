package main

import (
	"fmt"
	"log"
	"strconv"
	"errors"
	"hockey-stick/calculations/curve"
	"hockey-stick/calculations/flex"
	"hockey-stick/calculations/hand"
	"hockey-stick/calculations/kick"
	"hockey-stick/calculations/length"
	"charm.land/huh/v2"
)

func main(){
	var (
		hockeyType string
		position string
		strWeight string
		strHeight string
		broomTopHandPosition string
	)
	// hockeyType="Ball"
	// position="Defenseman"
	// weight=47
	strHeight="162"
	broomTopHandPosition="Left"

	var form=huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What kind of hockey do you play?").
				Options(
					huh.NewOption("Ice hockey", "Ice"),
					huh.NewOption("Roller hockey", "Roller"),
					huh.NewOption("Ball hockey", "Ball"),
				).
				Value(&hockeyType),
			huh.NewSelect[string]().
				Title("What position do you play?").
				Options(
					huh.NewOption("Winger (Forward)", "Winger"),
					huh.NewOption("Center (Forward)", "Center"),
					huh.NewOption("Defenseman", "Defenseman"),
				).
				Value(&position),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("What much do you weigh? (In kg)").
				Value(&strWeight).
				Validate(func(str string) error {
					if _, err:=strconv.Atoi(strWeight); err!=nil {
						return errors.New("Weight is not an integer")
					}
					return nil
				}),
		),
	)



	err:=form.Run()
	if err!=nil{
		log.Fatal(err)
	}

	weight, _:=strconv.Atoi(strWeight)
	height, _:=strconv.Atoi(strHeight)

	var curve=curve.Get(hockeyType)
	var flex=flex.Calculate(weight)
	var hand=hand.Get(broomTopHandPosition)
	var kick=kick.Get(position)
	var length=length.Calculate(height, hockeyType, position)

	fmt.Println(curve)
	fmt.Println(flex)
	fmt.Println(hand)
	fmt.Println(kick)
	fmt.Println(length)
}
