package process

import (
	// calc "hockey-stick/calculations"

	"hockey-stick/hockey/types"
)

func ProcessResult() (stick types.Stick) {
// func ProcessResult(player types.Player) (stick types.Stick) {
	// stick.Curve = calc.GetCurve(player.HockeyType)
	// stick.Flex = calc.GetFlex(player.Weight)
	// stick.Hand = calc.GetHand(player.BroomTopHandPosition)
	// stick.Kick = calc.GetKick(player.Position)
	// stick.Length = calc.GetLength(player.Height, player.HockeyType, player.Position)
	stick.Curve = "P92"
	stick.Flex = 55
	stick.Hand = "Right"
	stick.Kick = "Low"
	stick.Length = 58
	return
}
