package main

import (
	"net/http"

	"github.com/spf13/cobra"

	_ "net/http/pprof"

	"github.com/btwiuse/baize/pkg/baize"
	"github.com/btwiuse/baize/pkg/config"
)

func NewBazelExecutorCommand() *cobra.Command {
	cmd := &cobra.Command{}
	cfgPath := cmd.Flags().String("config", "/config.toml", "config file to use")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		cfg, err := config.NewConfigFromFile(*cfgPath)
		if err != nil {
			return err
		}
		s, err := baize.New(cfg)
		if err != nil {
			return err
		}
		if pprofAddr := cfg.GetExecutorConfig().PprofAddr; pprofAddr != "" {
			go func() {
				http.ListenAndServe(pprofAddr, nil)
			}()
		}
		return s.Run()
	}
	return cmd
}

func main() {
	err := NewBazelExecutorCommand().Execute()
	if err != nil {
		panic(err)
	}
}
