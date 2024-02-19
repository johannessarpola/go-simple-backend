package cmd

import (
	"fmt"

	"github.com/johannessarpola/go-simple-backend/pkg/server"
	"github.com/spf13/cobra"
	v "github.com/spf13/viper"
)

type config struct {
	Host  string `mapstructure:"host,omitempty"`
	Ports []int  `mapstructure:"ports,omitempty"`
}

func loadConfig() {
	//v.AutomaticEnv()
	v.SetConfigName("config.dev")
	v.SetConfigType("yaml")
	v.AddConfigPath("config")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(v.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic(fmt.Errorf("could not find config file: %w", err))
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}
}

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "runs either single server or multiple",
	Long:  "runs either single server or multiple",
	Run: func(cmd *cobra.Command, args []string) {

		var c config

		// Unmarshal the config into the struct
		if err := v.Unmarshal(&c); err != nil {
			panic(fmt.Errorf("could not unmarshal config file: %w", err))
		}

		for _, port := range c.Ports {
			server := server.NewServer(c.Host, port)
			go server.Run()
		}

	},
}

func init() {
	cobra.OnInitialize(loadConfig)

	rootCmd.Flags().String("host", "localhost", "host to run on")
	v.BindPFlag("host", rootCmd.Flags().Lookup("host"))

	rootCmd.Flags().IntSlice("ports", []int{8080}, "port(s) to run on")
	v.BindPFlag("ports", rootCmd.Flags().Lookup("ports"))
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
