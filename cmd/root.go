package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "basket-api",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Cobra supports persistent flags, which, if defined here will be global for your application.
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is ./.config_default.yml)")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".config_default")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read config file %s\n", viper.ConfigFileUsed())
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, "Using default config file:", viper.ConfigFileUsed())

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName(".config")
	}

	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())

	if err := viper.MergeInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "failed merge config file %s\n", viper.ConfigFileUsed())
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
