package services

import (
  "github.com/vlle/text_adventure/restful_db/internal/database"
  "github.com/vlle/text_adventure/restful_db/internal/models"
  "log"
)

type SerivceError struct {
  E error
  ProposedCode int
}

func (e SerivceError) Error() string {
  return e.E.Error()
}

func (e SerivceError) ProposedHttpCode() int {
  return e.ProposedCode
}

// GetItem retrieves an item from the database
func GetItem(name string) (models.Item, SerivceError) {
  dbpool, err := database.ConnectDatabase()
  if err != nil {
    log.Println(err)
    return models.Item{}, SerivceError{E: err, ProposedCode: 503} 
  }

  item, err := database.SelectItem(dbpool, name)
  if err != nil {
    log.Println(err)
    if err.Error() == "no rows in result set" {
     return models.Item{}, SerivceError{E: err, ProposedCode: 404}
    }
    return models.Item{}, SerivceError{E: err, ProposedCode: 500}
  }
  return item, SerivceError{}
}

func GetLocation(id int) (models.Location, SerivceError) {
  dbpool, err := database.ConnectDatabase()
  if err != nil {
    log.Println(err)
    return models.Location{}, SerivceError{E: err, ProposedCode: 503} 
  }

  location, err := database.SelectLocation(dbpool, id)
  if err != nil {
    log.Println(err)
    if err.Error() == "no rows in result set" {
     return models.Location{}, SerivceError{E: err, ProposedCode: 404}
    }
    return models.Location{}, SerivceError{E: err, ProposedCode: 500}
  }
  return location, SerivceError{}
}

func GetMonster(id int) (models.Monster, SerivceError) {
  dbpool, err := database.ConnectDatabase()
  if err != nil {
    log.Println(err)
    return models.Monster{}, SerivceError{E: err, ProposedCode: 503} 
  }

  monster, err := database.SelectMonster(dbpool, id)
  if err != nil {
    log.Println(err)
    if err.Error() == "no rows in result set" {
     return models.Monster{}, SerivceError{E: err, ProposedCode: 404}
    }
    return models.Monster{}, SerivceError{E: err, ProposedCode: 500}
  }
  return monster, SerivceError{}
}

func GetUser(id int) (models.User, SerivceError) {
  dbpool, err := database.ConnectDatabase()
  if err != nil {
    log.Println(err)
    return models.User{}, SerivceError{E: err, ProposedCode: 503} 
  }

  user, err := database.SelectUser(dbpool, id)
  if err != nil {
    log.Println(err)
    if err.Error() == "no rows in result set" {
     return models.User{}, SerivceError{E: err, ProposedCode: 404}
    }
    return models.User{}, SerivceError{E: err, ProposedCode: 500}
  }
  return user, SerivceError{}
}

func CreateUser(name string, password string, img_id int, location_id int) (int, SerivceError) {
  dbpool, err := database.ConnectDatabase()
  if err != nil {
    log.Println(err)
    return -1, SerivceError{E: err, ProposedCode: 503} 
  }

  id, err := database.InsertUser(dbpool, name, password, img_id, location_id)
  if err != nil {
    log.Println(err)
    return -1, SerivceError{E: err, ProposedCode: 500}
  }
  return id, SerivceError{}
}
