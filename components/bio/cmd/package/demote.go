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

import "github.com/spf13/cobra"

// DemoteCmd ...
var DemoteCmd = &cobra.Command{
	Use:   "demote <PKG_IDENT> <CHANNEL>",
	Short: "Demote a package from a specified channel",
	Args:  cobra.ExactArgs(2),
}

func init() {
	DemoteCmd.Flags().StringP("auth", "z", "", "Authentication token for Builder")
	DemoteCmd.Flags().StringP("url", "u", "https://bldr.habitat.sh/v1", `Specify an alternate Builder endpoint. If not specified, the value will be taken from the
HAB_BLDR_URL environment variable if defined.`)
}
