// Copyright © 2018 Elliott Davis elliott@excellent.io
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

package channel

import (
	"fmt"
	"os"

	"github.com/biome-sh/biome-go/components/builder-depot-client/client"
	"github.com/biome-sh/biome-go/components/bio/pkg/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ListCmd ...
var ListCmd = &cobra.Command{
	Use:   "list [OPTIONS] <CHANNEL>",
	Short: "Lists channels",
	Run: func(cmd *cobra.Command, args []string) {
		cli := client.New(viper.GetString("bldr_url"), viper.GetString("auth_token"))
		channels, err := cli.ListChannels(viper.GetString("origin"))
		if err != nil {
			ui.Fatal(err)
			os.Exit(1)
		}
		for _, channel := range channels {
			fmt.Println(channel.Name)
		}
	},
}
