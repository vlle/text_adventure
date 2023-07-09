package app_game

import (
	"net/http"
	"os"
  "fmt"
  // "io"
  "encoding/json"


	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vlle/text_adventure/game/internal/utils"
)

type TokenStruct struct {
  Token string `json:"jwt_token"`
}

func signup(w http.ResponseWriter, r *http.Request) {

    newUrl := os.Getenv("REST_URL")
    if newUrl == "" {
      newUrl = "http://127.0.0.1:3000/"
    }

    resp, err := http.Post(newUrl + "user/signup", "application/json", r.Body)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    defer resp.Body.Close()

    var user_id int
    json.NewDecoder(resp.Body).Decode(&user_id)
    fmt.Println(user_id)

    token := jwt.GenerateToken(user_id)
    t := TokenStruct{Token: token}
    fmt.Println(token)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(t)
}

func login(w http.ResponseWriter, r *http.Request) {
  newUrl := os.Getenv("REST_URL")
  if newUrl == "" {
    newUrl = "http://127.0.0.1:3000/"
  }

  resp, err := http.Post(newUrl + "user/login", "application/json", r.Body)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  if resp.StatusCode != 200 {
    http.Error(w, "Invalid credentials", http.StatusUnauthorized)
    return
  }
  defer resp.Body.Close()

  var user_id int
  json.NewDecoder(resp.Body).Decode(&user_id)

  token := jwt.GenerateToken(user_id)
  t := TokenStruct{Token: token}

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(t)
}


func RouteGame(r chi.Router) {


  r.Post("/signup", signup)


  r.Post("/login", login)
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

