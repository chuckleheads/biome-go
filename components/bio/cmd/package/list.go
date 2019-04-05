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

// ListCmd ...
var ListCmd = &cobra.Command{
	Use:   "list [ORIGIN | PKG_IDENT]",
	Short: "List all versions of installed packages",
	Long:  "This is poor UX but replicated exactly. You must list --all OR a package ident OR an origin",
	Args:  cobra.ExactArgs(1),
}

func init() {
	ListCmd.Flags().BoolP("all", "a", false, "List all installed packages")
	ListCmd.Flags().StringP("origin", "o", "", "An origin to list")
}
