package game_logic

import "strings"


func SignupUser() {

}

type ReaderMap interface {
  ReadMap() ([][]int, []string)
}


func DrawMap(r ReaderMap) string {
  var sb strings.Builder
  _, img := r.ReadMap()
  for i:=0; i < len(img); i++ {
    sb.WriteString(img[i])
  }
  return sb.String()
}
