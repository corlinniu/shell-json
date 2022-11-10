package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"os"
)

func main() {
	url := os.Args[1]
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	for _, xpath := range os.Args[2:] {
		node := htmlquery.FindOne(doc, xpath)
		fmt.Printf(node.FirstChild.Data)
	}
}
