package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

type StatusComponent struct {
	icon    string
	message string
	color   color.Attribute
}

type StatusEnum int

const (
	Added StatusEnum = iota
	Adding
	Applying
	Cached
	Canceled
	Canceling
	Created
	Creating
	Deleting
	Deleted
	Demoted
	Demoting
	Determining
	Downloading
	Encrypting
	Encrypted
	Found
	Generated
	Generating
	Installed
	Missing
	Promoted
	Promoting
	Signed
	Signing
	Uploaded
	Uploading
	Using
	Verified
	Verifying
)

func getStatus(status StatusEnum) StatusComponent {
	switch status {
	case Added:
		return StatusComponent{`↑`, "Added", color.FgGreen}
	case Adding:
		return StatusComponent{`☛`, "Adding", color.FgGreen}
	case Applying:
		return StatusComponent{`↑`, "Applying", color.FgGreen}
	case Cached:
		return StatusComponent{`☑`, "Cached", color.FgGreen}
	case Canceled:
		return StatusComponent{`✓`, "Cancelled", color.FgGreen}
	case Canceling:
		return StatusComponent{`☛`, "Cancelling", color.FgGreen}
	case Created:
		return StatusComponent{`✓`, "Created", color.FgGreen}
	case Creating:
		return StatusComponent{`Ω`, "Creating", color.FgGreen}
	case Deleting:
		return StatusComponent{`☒`, "Deleting", color.FgGreen}
	case Deleted:
		return StatusComponent{`✓`, "Deleted", color.FgGreen}
	case Demoted:
		return StatusComponent{`✓`, "Demoted", color.FgGreen}
	case Demoting:
		return StatusComponent{`→`, "Demoting", color.FgGreen}
	case Determining:
		return StatusComponent{`☁`, "Determining", color.FgGreen}
	case Downloading:
		return StatusComponent{`↓`, "Downloading", color.FgGreen}
	case Encrypting:
		return StatusComponent{`☛`, "Encrypting", color.FgGreen}
	case Encrypted:
		return StatusComponent{`✓`, "Encrypted", color.FgGreen}
	case Found:
		return StatusComponent{`→`, "Found", color.FgCyan}
	case Generated:
		return StatusComponent{`→`, "Generated", color.FgCyan}
	case Generating:
		return StatusComponent{`☛`, "Generating", color.FgGreen}
	case Installed:
		return StatusComponent{`✓`, "Installed", color.FgGreen}
	case Missing:
		return StatusComponent{`∵`, "Missing", color.FgRed}
	case Promoted:
		return StatusComponent{`✓`, "Promoted", color.FgGreen}
	case Promoting:
		return StatusComponent{`→`, "Promoting", color.FgGreen}
	case Signed:
		return StatusComponent{`✓`, "Signed", color.FgCyan}
	case Signing:
		return StatusComponent{`☛`, "Signing", color.FgCyan}
	case Uploaded:
		return StatusComponent{`✓`, "Uploaded", color.FgGreen}
	case Uploading:
		return StatusComponent{`↑`, "Uploading", color.FgGreen}
	case Using:
		return StatusComponent{`→`, "Using", color.FgGreen}
	case Verified:
		return StatusComponent{`✓`, "Verified", color.FgGreen}
	case Verifying:
		return StatusComponent{`☛`, "Verifying", color.FgGreen}
	default:
		return StatusComponent{}
	}
}

// Begin logs begin messages
func Begin(message string) {
	symbol := `»`
	paint(color.FgYellow).Printf("%s %s\n", symbol, message)
}

// End logs end messages
func End(message string) {
	symbol := `★`
	paint(color.FgBlue).Printf("%s %s\n", symbol, message)
}

// Status logs status messages with a given const value and message string
func Status(status StatusEnum, message string) {
	stat := getStatus(status)
	paint(stat.color).Printf("%s %s ", stat.icon, stat.message)
	fmt.Println(message)
}

// Info logs info messages
func Info(message string) {
	fmt.Println(message)
}

// Warn logs warn messages
func Warn(message string) {
	symbol := `∅`
	paint(color.FgYellow).Printf("%s %s\n", symbol, message)
}

// Fatal logs fatal messages
func Fatal(message error) {
	symbol := `✗✗✗`
	paint(color.FgRed).Println(symbol)
	for _, str := range strings.Split(message.Error(), "\n") {
		paint(color.FgRed).Printf("%s %s\n", symbol, str)
	}
	paint(color.FgRed).Println(symbol)
}

func Title(message string) {
	paint(color.FgGreen, color.Underline).Println(message)
}

func Heading(message string) {
	paint(color.FgGreen).Println(message)
}

func PromptYesNo(labelText string) bool {
	prompt := promptui.Select{
		Label: labelText,
		Items: []string{"Yes", "No", "Quit"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	switch result {
	case "Yes":
		return true
	case "No":
		return false
	case "Quit":
		os.Exit(0)
	}
	return false
}

func paint(pcolor color.Attribute, args ...color.Attribute) *color.Color {
	return color.New(pcolor, color.Bold).Add(args...)
}
