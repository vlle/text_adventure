package app_game

import (
  "net/http"
  "os"
	"github.com/go-chi/chi/v5"
  "github.com/golang-jwt/jwt/v5"
	"github.com/go-chi/chi/v5/middleware"
  // "github.com/vlle/text_adventure/game/game_logic"
)

func generate_token() {

  var (
    key []byte
    t   *jwt.Token
    s   string
  )
  
  key = []byte(os.Getenv("key"))  /* Load key from somewhere, for example an environment variable */
  t = jwt.New(jwt.SigningMethodHS256) 
  s, err := t.SignedString(key)
  if err != nil {
    panic(err)
  }
}

func RouteGame(r chi.Router) {
  r.Get("/signup", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("signup"))
  })
  r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("login"))
  })
  r.Get("/whereami", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("you are in the game"))
  })
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

