package main

import (
	"github.com/spf13/cobra"

	"github.com/dashjay/baize/pkg/config"
)

func NewBazelServerCommand() *cobra.Command {
	cmd := &cobra.Command{}
	cfgPath := cmd.Flags().String("config", "/config.toml", "config file to use")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		_, err := config.NewConfigFromFile(*cfgPath)
		if err != nil {
			return err
		}
		return nil
	}
	return cmd
}

func main() {
	err := NewBazelServerCommand().Execute()
	if err != nil {
		panic(err)
	}
}
