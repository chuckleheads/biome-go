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
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/biome-sh/biome-go/components/bio/pkg/crypto/keys"
	"github.com/biome-sh/biome-go/components/bio/pkg/fs"
	"github.com/biome-sh/biome-go/components/bio/pkg/ui"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// These should 100% be piped through viper
const habStudioSecret = "HAB_STUDIO_SECRET_"
const habStudioImage = "habitat/default-studio-x86_64-linux:0.79.1"

// RunCmd ...
var RunCmd = &cobra.Command{
	Use:     "enter",
	Aliases: []string{"build"},
	Short:   "Command to remove a Biome studio",
	Run: func(command *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		containerName := generateContainerName(dir)
		cli, err := client.NewEnvClient()
		if err != nil {
			panic(err)
		}
		// prepend the initial arg to studio to our arg chain
		args = append([]string{command.CalledAs()}, args...)
		// Run the container first, catch any errors but don't process them right away.
		// We still want to clean up the container so we don't get conflicts
		err = runContainer(cli, args, dir, containerName)
		// At this point the container has exited. We need to commit the changes.
		commitContainer(cli, containerName)
		// Remove old container we did work in since we saved our changes to the image
		removeContainer(cli, containerName)
		if err != nil {
			ui.Fatal(err)
		}
	},
}

func parseEnvVars() []string {
	// This is all a terrible mistake and should be piped through viper
	envVars := []string{"DEBUG",
		"DO_CHECK",
		"HAB_AUTH_TOKEN",
		"HAB_BLDR_URL",
		"HAB_BLDR_CHANNEL",
		"HAB_NOCOLORING",
		"HAB_LICENSE",
		"HAB_ORIGIN",
		"HAB_ORIGIN_KEYS",
		"HAB_STUDIO_BACKLINE_PKG",
		"HAB_STUDIO_NOSTUDIORC",
		"HAB_STUDIO_SUP",
		"HAB_UPDATE_STRATEGY_FREQUENCY_MS",
		"http_proxy",
		"https_proxy",
		"RUST_LOG",
	}

	for _, key := range os.Environ() {
		if strings.HasPrefix(key, habStudioSecret) {
			envVars = append(envVars, key)
		}
	}
	return envVars
}

func findDocker() string {
	path, err := exec.LookPath("docker")
	if err != nil {
		ui.Fatal(errors.New("Docker was not found on your path. Please install it"))
	}
	return path
}

func generateContainerName(dir string) string {
	return strings.ToLower(strings.Replace(strings.TrimPrefix(dir, "/"), "/", "--", -1))
}

func generateImageName(containerName string) string {
	return fmt.Sprintf("bio-studio.local/%s:latest", containerName)
}

// TED TODO: find a way to use the docker client sanely to get an interactive tty
func runContainer(cli *client.Client, args []string, currDir string, containerName string) error {
	cmdArgs := []string{"run"}

	cmdArgs = append(cmdArgs, "--name", containerName)
	if isatty.IsTerminal(os.Stdout.Fd()) {
		cmdArgs = append(cmdArgs, "-it")
	}

	for _, envVar := range parseEnvVars() {
		cmdArgs = append(cmdArgs, "--env", fmt.Sprintf("%s=%s", envVar, os.Getenv(envVar)))
	}

	volumes := []string{
		fmt.Sprintf("%s:/%s", currDir, "src"),
		fmt.Sprintf("%s:/%s", keys.DefaultCacheKeyPath(), fs.CacheKeyPath),
		fmt.Sprintf("%s:/%s", viper.GetString("artifact_cache"), fs.CacheArtifactPath),
	}

	for _, volume := range volumes {
		cmdArgs = append(cmdArgs, "--volume", volume)
	}
	if containerExists(cli, containerName) {
		cmdArgs = append(cmdArgs, generateImageName(containerName))
	} else {
		cmdArgs = append(cmdArgs, habStudioImage)
	}
	// This looks gross but we need a way to append the args to cmdArgs and also remove a -D flag
	// for legacy hab support
	for _, arg := range args {
		if arg != "-D" {
			cmdArgs = append(cmdArgs, arg)
		}
	}
	cmd := exec.Command(findDocker(), cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	return err
}

func containerExists(cli *client.Client, containerName string) bool {
	args := filters.NewArgs()
	args.Add("reference", generateImageName(containerName))
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{Filters: args})
	if err != nil {
		ui.Warn(err.Error())
	}
	if len(images) == 1 {
		return true
	}
	return false
}

func commitContainer(cli *client.Client, containerName string) {
	_, err := cli.ContainerCommit(context.Background(), containerName, types.ContainerCommitOptions{Reference: generateImageName(containerName)})
	if err != nil {
		ui.Warn(err.Error())
	}
}

func removeContainer(cli *client.Client, containerName string) {
	err := cli.ContainerRemove(context.Background(), containerName, types.ContainerRemoveOptions{RemoveVolumes: true})
	if err != nil {
		ui.Warn(err.Error())
	}
}
