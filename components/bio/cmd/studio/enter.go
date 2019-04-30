package studio

import (
	"github.com/spf13/cobra"
)

// EnterCmd represents the studio enter command
var EnterCmd = &cobra.Command{
	Use:   "enter",
	Short: "Enter a Biome studio",
	Run: func(komand *cobra.Command, args []string) {
	},
}
