package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/oliveagle/jsonpath"
	"os"
	"strings"
)

func main() {
	var paths = os.Args[1:]
	fileInfo, _ := os.Stdin.Stat()
	if (fileInfo.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
		_, _ = os.Stderr.WriteString("元数据必须通过管道输入\n")
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

		var rst []string
		for _, path := range paths {
			res, _ := jsonpath.JsonPathLookup(jsonData, path)
			str, _ := json.Marshal(res)
			rst = append(rst, string(str))
		}
		fmt.Printf("%s\n", strings.Join(rst, " "))
	}
}
