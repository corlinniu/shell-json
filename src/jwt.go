package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: jwt SECRET TOKEN\n")
		return
	}
	token, _ := jwt.Parse(os.Args[2], func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Args[1]), nil
	})
	if token == nil {
		return
	}
	rst, _ := json.Marshal(token.Claims)
	fmt.Printf("%s\n", rst)
}
