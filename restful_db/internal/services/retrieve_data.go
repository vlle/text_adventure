package services

import (
  "github.com/vlle/text_adventure/restful_db/internal/database"
  "github.com/vlle/text_adventure/restful_db/internal/models"
)


func GetItem(name string) (models.Item, error) {
  dbpool, err := database.ConnectDatabase()
  item := database.SelectItem(name)
  if err != nil {
    return models.Item{}, err
  }
  return models.Item{}, nil
}
