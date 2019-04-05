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

// DemoteCmd ...
var DemoteCmd = &cobra.Command{
	Use:   "demote <GROUP_ID> <CHANNEL>",
	Short: "Demote packages from a completed build job to a specified channel",
	Args:  cobra.ExactArgs(2),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	DemoteCmd.Flags().BoolP("interactive", "i", false, "Allow editing the list of demotable packages")
}
