package game_logic

import (
  "strings"
  "fmt"
)


func SignupUser() {

}

type ReaderMap interface {
  ReadMap() ([][]int, []string)
}

// add '\n' to the end of each symbol in string (so it would be correctly encoded in json


func DrawMap(r ReaderMap) string {
  var sb strings.Builder
  _, img := r.ReadMap()
  for i:=0; i < len(img); i++ {
    if i % 3 == 0 && i != 0 {
      sb.WriteString("E")
    }
    sb.WriteString(img[i])
  }
  fmt.Println(sb.String())
  return sb.String()
}
