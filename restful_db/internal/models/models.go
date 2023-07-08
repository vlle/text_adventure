package models

type Image struct {
  ID int
  Name string
  Url string
  Emoji string
}

type Point struct {
  X int
  Y int
}

type Location struct {
  ID int
  Title string
  XY Point
  Description string
  ImageID int
}

type User struct {
  ID int
  Name string
  Password string

  ImageID int
  LocationID int
}

type Item struct {
  ID          int
  Title       string
  Description string
  LocationID int
  ImageID    int
}

type Monster struct {
  ID int
  Title string
  Description string
  LocationID int
  ImageID int
}
