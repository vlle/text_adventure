package main

import (
	"bufio"
	"encoding/json"
  "strings"
	"fmt"
	"net/http"
	"os"
	"unicode/utf8"
)

func outputMapNetwork() {
  resp, err := http.Get("http://localhost:3333/game/whereami") 
    if err != nil {
      panic(err)
    }
    var result map[string]any
    json.NewDecoder(resp.Body).Decode(&result)
    if result["map"] != nil {
      outputMap(result["map"].(string))
    } else if result["message"] != nil {
    }
  }

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

func sendAuthRequest(scanner *bufio.Scanner, authType string) {

  green := "\033[1;32m"
  color_off := "\033[0m"
  fmt.Println("Please enter your username")
  fmt.Printf("%s%s%s", green, ">", color_off)
  scanner.Scan()
  username := scanner.Text()
  fmt.Println("Please enter your password")
  fmt.Printf("%s%s%s", green, ">", color_off)
  scanner.Scan()
  password := scanner.Text()

  resp, err := http.Post("http://localhost:3333/game/"+authType, "application/json", strings.NewReader(`{"username":"` + username + `", "password":"` + password + `"}`))
  if err != nil {
    panic(err)
  }
  var result map[string]any
  json.NewDecoder(resp.Body).Decode(&result)
  if result["jwt_token"] != nil {
    fmt.Println("Your login token:", result["jwt_token"].(string))
    fmt.Println("It is memorized during this session. Write it down to access your progress.")
    return 
  } else {
    fmt.Println("Something gone wrong during command")
  }
}

func signupHandler(scanner *bufio.Scanner) {
  sendAuthRequest(scanner, "signup")
}

func loginHandler(scanner *bufio.Scanner) {
  sendAuthRequest(scanner, "login")
}


func checkServer() {
  resp, err := http.Get("http://localhost:3333/game/healthcheck")
  if resp.StatusCode != 200 {
    fmt.Println("Error: ", resp.Status)
    os.Exit(1)
  }
  if err != nil {
    fmt.Println("Error: ", err)
    os.Exit(1)
  }
}

func main() {
  checkServer()

  green := "\033[1;32m"
  red := "\033[1;31m"
  color_off := "\033[0m"

  reader := bufio.NewReader(os.Stdin)
  scanner := bufio.NewScanner(reader)
  for {

    // fmt.Printf("%s%s%s\n", green,"Please enter your command", color_off)
    fmt.Printf("%s%s%s", green, ">", color_off)
    f := scanner.Scan()
    if f == false {
      fmt.Println("Error reading from console")
      continue
    }
    switch scanner.Text() {
      case "exit":
        fmt.Println("Bye!")
        os.Exit(0)
      case "help":
        fmt.Printf("%s%s%s\n", green,"Commands: help, exit, signup, login, move, whereami, use", color_off)
      case "signup":
        signupHandler(scanner)
      case "login":
        loginHandler(scanner)
      case "whereami":
        outputMapNetwork()
      default:
        fmt.Printf("%s%s%s\n", red, "Unknown command", color_off)
      }
  }
}
