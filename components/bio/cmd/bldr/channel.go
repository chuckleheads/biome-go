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

package bldr

import (
	"github.com/biome-sh/biome-go/components/bio/cmd/bldr/channel"
	"github.com/spf13/cobra"
)

// ChannelCmd ...
var ChannelCmd = &cobra.Command{
	Use:   "channel [SUBCOMMAND]",
	Short: "Commands relating to Biome Builder channels",
}

func init() {
	ChannelCmd.AddCommand(channel.CreateCmd)
	ChannelCmd.AddCommand(channel.DestroyCmd)
	ChannelCmd.AddCommand(channel.ListCmd)
}
