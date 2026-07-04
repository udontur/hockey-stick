package display

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/table"
	"charm.land/lipgloss/v2"
	"fmt"
	"hockey-stick/hockey/types"
	"os"
	"strconv"
)

var primary=lipgloss.Color("#FFFFFF")
var secondary=lipgloss.Color("#808080")
var theme=lipgloss.Color("#01B2BA")

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
		switch msg := msg.(type) {
		case tea.KeyPressMsg:
			switch msg.String() {
			case "q", "ctrl+c":
				return m, tea.Quit
			}
		}
		m.table, cmd = m.table.Update(msg)
		return m, cmd
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(secondary)

func (m model) View() tea.View {
	return tea.NewView(baseStyle.Render(m.table.View())+"\n"+m.table.HelpView()+"\n")
}

func DisplayResult(stick types.Stick) {
	var columns = []table.Column{
		{Title: "Spec", Width: 7},
		{Title: "Val", Width: 5},
	}

	var rows = []table.Row{
		{"Hand", stick.Hand},
		{"Curve", stick.Curve},
		{"Flex", strconv.Itoa(stick.Flex)},
		{"Kick", stick.Kick},
		{"Length", (strconv.Itoa(stick.Length)+"\"")},
	}

	var style=table.DefaultStyles()
	style.Header=lipgloss.NewStyle().
		Bold(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(secondary).
		BorderBottom(true).
		Foreground(secondary).
		PaddingLeft(1).
		PaddingRight(1).
		AlignVertical(lipgloss.Center)
	style.Cell=lipgloss.NewStyle().
		Foreground(primary).
		PaddingLeft(1).
		PaddingRight(1)

	var tableWidth=0
	for _, v := range columns{
		tableWidth+=v.Width
	}

	var displayTable=table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(len(rows)+1),
		table.WithWidth(tableWidth+4),
		table.WithStyles(style),
	)

	var p = tea.NewProgram(model{displayTable})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Bad bad: %v", err)
		os.Exit(1)
	}
}
