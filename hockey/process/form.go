package process

import (
	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
	"errors"
	"log"
	"strconv"

	"hockey-stick/hockey/theme"
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

			// Question title
			t.Focused.Title = t.Focused.Title.
				Foreground(theme.Primary).
				Bold(true)
			// Sidebar select
			t.Focused.Base = t.Focused.Base.
				BorderForeground(theme.Secondary)
			// Selection
			t.Focused.SelectSelector = t.Focused.SelectSelector.
				Background(theme.Accent).
				Bold(true)
			t.Focused.SelectedOption = t.Focused.SelectedOption.
				Background(theme.Accent).
				Bold(true)
			t.Focused.UnselectedOption = t.Focused.UnselectedOption.
				Foreground(theme.Primary)
			// Text input
			t.Focused.TextInput.Text = t.Focused.TextInput.Text.
				Foreground(theme.Primary).
				Background(theme.Accent).
				Bold(true)
			t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.
				Foreground(theme.Primary).
				Background(theme.Accent).
				Bold(true)

			// Question title
			t.Blurred = t.Focused
			t.Blurred.Title = t.Focused.Title
			// Sidebar select
			t.Blurred.Base = t.Blurred.Base.
				BorderStyle(lipgloss.HiddenBorder())
			t.Blurred.SelectSelector = t.Blurred.SelectSelector.
				Foreground(theme.Accent).
				UnsetBackground().
				Bold(true)
			t.Blurred.SelectedOption = t.Blurred.SelectedOption.
				Foreground(theme.Accent).
				UnsetBackground().
				Bold(true)
			// Text input
			// TODO: It doesn't apply, blocked up upstream fix: charmbracelet/huh #770
			t.Blurred.TextInput.Text = t.Blurred.TextInput.Text.
				Foreground(theme.Primary).
				Background(theme.Accent).
				Bold(true)
			t.Blurred.TextInput.Prompt = t.Blurred.TextInput.Prompt.
				Foreground(theme.Primary).
				Background(theme.Accent).
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
