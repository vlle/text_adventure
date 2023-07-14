package app_game

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vlle/text_adventure/game/internal/transport"
)

type TokenStruct struct {
  Token string `json:"jwt_token"`
}



func RouteGame(r chi.Router) {


  r.Post("/signup", transport.Signup)
  r.Post("/login", transport.Login)

  r.Get("/whereami", transport.Whereami)
  // r.Get("/items", transport.items)
  // r.Post("/use", transport.use)
  // r.Post("/fight", transport.fight)
  // r.Post("/move", transport.move)

}

func LaunchGame() {
  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
  r.Route("/game", RouteGame)
  http.ListenAndServe(":3333", r)
}

