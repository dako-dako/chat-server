package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"net/http"
)

var Cmd = &cobra.Command{
	Use:   "runserver",
	Short: "run chat API server",
	Run: func(cmd *cobra.Command, args []string) {
		r := chi.NewRouter()
		r.Use(middleware.Logger)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte("welcome"))
			if err != nil {
				log.Fatal().Err(err).Msg("home endpoint don't work")
			}
		})
		err := http.ListenAndServe(":3000", r)
		if err != nil {
			log.Fatal().Err(err).Msg("couldn't start server on specified port")
		}
	},
}
