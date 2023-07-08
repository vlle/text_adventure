package jwt

import (
  "github.com/golang-jwt/jwt/v5"
  "os"
  "fmt"
  "strconv"
)

func GenerateToken(id int) string {
  var (
    key []byte
    t   *jwt.Token
  )
  
  str_id := strconv.Itoa(123)
  key = []byte(os.Getenv("key"))  /* Load key from somewhere, for example an environment variable */
  t = jwt.NewWithClaims(jwt.SigningMethodHS256,
  jwt.MapClaims{ 
    "iss": "vlle_game", 
    "sub": str_id, 
  })
  s, err := t.SignedString(key)
  if err != nil {
    panic(err)
  }
  return s
}


func DecodeToken(token string) string {
  key := []byte(os.Getenv("key"))  /* Load key from somewhere, for example an environment variable */
  decoded_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
    }

    return key, nil
  })

  if claims, ok := decoded_token.Claims.(jwt.MapClaims); ok && decoded_token.Valid {
    fmt.Println(claims["foo"], claims["nbf"])
    return claims["sub"].(string)
  } else {
    fmt.Println(err)
    return ""
  }
}
