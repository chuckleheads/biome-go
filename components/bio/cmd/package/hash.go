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
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/biome-sh/biome-go/components/bio/pkg/crypto"
	"github.com/biome-sh/biome-go/components/bio/pkg/ui"
)

// HashCmd ...
var HashCmd = &cobra.Command{
	Use:   "hash <SOURCE>",
	Short: "Generates a blake2b hashsum from a target at any given filepath",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		hash, err := crypto.HashFile(filePath)
		if err != nil {
			ui.Fatal(err)
			os.Exit(1)
		}
		fmt.Printf("%s  %s", hash, filePath)
	},
}
