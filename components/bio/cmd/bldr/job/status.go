// Copyright © 2018 NAME HERE elliott@excellent.io
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

package job

import "github.com/spf13/cobra"

// StatusCmd ...
var StatusCmd = &cobra.Command{
	Use:   "status <GROUP_ID|--origin <ORIGIN>>",
	Short: "Get the status of one or more job groups",
	Args:  cobra.MaximumNArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	StatusCmd.Flags().IntP("limit", "l", 10, "Limit how many job groups to retrieve, ordered by most recent")
	StatusCmd.Flags().BoolP("showjobs", "s", false, "Show the status of all build jobs for a retrieved job group")
}
