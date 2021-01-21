package cmd

/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hadefication/ready/lib"

	"github.com/artdarek/go-unzip"
	"github.com/cavaliercoder/grab"
	"github.com/otiai10/copy"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize docker-compose and runtime files",
	Long: `This command will initialize supplied runtime to the current 
working directory. The runtime should contain at least
a docker-compose.yml file.

You can either supply a path or a URL for the runtime.
URL runtimes must be a zip file of a directory
that contains the docker-compose.yml.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(lib.Fatal("Definitely not ready, no runtime is supplied."))
			return
		}

		runtime := args[0]
		path, err := os.Getwd()

		fmt.Println(lib.Info("Initializing Ready files"))

		if err != nil {
			log.Println(err)
		}

		path = path + "/example"

		if lib.IsRuntimeURL(runtime) {
			// Download from url
			readyPath := path + "/ready"
			tempZipFilePath := path + "/url-runtime-temp"
			zipFile := tempZipFilePath + "/runtime.zip"

			grab.Get(zipFile, runtime)

			uz := unzip.New(zipFile, tempZipFilePath)
			err := uz.Extract()

			if err != nil {
				fmt.Println(err)
			}

			os.RemoveAll(zipFile)

			globs, err := filepath.Glob(tempZipFilePath + "/*")
			unzippedDir := globs[0]
			runtimeDir := filepath.Base(unzippedDir)
			runtimePath := readyPath + "/" + runtimeDir

			copy.Copy(tempZipFilePath+"/"+runtimeDir, readyPath+"/"+runtimeDir)

			os.RemoveAll(tempZipFilePath)

			lib.BackupDockerComposeFile(path, runtimePath)
		} else {

			readyPath := path + "/ready"
			runtimeName := filepath.Base(runtime)
			runtimePath := readyPath + "/" + runtimeName

			os.Mkdir(readyPath, 0700)

			copy.Copy(runtime, runtimePath)

			lib.BackupDockerComposeFile(path, runtimePath)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
