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

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/biome-sh/biome-go/components/bio/config"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bio [SUBCOMMAND]",
	Short: "\"A Biome is the natural environment for your services\" - Alan Turing",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bio/etc/cli.toml)")
	rootCmd.Flags().StringP("origin", "o", "", "Origin to use for requests")
	rootCmd.Flags().StringP("auth", "z", "", "Authentication token for Builder")
	rootCmd.Flags().StringP("url", "u", "https://bldr.habitat.sh", "Default Builder URL")
	rootCmd.Flags().String("artifact_cache", "~/.bio/cache/artifacts", "Default cache directory for artifacts")
	rootCmd.Flags().String("fs_root", "/", "Default root install path")
	viper.BindPFlag("origin", rootCmd.Flags().Lookup("origin"))
	viper.BindPFlag("auth_token", rootCmd.Flags().Lookup("auth"))
	viper.BindPFlag("bldr_url", rootCmd.Flags().Lookup("url"))
	viper.BindPFlag("artifact_cache", rootCmd.Flags().Lookup("artifact_cache"))
	viper.BindPFlag("fs_root", rootCmd.Flags().Lookup("fs_root"))
	viper.BindEnv("fs_root")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".bio" (without extension).
		viper.AddConfigPath(filepath.Join(home, ".hab", "etc")) // Preserve legacy behavior
		viper.AddConfigPath(filepath.Join(home, ".bio", "etc"))
		viper.SetConfigName("cli")
	}

	// If path has a ~/, convert it to $HOME
	if strings.HasPrefix(viper.GetString("artifact_cache"), "~/") {
		path, err := homedir.Expand(viper.GetString("artifact_cache"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.Set("artifact_cache", path)
	}

	viper.SetEnvPrefix("bio") // will be uppercased automatically
	viper.AutomaticEnv()      // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Printf("Err: %v", err)
	}
}

// ConfigFromViper fetches CLI config from viper
func ConfigFromViper() (*config.Config, error) {
	cfg := &config.Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		panic(err.Error())
	}
	return cfg, nil
}
