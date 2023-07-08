package app_game

import (
  "net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
  // "github.com/vlle/text_adventure/game/game_logic"
)

func RouteGame(r chi.Router) {
  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("welcome"))
  })
}

func LaunchGame() {
  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
  r.Route("/game", RouteGame)
  http.ListenAndServe(":3333", r)
}

