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

package plan

import "github.com/spf13/cobra"

// InitCmd ...
var InitCmd = &cobra.Command{
	Use: "init [PKG_NAME]",
	Long: `
Generates common package specific configuration files. Executing without argument will create a 'habitat' directory in
your current folder for the plan. If 'PKG_NAME' is specified it will create a folder with that name. Environment
variables (those starting with 'pkg_') that are set will be used in the generated plan`,
	Args: cobra.ExactArgs(1),
}

func init() {
	InitCmd.Flags().Bool("windows", false, "Use a Windows Powershell plan template")
	InitCmd.Flags().Bool("with-all", false, "Generate omnibus plan with all available plan options")
	InitCmd.Flags().Bool("with-callbacks", false, "Include callback functions in template")
	InitCmd.Flags().Bool("with-docs", false, "Include plan options documentation")

	InitCmd.Flags().StringP("origin", "o", "", "Origin for the new app")
	InitCmd.Flags().StringP("scaffolding", "s", "", "Specify explicit Scaffolding for your app (ex: node, ruby)")
}
