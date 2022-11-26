package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
)

func main() {
	token, _ := jwt.Parse(os.Args[2], func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Args[1]), nil
	})

	rst, _ := json.Marshal(token.Claims)
	fmt.Printf("%s\n", rst)
}
