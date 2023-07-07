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
  urlExample := "postgres://postgres:postgres@localhost:5500/rec"
  dbpool, err := pgxpool.New(context.Background(), urlExample)
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
  err := conn.QueryRow(context.Background(), query , name).Scan(&i.Id, &i.Title, &i.Description, &location_id, &i.Image_id)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error: %v\n", err)
    return i, err
  }
  return i, nil
}
