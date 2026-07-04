package display

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/table"
	"charm.land/lipgloss/v2"
	"fmt"
	"hockey-stick/hockey/types"
	"os"
	"strconv"
	"strings"
)

var primary=lipgloss.Color("#FFFFFF")
var secondary=lipgloss.Color("#808080")
var theme=lipgloss.Color("#01B2BA")

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(secondary).
	Foreground(primary)

func cutHeader(body string) string{
	var tableHeaderLinePos=strings.Index(body, "\n")
	body=body[tableHeaderLinePos+1:]
	return body
}

type model struct {
	table table.Model
	title string
	tableWidth int
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

func (m model) View() tea.View {
	var header=lipgloss.NewStyle().
		Width(m.tableWidth).
		Bold(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(secondary).
		BorderBottom(true).
		Foreground(secondary).
		Padding(0, 1).
		Align(lipgloss.Center).
		Render(m.title)

	var body = cutHeader(m.table.View())
	var document=baseStyle.Render(header+"\n"+body)
	return tea.NewView(document+"\n")
}

func DisplayResult(stick types.Stick) {
	var columns = []table.Column{
		{Title: "", Width: 7},
		{Title: "", Width: 5},
	}

	var rows = []table.Row{
		{"Hand", stick.Hand},
		{"Curve", stick.Curve},
		{"Flex", strconv.Itoa(stick.Flex)},
		{"Kick", stick.Kick},
		{"Length", (strconv.Itoa(stick.Length)+"\"")},
	}

	var style=table.DefaultStyles()
	style.Header=lipgloss.NewStyle()
	style.Cell=lipgloss.NewStyle().
		Padding(0, 1)
	style.Selected=lipgloss.NewStyle().
		Bold(true).
		Background(theme)

	var tableWidth=0
	for _, v := range columns{
		tableWidth+=v.Width+2
	}

	var displayTable=table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(len(rows)+1),
		table.WithWidth(tableWidth),
		table.WithStyles(style),
	)

	var display = tea.NewProgram(model{
		table: displayTable,
		title: "Your Ideal Stick Specs",
		tableWidth: tableWidth,
	})
	if _, err := display.Run(); err != nil {
		fmt.Printf("Bad bad: %v", err)
		os.Exit(1)
	}
}
