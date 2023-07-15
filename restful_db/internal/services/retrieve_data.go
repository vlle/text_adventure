package services

import (
	"fmt"
	"log"
	"os"

	"github.com/vlle/text_adventure/restful_db/internal/database"
	"github.com/vlle/text_adventure/restful_db/internal/models"
	"golang.org/x/crypto/bcrypt"
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

func GetLocations() ([]models.LocationIMG, SerivceError) {
  dbpool, err := database.ConnectDatabase()
  if err != nil {
    log.Println(err)
    return []models.LocationIMG{}, SerivceError{E: err, ProposedCode: 503} 
  }

  locations, err := database.SelectLocations(dbpool)
  if err != nil {
    log.Println(err)
    if err.Error() == "no rows in result set" {
     return []models.LocationIMG{}, SerivceError{E: err, ProposedCode: 404}
    }
    return []models.LocationIMG{}, SerivceError{E: err, ProposedCode: 500}
  }
  return locations, SerivceError{}
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

  bytes, crypt_err := bcrypt.GenerateFromPassword([]byte(password), 10)
  hash := string(bytes)
  if crypt_err != nil {
    fmt.Fprintf(os.Stderr, "InsertUser.Error: %v\n", crypt_err)
    return 0, SerivceError{E: crypt_err, ProposedCode: 500}
  }

  id, err := database.InsertUser(dbpool, name, hash, img_id, location_id)
  if err != nil {
    log.Println(err)
    return -1, SerivceError{E: err, ProposedCode: 500}
  }
  return id, SerivceError{}
}

func LoginUser(name string, password string) (models.User, SerivceError) {
  dbpool, err := database.ConnectDatabase()
  if err != nil {
    log.Println(err)
    return models.User{}, SerivceError{E: err, ProposedCode: 503} 
  }

  user, err := database.SelectUserByName(dbpool, name)
  if err != nil {
    log.Println(err)
    if err.Error() == "no rows in result set" {
     return models.User{}, SerivceError{E: err, ProposedCode: 404}
    }
    return models.User{}, SerivceError{E: err, ProposedCode: 500}
  }

  err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
  if err != nil {
    fmt.Fprintf(os.Stderr, "LoginUser.Error: %v\n", err)
    return models.User{}, SerivceError{E: err, ProposedCode: 401}
  }
  return user, SerivceError{}
}
