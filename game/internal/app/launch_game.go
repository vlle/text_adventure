package app_game

import (
  "net/http"
  "os"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
  // "github.com/vlle/text_adventure/game/game_logic"
)

func RouteGame(r chi.Router) {
  // send post query to another http

  newUrl := os.Getenv("REST_URL")
  if newUrl == "" {
    newUrl = "http://127.0.0.1:3000/"
  }
  // redirect post request to another http

  r.Post("/signup", func(w http.ResponseWriter, r *http.Request) {
    _, err := http.Post(newUrl + "user/signup", "application/json", r.Body)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    //defer resp.Body.Close()
    w.Write([]byte("user created"))
  })


  r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("login"))
  })
  r.Get("/whereami", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("you are in the game"))
  })
  // r.Get("/", func(w http.ResponseWriter, r *http.Request) {
  //   w.Write([]byte("welcome"))
  // })
}

func LaunchGame() {
  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
  r.Route("/game", RouteGame)
  http.ListenAndServe(":3333", r)
}

