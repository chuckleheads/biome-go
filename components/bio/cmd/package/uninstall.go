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

// UninstallCmd ...
var UninstallCmd = &cobra.Command{
	Use:   "uninstall <PKG_IDENT>",
	Short: "Safely uninstall a package and dependencies from the local filesystem",
	Args:  cobra.ExactArgs(1),
}

func init() {
	UninstallCmd.Flags().BoolP("dryrun", "d", false, "Just show what would be uninstalled, don't actually do it")
	UninstallCmd.Flags().Bool("no-deps", false, "Don't uninstall dependencies")

	UninstallCmd.Flags().String("exclude", "", "Identifier of one or more packages that should not be uninstalled. (ex: core/redis, core/busybox-static/1.42.2/21120102031201)")
}
