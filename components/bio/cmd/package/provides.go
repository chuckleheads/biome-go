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

// ProvidesCmd ...
var ProvidesCmd = &cobra.Command{
	Use:   "provides <FILE>",
	Short: "Search installed Biome packages for a given file",
	Args:  cobra.ExactArgs(1),
}

func init() {
	ListCmd.Flags().BoolP("path", "p", false, "Show full path to file")
	ListCmd.Flags().BoolP("real", "r", false, "Show fully qualified package names (ex: core/busybox-static/1.24.2/20160708162350)")
}
