package setup

import (
	"fmt"
	"os"

	"github.com/biome-sh/biome-go/components/bio/pkg/ui"
	"github.com/manifoldco/promptui"
)

func setupSupSecret() string {
	supSecretHelpText()
	if ui.PromptYesNo("Set up a default Biome Supervisor CtlGateway secret?") {
		prompt := promptui.Prompt{
			Label:   "Biome Supervisor CtlGateway secret",
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

func supSecretHelpText() {
	fmt.Println(`
  Enter your Biome Supervisor CtlGateway secret.`)
	fmt.Println()
}
