package process

import (
	calc "hockey-stick/calculations"

	"hockey-stick/hockey/types"
)

func ProcessResult(player types.Player) (stick types.Stick){
	stick.Curve=calc.GetCurve(player.HockeyType)
	stick.Flex=calc.GetFlex(player.Weight)
	stick.Hand=calc.GetHand(player.BroomTopHandPosition)
	stick.Kick=calc.GetKick(player.Position)
	stick.Length=calc.GetLength(player.Height, player.HockeyType, player.Position)
	return
}
