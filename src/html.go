package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("usage: bin/html sep file.html xpath1 xpath2 ...\n")
		return
	}

	sep := os.Args[1]
	file := os.Args[2]
	var err error
	var doc *html.Node
	if strings.Index(file, "http://") >= 0 || strings.Index(file, "https://") >= 0 {
		doc, err = htmlquery.LoadURL(file)
	} else {
		doc, err = htmlquery.LoadDoc(file)
	}
	if err != nil {
		fmt.Printf("err: %v", err.Error())
		return
	}

	var row []string
	for _, xpath := range os.Args[3:] {
		node := htmlquery.FindOne(doc, xpath)
		if node == nil || node.FirstChild == nil {
			row = append(row, "")
		} else {
			s := strings.ReplaceAll(node.FirstChild.Data, "\n", "")
			s = strings.ReplaceAll(s, "\r", "")
			s = strings.Trim(s, " ")
			row = append(row, s)
		}
	}

	fmt.Printf("%s\n", strings.Join(row, sep))
}
