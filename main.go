package main

import (
	"fmt"
	"hockey-stick/calculations/curve"
	"hockey-stick/calculations/flex"
	"hockey-stick/calculations/hand"
	"hockey-stick/calculations/kick"
	"hockey-stick/calculations/length"
)

func main(){
	var hockeyType="Ball"
	var position="Defenseman"
	var weight=47
	var height=162
	var broomTopHand="Left"

	var curve=curve.Get(hockeyType)
	var flex=flex.Calculate(weight)
	var hand=hand.Get(broomTopHand)
	var kick=kick.Get(position)
	var length=length.Calculate(height, hockeyType, position)

	fmt.Println(curve)
	fmt.Println(flex)
	fmt.Println(hand)
	fmt.Println(kick)
	fmt.Println(length)
}
