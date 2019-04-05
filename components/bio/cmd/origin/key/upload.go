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

// UploadCmd ...
var UploadCmd = &cobra.Command{
	Use:   "download <ORIGIN>",
	Short: "Upload origin keys to Builder",
	Args:  cobra.ExactArgs(1),
}

func init() {
	UploadCmd.Flags().BoolP("secret", "s", false, "Upload secret key in addition to the public key")

	UploadCmd.Flags().StringP("auth", "z", viper.GetString("auth_token"), "Authentication token for Builder")
	UploadCmd.Flags().StringP("url", "u", viper.GetString("bldr_url"), `Specify an alternate Builder endpoint. If not specified, the value will be taken from the
HAB_BLDR_URL environment variable if defined.`)

	UploadCmd.Flags().String("pubfile", "", "Path to a local public origin key file on disk")
	UploadCmd.Flags().String("secfile", "", "Path to a local secret origin key file on disk")

}
