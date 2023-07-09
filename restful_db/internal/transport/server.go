package server

import (
	"encoding/json"
	"net/http"
  "strconv"
  //"fmt"
  // "log"

	"github.com/go-chi/chi/v5"
	"github.com/vlle/text_adventure/restful_db/internal/services"
)

func GetItem (w http.ResponseWriter, r *http.Request) {
  item_title := chi.URLParam(r, "key")
  item, err := services.GetItem(item_title)
  if err.E != nil {
    http.Error(w, http.StatusText(err.ProposedHttpCode()), err.ProposedHttpCode())
    return 
  }
  json_item, error := json.Marshal(item)
  if error != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Write(json_item)
}

func GetLocation(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "location_id")
  i, error := strconv.Atoi(id)
  if error != nil {
      panic(error)
  }
  location, err := services.GetLocation(i)
  if err.E != nil {
    http.Error(w, http.StatusText(err.ProposedHttpCode()), err.ProposedHttpCode())
    return 
  }
  json_location, error := json.Marshal(location)
  if error != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Write(json_location)
}

func GetMonster(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "monster_id")
  i, error := strconv.Atoi(id)
  if error != nil {
      panic(error)
  }
  monster, err := services.GetMonster(i)
  if err.E != nil {
    http.Error(w, http.StatusText(err.ProposedHttpCode()), err.ProposedHttpCode())
    return 
  }
  json_monster, error := json.Marshal(monster)
  if error != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Write(json_monster)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "user_id")
  i, error := strconv.Atoi(id)
  if error != nil {
      panic(error)
  }
  user, err := services.GetUser(i)
  if err.E != nil {
    http.Error(w, http.StatusText(err.ProposedHttpCode()), err.ProposedHttpCode())
    return 
  }
  json_user, error := json.Marshal(user)
  if error != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Write(json_user)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
  
  var N struct {
    Name string 
    Password string 
  }
  // bytes := make([]byte, 128)
  // r.Body.Read(bytes)
  decoder := json.NewDecoder(r.Body)
  error := decoder.Decode(&N)
  if error != nil {
      panic(error)
  }
  Name := N.Name
  Password := N.Password
  id, err := services.CreateUser(Name, Password, -1, -1)
  if err.E != nil {
    http.Error(w, http.StatusText(err.ProposedHttpCode()), err.ProposedHttpCode())
    return 
  }

  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(id)
}


func LoginUser(w http.ResponseWriter, r *http.Request) {
  
  var N struct {
    Name string 
    Password string 
  }
  decoder := json.NewDecoder(r.Body)
  error := decoder.Decode(&N)
  if error != nil {
      panic(error)
  }
  id, err := services.LoginUser(N.Name, N.Password)
  if err.E != nil {
    http.Error(w, http.StatusText(err.ProposedHttpCode()), err.ProposedHttpCode())
    return 
  }

  json.NewEncoder(w).Encode(id)
}
