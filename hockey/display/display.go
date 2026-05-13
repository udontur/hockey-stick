package display

import (
	"fmt"
	"os"
	"hockey-stick/hockey/types"
	"strconv"
	tea "charm.land/bubbletea/v2"
)

type info struct {
	name string
	value string
}

type model struct {
	stickInfo []info
	cursor int
}

func initialModel(stick types.Stick) model {
	var arr = make([]info, 5)
	arr[0].name="Hand"
	arr[0].value=stick.Hand
	arr[1].name="Curve"
	arr[1].value=stick.Curve
	arr[2].name="Flex"
	arr[2].value=strconv.Itoa(stick.Flex)
	arr[3].name="Kick"
	arr[3].value=stick.Kick
	arr[4].name="Length"
	arr[4].value=strconv.Itoa(stick.Length)
	return model{
		stickInfo: arr,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg:=msg.(type){
		case tea.KeyPressMsg:
			switch msg.String() {
				case "ctrl+c", "q":
					return m, tea.Quit
				case "up":
					if m.cursor > 0 {
						m.cursor--
					}
				case "down":
					if m.cursor < len(m.stickInfo)-1{
						m.cursor++
					}
			}
	}
	return m, nil
}

func (m model) View() tea.View {
	var header="Your ideal hockey stick\n\n"
	for i, spec := range m.stickInfo {
		cursor:=" "
		if m.cursor==i{
			cursor=">"
		}

		header+=fmt.Sprintf("%s #%v %v: %s\n", cursor, i, spec.name, spec.value)
	}
	header+="\nPress q to quite\n"
	return tea.NewView(header)
}

func DisplayResult(stick types.Stick) {
	var p=tea.NewProgram(initialModel(stick))
	if _, err:=p.Run(); err!=nil{
		fmt.Printf("Oh hell naw error: %v", err)
		os.Exit(1)
	}
	// fmt.Println(stick.Curve)
	// fmt.Println(stick.Flex)
	// fmt.Println(stick.Hand)
	// fmt.Println(stick.Kick)
	// fmt.Println(stick.Length)


}
