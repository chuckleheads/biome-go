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
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/biome-sh/biome-go/components/bio/pkg/crypto/keys"
	"github.com/biome-sh/biome-go/components/bio/pkg/fs"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

// These should 100% be piped through viper
const habStudioSecret = "HAB_STUDIO_SECRET_"
const habStudioImage = "habitat/default-studio-x86_64-linux:0.79.1"
const artifactPathEnvVar = "ARTIFACT_PATH"

var studioCmd = &cobra.Command{
	Use:   "studio",
	Short: "Commands relating to Biome studio",
	Args:  cobra.MaximumNArgs(1),
	Run: func(command *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		containerName := generateContainerName(dir)
		// Run the container first
		executeDocker(runContainer(args, dir, containerName))

		// At this point the container has exited. We need to commit the changes.
		commitContainer(containerName)
		removeContainer(containerName)
	},
}

func init() {
	rootCmd.AddCommand(studioCmd)
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
		panic("Docker was not found on your path. Please install it")
	}
	return path
}

func generateContainerName(dir string) string {
	return strings.ToLower(strings.Replace(strings.TrimPrefix(dir, "/"), "/", "--", -1))
}

func runContainer(args []string, currDir string, containerName string) []string {
	cmdArgs := []string{"run"}

	cmdArgs = append(cmdArgs, "--name", containerName)
	if isatty.IsTerminal(os.Stdout.Fd()) {
		cmdArgs = append(cmdArgs, "-it")
	}

	for _, envVar := range parseEnvVars() {
		cmdArgs = append(cmdArgs, "--env", fmt.Sprintf("%s=%s", envVar, os.Getenv(envVar)))
	}

	volumes := []string{fmt.Sprintf("%s:%s",
		currDir,
		"/src"),
		fmt.Sprintf("%s:/%s",
			keys.DefaultCacheKeyPath(),
			fs.CacheKeyPath)}
	cacheArtifactPath, present := os.LookupEnv(artifactPathEnvVar)
	if present {
		volumes = append(volumes, fmt.Sprintf("%s:/%s", cacheArtifactPath, fs.CacheArtifactPath))
	}

	for _, volume := range volumes {
		cmdArgs = append(cmdArgs, "--volume", volume)
	}
	if containerExists(containerName) {
		cmdArgs = append(cmdArgs, fmt.Sprintf("bio-studio.local/%s:latest", containerName))
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
	return cmdArgs
}

func containerExists(containerName string) bool {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	args := filters.NewArgs()
	args.Add("reference", fmt.Sprintf("bio-studio.local/%s:latest", containerName))
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{Filters: args})
	if err != nil {
		panic(err)
	}
	if len(images) == 1 {
		return true
	}
	return false
}

func commitContainer(containerName string) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	_, err = cli.ContainerCommit(context.Background(), containerName, types.ContainerCommitOptions{Reference: fmt.Sprintf("bio-studio.local/%s:latest", containerName)})
	if err != nil {
		panic(err)
	}
}

func removeContainer(containerName string) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	err = cli.ContainerRemove(context.Background(), containerName, types.ContainerRemoveOptions{RemoveVolumes: true})
	if err != nil {
		panic(err)
	}
}

func executeDocker(args []string) {
	cmd := exec.Command(findDocker(), args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

// Ideal flow
// docker run -it --name <NAME> --volumes <VOLUMES> --envs <ENVS> IMAGE <COMMAND>
// exit
// docker commit NAME NAME:latest
// docker rm --volumes NAME
// docker run -it --name NAME --volumes <VOLUMES> --envs <ENVS> NAME <COMMAND>
// repeat
// hab studio rm -> docker rmi --volumes NAME
