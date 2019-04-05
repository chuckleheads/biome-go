package setup

import (
	"fmt"
	"os"

	"github.com/biome-sh/biome-go/components/bio/pkg/ui"

	"github.com/manifoldco/promptui"
)

func setupPAT() string {
	patHelpText()
	if ui.PromptYesNo("Set up a default Biome personal access token?") {
		prompt := promptui.Prompt{
			Label:   "Biome personal access token",
			Default: "",
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

func patHelpText() {
	fmt.Println()
	ui.Heading("Biome Personal Access Token")
	fmt.Println(`
  While you can perform tasks like building and running Biome packages 
  without needing to authenticate with Builder, some operations like 
  uploading your packages to Builder, or checking status of your build 
  jobs from the Biome client will require you to use an access token.`)
	fmt.Println(`
  The Biome Personal Access Token can be generated via the Builder 
  Profile page (https://bldr.habitat.sh/#/profile). Once you have 
  generated your token, you can enter it here.`)
	fmt.Println(`
  If you would like to save your token for use by the Biome client, 
  please enter your access token. Otherwise, just enter No. `)
	fmt.Println(`
  For more information on using Builder, please read the documentation at 
  https://www.habitat.sh/docs/using-builder/`)
	fmt.Println()
}
