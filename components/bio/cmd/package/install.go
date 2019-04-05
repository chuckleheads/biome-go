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
	"strings"

	"github.com/spf13/cobra"

	"github.com/biome-sh/biome-go/components/core/archive"
	"github.com/biome-sh/biome-go/components/core/ident"
	pkg "github.com/biome-sh/biome-go/components/bio/pkg/package"
)

var channel string

// InstallCmd ...
var InstallCmd = &cobra.Command{
	Use:   "install <PKG_IDENT_OR_ARTIFACT>",
	Short: "installs a Biome package",
	Long:  "TODO: Right now this is just for local artifacts",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		// Right now it better end in a .hart, we could unpack the archive to see if it is a hart
		// but it's not worth the time right now
		if strings.HasSuffix(path, ".hart") {
			pkg.FromArchive(archive.New(path))
		} else {
			// cheap validation check
			pkgIdent, err := ident.FromString(args[0])
			if err != nil {
				fmt.Println("Error with supplied ident: ", err)
			}
			pkg.FromIdent(pkgIdent, channel)
		}
	},
}

func init() {
	InstallCmd.Flags().StringP("auth", "z", "", "Authentication token for Builder")
	InstallCmd.Flags().StringP("url", "u", "https://bldr.habitat.sh/v1", `Specify an alternate Builder endpoint. If not specified, the value will be taken from the
HAB_BLDR_URL environment variable if defined. (default: https://bldr.habitat.sh)`)
	InstallCmd.Flags().StringVarP(&channel, "channel", "c", "stable", "Install from the specified release channel")

	InstallCmd.Flags().BoolP("binlink", "b", false, "Binlink all binaries from installed package(s)")
	InstallCmd.Flags().BoolP("force", "f", false, "Overwrite existing binlinks")

}
