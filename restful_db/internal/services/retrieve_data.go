package services

import (
  "github.com/vlle/text_adventure/restful_db/internal/database"
  "github.com/vlle/text_adventure/restful_db/internal/models"
)


func GetItem(name string) (models.Item, error) {
  dbpool, err := database.ConnectDatabase()
  if err != nil {
    return models.Item{}, err
  }

  item, err := database.SelectItem(dbpool, name)
  if err != nil {
    return models.Item{}, err
  }

  return item, nil
}
