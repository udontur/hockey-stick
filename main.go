package main

import (
	"fmt"
	// "log"
	"hockey-stick/hockey/form"
	"hockey-stick/hockey/result"
)

func main(){

	curve, flex, hand, kick, length:=result.Process(form.QuestionForm())

	fmt.Println(curve)
	fmt.Println(flex)
	fmt.Println(hand)
	fmt.Println(kick)
	fmt.Println(length)
}
