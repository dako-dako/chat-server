package cmd

import (
	"github.com/AXYGEN0141/chat-server/cmd/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "app",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal().Err(err).Msg("can't execute root cmd")
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(
		server.Cmd,
	)
}
