package process

import (
	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
	"errors"
	"log"
	"strconv"

	"hockey-stick/hockey/types"
)

func QuestionForm() (player types.Player) {
	var strHeight, strWeight string
	var form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What kind of hockey do you play?").
				Options(
					huh.NewOption("Ice hockey", "Ice"),
					huh.NewOption("Roller hockey", "Roller"),
					huh.NewOption("Ball hockey", "Ball"),
				).
				Value(&player.HockeyType),
			huh.NewSelect[string]().
				Title("What position do you play?").
				Options(
					huh.NewOption("Winger (Forward)", "Winger"),
					huh.NewOption("Center (Forward)", "Center"),
					huh.NewOption("Defenseman", "Defenseman"),
				).
				Value(&player.Position),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("How much do you weigh? (In kg)").
				Value(&strWeight).
				Validate(func(str string) error {
					if _, err := strconv.Atoi(strWeight); err != nil {
						return errors.New("Please enter an integer")
					}
					return nil
				}),
			huh.NewInput().
				Title("How tall are you? (In cm)").
				Value(&strHeight).
				Validate(func(str string) error {
					if _, err := strconv.Atoi(strHeight); err != nil {
						return errors.New("Please enter an integer")
					}
					return nil
				}),
			huh.NewSelect[string]().
				Title("How do you hold a broom? (Try it now!)").
				Options(
					huh.NewOption("Left hand on top", "Left"),
					huh.NewOption("Right hand on top", "Right"),
				).
				Value(&player.BroomTopHandPosition),
		),
	).WithTheme(
		huh.ThemeFunc(func(isDark bool) *huh.Styles {
			t := huh.ThemeBase(isDark)

			var (
				primary=lipgloss.Color("#FFFFFF")
				secondary=lipgloss.Color("#808080")
				theme=lipgloss.Color("#01B2BA")
			)

			// Question title
			t.Focused.Title = t.Focused.Title.
				Foreground(primary).
				Bold(true)
			// Sidebar select
			t.Focused.Base=t.Focused.Base.
				BorderForeground(secondary)
			// Selection
			t.Focused.SelectSelector = t.Focused.SelectSelector.
				Background(theme).
				Bold(true)
			t.Focused.SelectedOption = t.Focused.SelectedOption.
				Background(theme).
				Bold(true)
			t.Focused.UnselectedOption = t.Focused.UnselectedOption.
				Foreground(primary)
			// Text input
			t.Focused.TextInput.Text = t.Focused.TextInput.Text.
				Foreground(primary).
				Background(theme).
				Bold(true)
			t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.
				Foreground(primary).
				Background(theme).
				Bold(true)

			// Question title
			t.Blurred=t.Focused
			t.Blurred.Title=t.Focused.Title
			// Sidebar select
			t.Blurred.Base = t.Blurred.Base.
				BorderStyle(lipgloss.HiddenBorder())
			t.Blurred.SelectSelector = t.Blurred.SelectSelector.
				Foreground(theme).
				UnsetBackground().
				Bold(true)
			t.Blurred.SelectedOption = t.Blurred.SelectedOption.
				Foreground(theme).
				UnsetBackground().
				Bold(true)
			// Text input
			// TODO: It doesn't apply, blocked up upstream fix: charmbracelet/huh #770
			t.Blurred.TextInput.Text=t.Blurred.TextInput.Text.
				Foreground(primary).
				Background(theme).
				Bold(true)
			t.Blurred.TextInput.Prompt=t.Blurred.TextInput.Prompt.
				Foreground(primary).
				Background(theme).
				Bold(true)

			return t
		}))

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	player.Weight, _ = strconv.Atoi(strWeight)
	player.Height, _ = strconv.Atoi(strHeight)
	return
}
