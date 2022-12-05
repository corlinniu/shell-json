package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: csv file-to-path.csv 'The name is {Name}.'\n")
		return
	}

	fileName := os.Args[1]
	txtTpl := os.Args[2]

	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	rst := CsvToMap(string(fileContent))

	var usedKeys []string
	rg := regexp.MustCompile(`{[^}]+}`)
	vars := rg.FindAllString(txtTpl, -1)
	for _, vr := range vars {
		vr = vr[1 : len(vr)-1]
		usedKeys = append(usedKeys, vr)
	}

	for _, row := range rst {
		rawStr := txtTpl
		for _, key := range usedKeys {
			rawStr = strings.ReplaceAll(rawStr, "{"+key+"}", row[key])
		}
		fmt.Printf("%s\n", rawStr)
	}
}

func CsvToMap(fileContent string) []map[string]string {
	rows := strings.Split(fileContent, "\n")

	var head []string
	for _, s := range strings.Split(rows[0], ",") {
		head = append(head, s)
	}

	var rst []map[string]string
	for _, row := range rows[1:] {
		cols := strings.Split(row, ",")

		var rowMap = make(map[string]string)
		for idx, col := range cols {
			rowMap[head[idx]] = col
		}

		rst = append(rst, rowMap)
	}

	return rst
}
