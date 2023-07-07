package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vlle/text_adventure/restful_db/internal/services"
)

func GetItem (w http.ResponseWriter, r *http.Request) {
    item_title := chi.URLParam(r, "key")
    item, err := services.GetItem(item_title)
    if err != nil {
      http.Error(w, http.StatusText(422), 422)
      return 
    }
    json_item, err := json.Marshal(item)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
		w.Write(json_item)
}
