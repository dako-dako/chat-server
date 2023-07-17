package server

import (
	"fmt"

	"github.com/AXYGEN0141/chat-server/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "runserver",
	Short: "run chat API server",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := server.StartRouter()
		if err != nil {
			log.Fatal().Err(err).Msg("couldn't start server on specified port")
		}
		fmt.Println(app)
	},
}
