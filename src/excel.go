package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: excel file.xlsx sheet 'age = {age}'\n")
		return
	}

	file := os.Args[1]
	sheet := os.Args[2]
	txtTpl := os.Args[3]

	f, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows := f.GetRows(sheet)
	head := rows[0]
	for _, row := range rows[1:] {
		rawStr := txtTpl
		for idxCol, headCol := range head {
			rawStr = strings.ReplaceAll(rawStr, "{"+headCol+"}", row[idxCol])
		}
		fmt.Printf("%s\n", rawStr)
	}
}
