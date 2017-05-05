package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "orbi",
	Short: `Project structure generator`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetDefault("OrbiPath", fmt.Sprintf("%s/.orbi", os.Getenv("HOME")))
	viper.SetDefault("PlansPath", fmt.Sprintf("%s/plans", viper.GetString("OrbiPath")))
	viper.SetDefault("TemplatesDir", "templates")

	viper.SetConfigName(".config")     // name of config file (without extension)
	viper.AddConfigPath("$HOME/.orbi") // adding home directory as first search path
}
