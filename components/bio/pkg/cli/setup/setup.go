package setup

import (
	"fmt"

	"github.com/biome-sh/biome-go/components/bio/pkg/ui"
	"github.com/spf13/viper"

	"github.com/fatih/color"
)

var bioGreen = color.New(color.FgGreen, color.Bold)

// Run executes the setup cli sequence
func Run() {
	setupInit()
	origin := setupOrigin()
	if origin != "" {
		viper.Set("origin", origin)
	}
	viper.WriteConfig()
	pat := setupPAT()
	if pat != "" {
		viper.Set("auth", pat)
	}
	viper.WriteConfig()
	ctlSecret := setupSupSecret()
	if ctlSecret != "" {
		viper.Set("ctl_secret", ctlSecret)
	}
	viper.WriteConfig()
	setupAnalytics()
	setupEnd()
}

// SetupInit prints the initial help text for the cli setup
func setupInit() {
	fmt.Println()
	ui.Title("Biome CLI Setup")
	fmt.Println(`
  Welcome to bio setup. Let's get started.`)
}

func setupEnd() {
	fmt.Println()
	ui.Heading("CLI Setup Complete")
	fmt.Println()
	fmt.Println("  That's all for now. Thanks for using Biome!")
	fmt.Println()
}
