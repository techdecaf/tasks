package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// VERSION is converted to the git tag at compile time using the make build command.
var VERSION string

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "runs a list of tasks as defined in your taskfile.yaml",
	Long:  `runs a list of tasks as defined in your taskfile.yaml`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if ver, _ := cmd.Flags().GetBool("version"); ver {
			fmt.Println(VERSION)
		}
	},
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tasks.yaml)")
	rootCmd.PersistentFlags().BoolP("silent", "s", false, "suppress log messages <overrides log option in taskfile.yaml>")

	rootCmd.PersistentFlags().StringP("task-file", "f", "", "use a specific taskfile.")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("version", "v", false, "Prints application version")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// if cfgFile != "" {
	// 	// Use config file from the flag.
	// 	viper.SetConfigFile(cfgFile)
	// } else {
	// 	// Find home directory.
	// 	home, err := homedir.Dir()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}

	// 	// Search config in home directory with name ".tasks" (without extension).
	// 	viper.AddConfigPath(home)
	// 	viper.SetConfigName(".tasks")
	// }

	// viper.AutomaticEnv() // read in environment variables that match

	// // If a config file is found, read it in.
	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Println("Using config file:", viper.ConfigFileUsed())
	// }
}