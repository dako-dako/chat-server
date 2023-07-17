package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

func StartRouter() (*chi.Mux, error) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("welcome"))
		if err != nil {
			log.Fatal().Err(err).Msg("home endpoint don't work")
		}
	})
	err := http.ListenAndServe(":3000", r)
	return r, err
}
