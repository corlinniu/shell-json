package main

import (
	"net/url"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		print("Usage: urlencode xxx")
		return
	}

	print(url.QueryEscape(os.Args[1]))
}
