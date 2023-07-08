// database 
package database

import (
  "context"
  "fmt"
  "os"

  "github.com/jackc/pgx/v5/pgxpool"
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

func SelectMonster(conn *pgxpool.Pool, id int) (models.Monster, error) {
  var m models.Monster
  query := "SELECT id, name, description, coalesce(image_id, -1), coalesce(location_id, -1) FROM monster WHERE id = $1"
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
  query := "SELECT id, name, coalesce(image_id, -1), location_id FROM user WHERE id = $1"
  err := conn.QueryRow(context.Background(), query , id).Scan(&u.ID, &u.Name, &u.ImageID, &u.LocationID)
  if err != nil {
    fmt.Fprintf(os.Stderr, "SelectUser.Error: %v\n", err)
    return u, err
  }
  return u, nil
}
