package cmd

import (
	"github.com/AlecAivazis/survey"
	"github.com/devopsdays/devopsdays-cli/event"
	"github.com/devopsdays/devopsdays-cli/speaker"
	"github.com/devopsdays/devopsdays-cli/sponsor"
)

func mainPrompt() (err error) {
	var selection string

	for selection != "Quit the application" {
		prompt := &survey.Select{
			Message: "Select an action:",
			Options: []string{
				"Create a new event",
				"Create a new speaker",
				"Create a new sponsor",
				"Show a speaker",
				"Quit the application",
			},
		}
		err := survey.AskOne(prompt, &selection, nil)
		if err != nil {
			break
		}
		switch selection {
		case "Create a new event":
			event.CreateEvent("", "")
		case "Create a new speaker":
			createSpeakerPrompt("", "")
		case "Create a new sponsor":
			sponsor.CreateSponsor("")
		case "Show a speaker":
			showSpeakerPrompt("", "")
		}
	}

	return
}

func createSpeakerPrompt(city, year string) (err error) {
	var exitCode = true

	for exitCode {
		if city == "" {
			prompt := &survey.Input{
				Message: "Enter the city name:",
			}
			err := survey.AskOne(prompt, &city, survey.Required)
			// handle interrupts
			if err != nil {
				exitCode = false
				break
			}
		}

		if year == "" {
			prompt := &survey.Input{
				Message: "Enter the year:",
			}
			err := survey.AskOne(prompt, &year, survey.Required)
			// handle interrupts
			if err != nil {
				exitCode = false
				break
			}
		}
		speaker.CreateSpeaker("", city, year)
		prompt := &survey.Confirm{
			Message: "Do you want to add another speaker?",
		}
		err := survey.AskOne(prompt, &exitCode, nil)
		// handle interrupts
		if err != nil {
			exitCode = false
			break
		}
	}
	return
}

func showSpeakerPrompt(city, year string) (err error) {
	var exitCode = true
	for exitCode {
		if city == "" {
			prompt := &survey.Input{
				Message: "Enter the city name:",
			}
			err := survey.AskOne(prompt, &city, survey.Required)
			// handle interrupts
			if err != nil {
				exitCode = false
				break
			}
		}

		if year == "" {
			prompt := &survey.Input{
				Message: "Enter the year:",
			}
			err := survey.AskOne(prompt, &year, survey.Required)
			// handle interrupts
			if err != nil {
				exitCode = false
				break
			}
		}
		speaker.ShowSpeakers(city, year)

	}
	return
}
