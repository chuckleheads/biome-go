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

// BinlinkCmd ...
var BinlinkCmd = &cobra.Command{
	Use:   "binlink <PKG_IDENT> [BINARY]",
	Short: "Creates a binlink for a package binary in a common 'PATH' location",
	Args:  cobra.RangeArgs(1, 2),
}

func init() {
	BinlinkCmd.Flags().BoolP("force", "f", false, "Overwrite existing binlinks")
	BinlinkCmd.Flags().StringP("dest", "d", "", "Sets the destination directory (default: /bin)")
}
