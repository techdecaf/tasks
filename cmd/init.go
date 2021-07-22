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
		SAMPLE_TASKFILE = strings.ReplaceAll(SAMPLE_TASKFILE, "~~", "`")

		if _, err := os.Stat(taskfilePath); os.IsNotExist(err) {
			file, err := os.Create(taskfilePath)
			if err != nil {
				logger.Fatal("init", err)
			}
			defer file.Close()

			if _, err = io.WriteString(file, SAMPLE_TASKFILE); err != nil {
				logger.Fatal("init", err)
			}
		} else {
			logger.Fatal("init", fmt.Sprintf("taskfile already exists at %s, refusing to overwrite.", taskfilePath))
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
