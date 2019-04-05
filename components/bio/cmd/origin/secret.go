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

package origin

import (
	"github.com/biome-sh/biome-go/components/bio/cmd/origin/secret"
	"github.com/spf13/cobra"
)

// SecretCmd ...
var SecretCmd = &cobra.Command{
	Use:   "secret",
	Short: "Commands relating to Biome secrets",
}

func init() {
	SecretCmd.PersistentFlags().StringP("auth", "z", "", "Authentication token for Builder")
	SecretCmd.PersistentFlags().StringP("url", "u", "https://bldr.habitat.sh/v1", `Specify an alternate Builder endpoint. If not specified, the value will be taken from the
HAB_BLDR_URL environment variable if defined. (default: https://bldr.habitat.sh)`)
	SecretCmd.PersistentFlags().StringP("origin", "o", "", "The origin for which the secret will be uploaded. Default is from 'HAB_ORIGIN' or cli.toml")

	SecretCmd.AddCommand(secret.DeleteCmd)
	SecretCmd.AddCommand(secret.ListCmd)
	SecretCmd.AddCommand(secret.UploadCmd)
}
