package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"unicode/utf8"
)

func outputMap(v string) {
  for len(v) > 0 {
    r, size := utf8.DecodeRuneInString(v)
    v = v[size:]
    if string(r) == "E" {
      fmt.Printf("\n")
    } else {
      fmt.Printf("%c", r)
    }
  }
  fmt.Printf("\n")
}

func main() {

  resp, err := http.Get("http://localhost:3333/game/healthcheck")
  if resp.StatusCode != 200 {
    fmt.Println("Error: ", resp.Status)
    os.Exit(1)
  }
  reader := bufio.NewReader(os.Stdin)
  for {

    fmt.Println("Please enter your command")

    // Call the reader to read user's input
    scanner := bufio.NewScanner(reader)
    f := scanner.Scan()
    if f == false {
      fmt.Println("Error reading from console")
      continue
    }

    resp, err = http.Get("http://localhost:3333/game/" + scanner.Text()) 
    if err != nil {
      panic(err)
    }

    var result map[string]any
    json.NewDecoder(resp.Body).Decode(&result)
    if result["map"] != nil {
      outputMap(result["map"].(string))
    } else if result["message"] != nil {
      //
    }
  }

}
