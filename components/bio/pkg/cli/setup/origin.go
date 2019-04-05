package setup

import (
	"fmt"
	"os"
	"os/user"

	"github.com/biome-sh/biome-go/components/bio/pkg/ui"
	"github.com/manifoldco/promptui"
)

func setupOrigin() string {
	originHelpText()
	if ui.PromptYesNo("Set up default origin?") {
		var username string
		u, err := user.Current()
		if err == nil {
			username = u.Username
		}

		prompt := promptui.Prompt{
			Label:   "Origin",
			Default: username,
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}
		return result
	}
	fmt.Println("  Okay, maybe another time.")
	return ""
}

func originHelpText() {
	fmt.Println()
	ui.Heading("Set up a default origin")
	fmt.Println(`
  Every package in Biome belongs to an origin, which indicates the 
  person or organization responsible for maintaining that package. Each 
  origin also has a key used to cryptographically sign packages in that 
  origin.`)
	fmt.Println(`
  Selecting a default origin tells package building operations such as 
  'bio pkg build' what key should be used to sign the packages produced. 
  If you do not set a default origin now, you will have to tell package 
  building commands each time what origin to use.`)
	fmt.Println(`
  For more information on origins and how they are used in building 
  packages, please consult the docs at 
  https://www.habitat.sh/docs/create-packages-build/`)
	fmt.Println()
}
