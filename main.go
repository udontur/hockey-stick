package main

import (
	"fmt"
	// "log"

	"hockey-stick/hockey/form"
	"hockey-stick/hockey/result"

	"hockey-stick/hockey/types"
)



func main(){
	var player types.Player=form.QuestionForm()
	var stick=result.Process(player)

	fmt.Println(stick.Curve)
	fmt.Println(stick.Flex)
	fmt.Println(stick.Hand)
	fmt.Println(stick.Kick)
	fmt.Println(stick.Length)
}
