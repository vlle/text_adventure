// database
package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/vlle/text_adventure/restful_db/internal/models"
)

func ConnectDatabase() (*pgxpool.Pool, error) {
  url := os.Getenv("DATABASE_URL")
  if url == "" {
    url = "postgres://postgres:postgres@localhost:5500/rec"
  }
  dbpool, err := pgxpool.New(context.Background(), url)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
    return nil, err
  }
  return dbpool, nil
}


func SelectItem(conn *pgxpool.Pool, name string) (models.Item, error) {
  var i models.Item 
  var location_id int
  query := "SELECT id, name, description, coalesce(location_id, -1), coalesce(image_id, -1) FROM item WHERE name = $1"
  err := conn.QueryRow(context.Background(), query , name).Scan(&i.ID, &i.Title, &i.Description, &location_id, &i.ImageID)
  if err != nil {
    fmt.Fprintf(os.Stderr, "SelectItem.Error: %v\n", err)
    return i, err
  }
  return i, nil
}

func SelectLocation(conn *pgxpool.Pool, id int) (models.Location, error) {
  var l models.Location
  query := "SELECT id, title, description, coalesce(image_id, -1) FROM location WHERE id = $1"
  err := conn.QueryRow(context.Background(), query , id).Scan(&l.ID, &l.Title, &l.Description, &l.ImageID)
  if err != nil {
    fmt.Fprintf(os.Stderr, "SelectLocation.Error: %v\n", err)
    return l, err
  }
  return l, nil
}

func SelectLocations(conn *pgxpool.Pool) ([]models.LocationIMG, error) {
  var locations []models.LocationIMG
  var point pgtype.Point
  query := "SELECT location.id, title, description, emoji, xy FROM location JOIN image ON location.image_id = image.id"
  rows, err := conn.Query(context.Background(), query)
  if err != nil {
    fmt.Fprintf(os.Stderr, "SelectLocations.Error: %v\n", err)
    return locations, err
  }
  defer rows.Close()
  for rows.Next() {
    var l models.LocationIMG
    err := rows.Scan(&l.ID, &l.Title, &l.Description, &l.Image, &point)
    l.XY.X = int(point.P.X)
    l.XY.Y = int(point.P.Y)
    if err != nil {
      fmt.Fprintf(os.Stderr, "SelectLocations.Error: %v\n", err)
      return locations, err
    }
    locations = append(locations, l)
  }
  return locations, nil
}

func SelectMonster(conn *pgxpool.Pool, id int) (models.Monster, error) {
  var m models.Monster
  query := "SELECT id, name, description, image_id FROM monster WHERE id = $1"
  err := conn.QueryRow(context.Background(), query , id).Scan(&m.ID, &m.Title, &m.Description, &m.ImageID)
  if err != nil {
    fmt.Fprintf(os.Stderr, "SelectMonster.Error: %v\n", err)
    return m, err
  }
  return m, nil
}

func SelectImage(conn *pgxpool.Pool, id int) (models.Image, error) {
  var i models.Image
  query := "SELECT id, name, coalesce(url, 'n/e'), emoji FROM image WHERE id = $1"
  err := conn.QueryRow(context.Background(), query , id).Scan(&i.ID, &i.Name, &i.Url, &i.Emoji)
  if err != nil {
    fmt.Fprintf(os.Stderr, "SelectImage.Error: %v\n", err)
    return i, err
  }
  return i, nil
}

func SelectUser(conn *pgxpool.Pool, id int) (models.User, error) {
  var u models.User
  query := "SELECT id, name, coalesce(image_id, -1), location_id FROM p_user WHERE id = $1"
  err := conn.QueryRow(context.Background(), query , id).Scan(&u.ID, &u.Name, &u.ImageID, &u.LocationID)
  if err != nil {
    fmt.Fprintf(os.Stderr, "SelectUser.Error: %v\n", err)
    return u, err
  }
  return u, nil
}


func SelectUserByName(conn *pgxpool.Pool, name string) (models.User, error) {
  var u models.User
  query := "SELECT id, name, password, coalesce(image_id, -1), location_id FROM p_user WHERE name = $1"
  err := conn.QueryRow(context.Background(), query, name).Scan(&u.ID, &u.Name, &u.Password, &u.ImageID, &u.LocationID)
  if err != nil {
    fmt.Fprintf(os.Stderr, "SelectUserByName.Error: %v\n", err)
    return u, err
  }
  return u, nil
}

func InsertUser(conn *pgxpool.Pool, name string, password string, image_id int, location_id int) (int, error) {
  if image_id <= 0 {
    image_id = 11 // default image = '🕵️‍♂️'
  }
  if location_id <= 0 {
    location_id = 5 // starting_location = '🪨'
  }


  query := "INSERT INTO p_user (name, password, hp, armor, fight_power, image_id, location_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
  var id int
  err := conn.QueryRow(context.Background(), query, name, password, 100, 0, 1, image_id, location_id).Scan(&id)
  if err != nil {
    fmt.Fprintf(os.Stderr, "InsertUser.Error: %v\n", err)
    return id, err
  }
  return id, nil
}

func UpdateUserLocation(conn *pgxpool.Pool, user_id int, location_id int) (models.User, error) {
  var u models.User
  query := "UPDATE user SET location_id = $1 WHERE id = $2 RETURNING id, name, coalesce(image_id, -1), coalesce(location_id, -1)"
  err := conn.QueryRow(context.Background(), query , location_id, user_id).Scan(&u.ID, &u.Name, &u.ImageID, &u.LocationID)
  if err != nil {
    fmt.Fprintf(os.Stderr, "UpdateUserLocation.Error: %v\n", err)
    return u, err
  }
  return u, nil
}
