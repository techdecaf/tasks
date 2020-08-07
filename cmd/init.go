package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
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
		taskfilePath := filepath.Join(pwd, "taskfile.yaml")
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
}

var taskfileSample = `
# documentation for tasks can be found @ https://github.com/techdecaf/tasks
options:
  log: true

variables:
  CI_PROJECT_NAME: "{{PWD | base}}"
  CI_COMMIT_TAG: "{{TRY ~~git describe --tags --always --dirty --abbrev=0~~}}"
  CI_COMMIT_REF_NAME: "{{TRY ~~git rev-parse --abbrev-ref HEAD~~}}"
  CI_COMMIT_SHA: "{{TRY ~~git rev-parse HEAD~~}}"

tasks:
  default:
    description: is the task that runs when no tasks have been specified.
    commands: [tasks list]

  clean:
    description: remove temporary files or directories
    commands: []

  dependencies:
    description: install all required dependencies
	commands: []

  test:
    description: run tests
	commands: []
	
  coverage:
    description: run test including coverage
    commands: []

  build:
    description: build current project
    commands: []

  deploy:
    description: deploy the current project
	commands: []
	
  upgrade:
    description: upgrade the current project
    commands: []
`
