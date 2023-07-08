package app

import (
  "net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
  "github.com/vlle/text_adventure/restful_db/internal/transport"
)

func RouteItem(r chi.Router) {
  r.Get("/{key}", server.GetItem)
  // r.Post("/", server.CreateItem)
}

func RouteLocation(r chi.Router) {
  r.Get("/{location_id}", server.GetLocation)
  // r.Post("/", server.CreateLocation)
}

func RouteMonster(r chi.Router) {
  r.Get("/{monster_id}", server.GetMonster)
  // r.Post("/", server.CreateMonster)
}

func RouteUser(r chi.Router) {
  r.Get("/{user_id}", server.GetUser)
  // r.Post("/", server.CreatePlayer)
}

func LaunchServer() {
  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Route("/item", RouteItem)
  r.Route("/location", RouteLocation)
  r.Route("/user", RouteUser)
  r.Route("/monster", RouteMonster)
  http.ListenAndServe(":3000", r)
}
