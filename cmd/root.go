/*
Copyright © 2024 Nestri <>
*/
package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//go:embed nestri.ascii
var art string

var cfgFile string

type GameConfig struct {
	Directory  string
	Executable string
	GPU        int
	Vendor     string
	Resolution struct {
		Height int
		Width  int
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nestri",
	Short: "A CLI tool to manage your cloud gaming service",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nestri.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search for config in the current directory
		viper.AddConfigPath(".")
		viper.SetConfigName(".nestri")
		viper.SetConfigType("yaml")

		// If not found in current directory, check in $HOME/.nestri/
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// Find home directory.
				home, err := os.UserHomeDir()
				cobra.CheckErr(err)

				// Search config in home directory with name ".nestri" (without extension).
				viper.AddConfigPath(home)
				viper.SetConfigType("yaml")
				viper.SetConfigName(".nestri")
			}
		}

	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Could not find a config file in local directory or in $HOME directory:")
	}
}
