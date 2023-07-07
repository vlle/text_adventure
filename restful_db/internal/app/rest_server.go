package app

import (
  "net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
  "github.com/vlle/text_adventure/restful_db/internal/transport"
)

func LaunchServer() {
  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Get("/", server.GetItem)
  http.ListenAndServe(":3000", r)
}
