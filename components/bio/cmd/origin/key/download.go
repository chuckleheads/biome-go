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

package key

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// DownloadCmd ...
var DownloadCmd = &cobra.Command{
	Use:   "download <ORIGIN> [REVISION]",
	Short: "Download origin key(s) to HAB_CACHE_KEY_PATH",
	Args:  cobra.RangeArgs(1, 2),
}

func init() {
	DownloadCmd.Flags().BoolP("encryption", "e", false, "Download public encryption key instead of public signing key")
	DownloadCmd.Flags().BoolP("secret", "s", false, "Download secret signing key instead of public signing key")

	DownloadCmd.Flags().StringP("auth", "z", viper.GetString("auth_token"), "Authentication token for Builder")
	DownloadCmd.Flags().StringP("url", "u", viper.GetString("bldr_url"), `Specify an alternate Builder endpoint. If not specified, the value will be taken from the
HAB_BLDR_URL environment variable if defined.`)
}
