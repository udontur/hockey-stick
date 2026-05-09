package display

import (
	"fmt"
	"hockey-stick/hockey/types"
)

func DisplayResult(stick types.Stick) {
	fmt.Println(stick.Curve)
	fmt.Println(stick.Flex)
	fmt.Println(stick.Hand)
	fmt.Println(stick.Kick)
	fmt.Println(stick.Length)
}
