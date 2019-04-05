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
	"github.com/biome-sh/biome-go/components/bio/cmd/package"
	"github.com/spf13/cobra"
)

// packageCmd represents the pkg command
var packageCmd = &cobra.Command{
	Use:     "package [SUBCOMMAND]",
	Short:   "Commands relating to Biome packages",
	Aliases: []string{"pkg"},
}

func init() {
	rootCmd.AddCommand(packageCmd)
	packageCmd.AddCommand(pkg.BindsCmd)
	packageCmd.AddCommand(pkg.BinlinkCmd)
	packageCmd.AddCommand(pkg.BuildCmd)
	packageCmd.AddCommand(pkg.ChannelsCmd)
	packageCmd.AddCommand(pkg.ConfigCmd)
	packageCmd.AddCommand(pkg.DemoteCmd)
	packageCmd.AddCommand(pkg.DependenciesCmd)
	packageCmd.AddCommand(pkg.EnvCmd)
	packageCmd.AddCommand(pkg.ExecCmd)
	packageCmd.AddCommand(pkg.ExportCmd)
	packageCmd.AddCommand(pkg.HashCmd)
	packageCmd.AddCommand(pkg.InfoCmd)
	packageCmd.AddCommand(pkg.InstallCmd)
	packageCmd.AddCommand(pkg.ListCmd)
	packageCmd.AddCommand(pkg.PathCmd)
	packageCmd.AddCommand(pkg.PromoteCmd)
	packageCmd.AddCommand(pkg.ProvidesCmd)
	packageCmd.AddCommand(pkg.SearchCmd)
	packageCmd.AddCommand(pkg.SignCmd)
	packageCmd.AddCommand(pkg.UninstallCmd)
	packageCmd.AddCommand(pkg.UploadCmd)
	packageCmd.AddCommand(pkg.VerifyCmd)

}
