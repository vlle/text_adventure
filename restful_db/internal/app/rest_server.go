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
  r.Get("/map", server.GetMap)
}

func RouteMonster(r chi.Router) {
  r.Get("/{monster_id}", server.GetMonster)
  // r.Post("/", server.CreateMonster)
}

func RouteUser(r chi.Router) {
  r.Post("/signup", server.PostUser)
  r.Post("/login", server.LoginUser)
  r.Get("/{user_id}", server.GetUser)
}

func LaunchServer() {
  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
  r.Route("/item", RouteItem)
  r.Route("/location", RouteLocation)
  r.Route("/user", RouteUser)
  r.Route("/monster", RouteMonster)
  http.ListenAndServe(":3000", r)
}
