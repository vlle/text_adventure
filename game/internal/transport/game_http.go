package transport

import (
  "net/http"
  "os"
  "encoding/json"
  "fmt"
	"github.com/vlle/text_adventure/game/internal/utils"
	"github.com/vlle/text_adventure/game/internal/services"
)

type TokenStruct struct {
  Token string `json:"jwt_token"`
}

type Point struct {
  X int `json:"x"`
  Y int `json:"y"`
}

type LocationIMG struct {
  ID int `json:"id"`
  Title string `json:"title"`
  XY Point `json:"xy"`
  Description string `json:"description"`
  Image string `json:"image"`
}

// write implementation of fizzbuzz in golang

func (l *LocationIMG) ReadMap() ([]int, string) {
  locs := make([]int, 2)
  locs[0] = l.XY.X
  locs[1] = l.XY.X
  img := l.Image
  return locs, img
}

type locations  []LocationIMG

func (l locations) ReadMap() ([][]int, []string) {
  locs := make([][]int, len(l))
  img := []string{}
  for i:=0; i < len(l); i++ {
    locs[i] = make([]int, 2)
    locs[i][0] = l[i].XY.X
    locs[i][1] = l[i].XY.X
    img = append(img, l[i].Image)
  }
  return locs, img
}

type Map struct {
  Map string `json:"map"`
}

func Signup(w http.ResponseWriter, r *http.Request) {

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

func Login(w http.ResponseWriter, r *http.Request) {
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


func Whereami(w http.ResponseWriter, r *http.Request) {
  newUrl := os.Getenv("REST_URL")
  if newUrl == "" {
    newUrl = "http://127.0.0.1:3000/"
  }

  resp, err := http.Get(newUrl + "location/map")
  var locs locations

  err = json.NewDecoder(resp.Body).Decode(&locs)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  var m Map 

  m.Map = game_logic.DrawMap(locs)
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(m)

}
