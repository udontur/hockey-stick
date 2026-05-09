package main

import (
	"hockey-stick/hockey/display"
	"hockey-stick/hockey/process"
)

func main() {
	var player = process.QuestionForm()
	var stick = process.ProcessResult(player)
	display.DisplayResult(stick)
}
