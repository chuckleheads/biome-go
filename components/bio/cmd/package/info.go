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

// InfoCmd ...
var InfoCmd = &cobra.Command{
	Use:   "info <SOURCE>",
	Short: "Returns the Biome Artifact information",
	Args:  cobra.ExactArgs(1),
}

// var pkgArchiveCommand = &cobra.Command{
// 	Use:   "archive <PATH>",
// 	Short: "prints files in an archive",
// 	Args:  cobra.ExactArgs(1),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		path := args[0]
// 		archive := archive.New(path)
// 		fmt.Println(archive.GetMetadata())
// 	},
// }
