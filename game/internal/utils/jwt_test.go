package jwt

import (
	"strconv"
	"testing"
)

func TestGenerateToken(t *testing.T) {
  t.Setenv("key", "secret")
  value := 123
  got := GenerateToken(value)
  decoded := DecodeToken(got)
  if strconv.Itoa(value) != decoded {
    t.Errorf("GenerateToken(%d) = %s; want %s", value, got, decoded)
  }
}
