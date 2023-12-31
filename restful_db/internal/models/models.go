package models

type Image struct {
  ID int `json:"id"`
  Name string `json:"name"`
  Url string `json:"url"`
  Emoji string `json:"emoji"`
}

type Point struct {
  X int `json:"x"`
  Y int `json:"y"`
}

type Location struct {
  ID int `json:"id"`
  Title string `json:"title"`
  XY Point `json:"xy"`
  Description string `json:"description"`
  ImageID int `json:"image_id"`
}

type LocationIMG struct {
  ID int `json:"id"`
  Title string `json:"title"`
  XY Point `json:"xy"`
  Description string `json:"description"`
  Image string `json:"image"`
}

type User struct {
  ID int `json:"id"`
  Name string `json:"name"`
  Password string `json:"password"`

  ImageID int `json:"image_id"`
  LocationID int `json:"location_id"`
} 

type Item struct {
  ID          int `json:"id"`
  Title       string `json:"title"`
  Description string `json:"description"`
  LocationID int `json:"location_id"`
  ImageID    int `json:"image_id"`
}

type Monster struct {
  ID int `json:"id"`
  Title string `json:"title"`
  Description string `json:"description"`
  LocationID int `json:"location_id"`
  ImageID int `json:"image_id"`
}
