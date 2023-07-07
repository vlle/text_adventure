package server

import (
	"net/http"
  "github.com/vlle/text_adventure/restful_db/internal/services"
  "encoding/json"
)

func GetItem (w http.ResponseWriter, r *http.Request) {
    item, err := services.GetItem("Key")
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return 
    }
    json_item, err := json.Marshal(item)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
		w.Write(json_item)
}
