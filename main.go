package main

import (
	"hockey-stick/hockey/process"
	"hockey-stick/hockey/display"
)

func main(){
	var player=process.QuestionForm()
	var stick=process.ProcessResult(player)
	display.DisplayResult(stick)
}
