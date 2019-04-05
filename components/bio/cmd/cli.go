// Copyright © 2018 Elliott Davis <elliott@excellent.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"

	"github.com/biome-sh/biome-go/components/bio/pkg/cli/setup"
	"github.com/spf13/cobra"
)

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli [SUBCOMMAND]",
	Short: "Commands relating to Biome runtime config",
}

var cliSetupCommand = &cobra.Command{
	Use:   "setup",
	Short: "Sets up the CLI with reasonable defaults",
	Run: func(cmd *cobra.Command, args []string) {
		setup.Run()
	},
}

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completers",
	Short: "Creates command-line completers for your shell.",
	Long: `
To load completers run

. <(bio cli completers)

To configure your bash shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(bio cli completers)`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenBashCompletion(os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.AddCommand(cliSetupCommand)
	cliCmd.AddCommand(completionCmd)
}
