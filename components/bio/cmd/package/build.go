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

import (
	"os"

	"github.com/spf13/cobra"
)

// BuildCmd ...
var BuildCmd = &cobra.Command{
	Use:   "build <PLAN_CONTEXT>",
	Short: "Builds a Plan using a Studio",
	Args:  cobra.ExactArgs(1),
}

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	BinlinkCmd.Flags().StringSliceP("keys", "k", []string{}, "Installs secret origin keys (ex: 'unicorn', 'acme,other,acme-ops')")
	BinlinkCmd.Flags().StringP("root", "r", "/bio/studios/", "Sets the Studio root")
	BinlinkCmd.Flags().StringP("src", "s", pwd, "Sets the source path")
}
