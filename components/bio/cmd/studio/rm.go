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

package studio

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"

	"github.com/biome-sh/biome-go/components/bio/pkg/ui"
)

// RmCmd ...
var RmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Command to remove a Biome studio",
	Run: func(command *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		imageName := generateImageName(generateContainerName(dir))
		cli, err := client.NewEnvClient()
		if err != nil {
			panic(err)
		}
		removeImage(cli, imageName)
	},
}

func removeImage(cli *client.Client, imageName string) {
	_, err := cli.ImageRemove(context.Background(), imageName, types.ImageRemoveOptions{})
	if err != nil {
		if strings.Contains(err.Error(), "No such image") {
			ui.Warn(fmt.Sprintf("No Studio found for: %s", imageName))
			return
		}
		panic(err)
	}
	ui.Status(ui.Deleted, fmt.Sprintf("Studio for: %s", imageName))
}
