package main

import (
  "github.com/vlle/text_adventure/game/cmd"
  "github.com/vlle/text_adventure/restful_db/cmd"
)


func main() {
  go func() { (game_start.GameStart()) }()
  go func() { (rest.RestStart()) }()
  select {}
}
