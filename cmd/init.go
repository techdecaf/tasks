package cmd

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a task file in the current directory.",
	Long:  `initialize a task file in the current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		pwd, _ := os.Getwd()
		taskfilePath := path.Join(pwd, "taskfile.yaml")
		taskfileSample = strings.ReplaceAll(taskfileSample, "~~", "`")

		if _, err := os.Stat(taskfilePath); os.IsNotExist(err) {
			file, err := os.Create(taskfilePath)
			if err != nil {
				log.Fatal("init", err)
			}
			defer file.Close()

			if _, err = io.WriteString(file, taskfileSample); err != nil {
				log.Fatal("init", err)
			}
		} else {
			log.Fatal("init", fmt.Sprintf("taskfile already exists at %s, refusing to overwrite.", taskfilePath))
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("samples", "t", false, "Help message for toggle")
}

var taskfileSample = `
options:
  log: true

variables:
  CI_PROJECT_NAME: tasks
  CI_COMMIT_TAG: "{{EXEC ~~git describe --tags --always --dirty --abbrev=0~~}}"
  CI_COMMIT_REF_NAME: "{{EXEC ~~git rev-parse --abbrev-ref HEAD~~}}"
  CI_COMMIT_SHA: "{{EXEC ~~git rev-parse HEAD~~}}"
  S3_BUCKET: github.techdecaf.io

tasks:
  default:
    description: is the task that runs when no tasks have been specified.
    commands: [tasks list]

  dependencies:
    description: install all required dependencies
    commands: [go get, go install]

  build:
    description: build current project
    pre: [clean, dependencies]
    commands: []

  clean:
    description: remove files created as part of the build step.
    commands: ["rm -rf dist"]

  test:
    description: run tests
    commands: []

  lint:
    description: run linting
    commands: []

  deploy:
    description: deploy the current project
    commands: []
`
