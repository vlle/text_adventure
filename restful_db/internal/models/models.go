package models

type Image struct {
  id int
  name string
  url string
  emoji string
}


type User struct {
  id int
  name string
  password string
  image_id int
}

type Point struct {
  X int
  Y int
}


type Item struct {
  Id          int
  Title       string
  Description string
  Xy          Point
  Image_id    int
}

type Location struct {
  id int
  title string
  description string
  xy Point
  image_id int
}

type Monster struct {
  id int
  title string
  description string
  xy Point
  image_id int
}
