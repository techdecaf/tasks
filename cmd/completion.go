package cmd

import (
	"fmt"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// Plugin structure for zsh
type Plugin struct {
	name   string
	path   string
	script string
}

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "generates a zsh plugin with completion scripts",
	Long:  `generates a zsh plugin with completion scripts`,
	Run: func(cmd *cobra.Command, args []string) {
		plugin := &Plugin{}
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal("completion", err)
		}

		plugin.name = rootCmd.Name()
		plugin.path = path.Join(home, ".oh-my-zsh/plugins/", plugin.name)
		plugin.script = path.Join(plugin.path, fmt.Sprintf("_%s", plugin.name))

		if _, err := os.Stat(plugin.path); os.IsNotExist(err) {
			os.MkdirAll(plugin.path, 0700)
		}

		if err := rootCmd.GenZshCompletionFile(plugin.script); err != nil {
			log.Fatal("completion", err)
		}

		fmt.Printf("a zsh completion file has been generated in %s \n", plugin.path)
		fmt.Println()
		fmt.Printf("to utilize the plugin, please add %s to the plugins section of your \n", plugin.name)
		fmt.Println(".zshrc file, and add add `compinit` to the bottom .zshrc file")
		fmt.Println()
		fmt.Println(".zshrc")
		fmt.Printf("065: `plugins( %s )`\n", plugin.name)
		fmt.Println("EoF: `compinit`")
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
