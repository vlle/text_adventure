package services

import (
  "github.com/vlle/text_adventure/restful_db/internal/database"
  "github.com/vlle/text_adventure/restful_db/internal/models"
)

// type ItemSerivceError struct {
//   e error
//   error_type string
//   Proposed_http_code int
// }
// 
// func (e ItemSerivceError) Error() string {
//   return e.error_type
// }

// need custom-like error handling
func GetItem(name string) (models.Item, error) {
  dbpool, err := database.ConnectDatabase()
  if err != nil {
    return models.Item{}, err
  }

  item, err := database.SelectItem(dbpool, name)
  if err != nil {
    if err.Error() == "sql: no rows in result set" {

    }
    return models.Item{}, err
  }

  return item, nil
}
