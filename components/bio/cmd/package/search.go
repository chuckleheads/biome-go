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

package pkg

import (
	"fmt"

	"github.com/biome-sh/biome-go/components/builder-depot-client/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// SearchCmd ...
var SearchCmd = &cobra.Command{
	Use:   "search <SEARCH_TERM>",
	Short: "Search for a package in Builder",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cli := client.New(viper.GetString("bldr_url"), viper.GetString("auth_token"))
		pkgs := cli.SearchPackage(args[0])
		for _, pkg := range pkgs {
			fmt.Printf("%s/%s/%s/%s\n", pkg.Origin, pkg.Name, pkg.Version, pkg.Release)
		}
	},
}

func init() {
	SearchCmd.Flags().StringP("auth", "z", "", "Authentication token for Builder")
	SearchCmd.Flags().StringP("url", "u", "https://bldr.habitat.sh/v1", `Specify an alternate Builder endpoint. If not specified, the value will be taken from the
HAB_BLDR_URL environment variable if defined. (default: https://bldr.habitat.sh)`)
}
