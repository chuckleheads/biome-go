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

package job

import (
	"fmt"
	"os"

	"github.com/biome-sh/biome-go/components/builder-depot-client/client"
	pkgType "github.com/biome-sh/biome-go/components/core/ident"
	"github.com/biome-sh/biome-go/components/bio/pkg/ui"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var group bool

// StartCmd ...
var StartCmd = &cobra.Command{
	Use:   "start <PKG_IDENT>",
	Short: "Schedule a build job or group of jobs",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jobStart(args[0])
	},
}

func init() {
	StartCmd.Flags().BoolVarP(&group, "group", "g", false, "Schedule jobs for this package and all of its reverse dependencies")
}

// TED: This belongs in pkg/jobs
func jobStart(ident string) {
	// TED: this should check for dependent rebuilds. If there are dependent rebuilds we need
	// to prompt the user and ask if they'd like to reverse dep build or just build this package.
	cli := client.New(viper.GetString("bldr_url"), viper.GetString("auth_token"))
	pkg, err := pkgType.FromString(ident)
	if err != nil {
		ui.Fatal(err)
		os.Exit(1)
	}
	rdeps, err := cli.FetchRdeps(pkg)
	if err != nil {
		ui.Fatal(err)
		os.Exit(1)
	}
	if group {
		if len(rdeps) > 0 {
			ui.Warn("Found the following reverse dependencies:")
			for _, rdep := range rdeps {
				ui.Warn(rdep.String())
			}
			question := `If you choose to start a group build for this package,
	all of the above will be built as well. Is this what you want?"`
			if !ui.PromptYesNo(question) {
				ui.Fatal(fmt.Errorf("Aborted"))
				os.Exit(1)
			}
		}
	}
	ui.Status(ui.Creating, fmt.Sprintf("build job for %s.", ident))
	err = cli.ScheduleJob(pkg, !group)
	if err != nil {
		ui.Fatal(err)
		os.Exit(1)
	}

	ui.Status(ui.Created, fmt.Sprintf("build job %s.", ident))
}
