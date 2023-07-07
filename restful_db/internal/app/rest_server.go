package app

import (
  "net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
  "github.com/vlle/text_adventure/restful_db/internal/transport"
)

func RouteItem(r chi.Router) {
  r.Get("/{key}", server.GetItem)
}

func LaunchServer() {
  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Route("/item", RouteItem)
  http.ListenAndServe(":3000", r)
}
