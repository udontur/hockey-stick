package result

import (
	"hockey-stick/calculations/curve"
	"hockey-stick/calculations/flex"
	"hockey-stick/calculations/hand"
	"hockey-stick/calculations/kick"
	"hockey-stick/calculations/length"

	"hockey-stick/hockey/types"
)

func Process(player types.Player) (stick types.Stick){
	stick.Curve=curve.Get(player.HockeyType)
	stick.Flex=flex.Calculate(player.Weight)
	stick.Hand=hand.Get(player.BroomTopHandPosition)
	stick.Kick=kick.Get(player.Position)
	stick.Length=length.Calculate(player.Height, player.HockeyType, player.Position)
	return
}
