package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/oliveagle/jsonpath"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: cat raw.txt | json2 'update User set name = \"{name}\" where id = {id}'")
		return
	}

	fileInfo, _ := os.Stdin.Stat()
	if (fileInfo.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
		_, _ = os.Stderr.WriteString("元数据必须通过管道输入\n")
		return
	}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var stdIn = s.Bytes()
		var jsonData interface{}
		err := json.Unmarshal(stdIn, &jsonData)
		if err != nil {
			_, _ = os.Stderr.WriteString("ERR: 输入数据非Json格式\n")
			continue
		}

		rstStr := os.Args[1]
		rg := regexp.MustCompile(`{[^}]+}`)
		vars := rg.FindAllString(os.Args[1], -1)
		for _, vr := range vars {
			vr = vr[1 : len(vr)-1]
			res, _ := jsonpath.JsonPathLookup(jsonData, "$."+vr)
			rstStr = strings.ReplaceAll(rstStr, "{"+vr+"}", fmt.Sprintf("%v", res))
		}
		fmt.Printf(rstStr)
	}
}
