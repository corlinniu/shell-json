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
	//if len(os.Args) < 3 {
	//	_, _ = os.Stderr.WriteString("USAGE: json <SEP> <JSONPATH> ... \n")
	//	return
	//}
	//var paths = os.Args[2:]
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
			fmt.Printf("res = %v\n", res)
			rstStr = strings.ReplaceAll(rstStr, "{"+vr+"}", fmt.Sprintf("%v", res))
		}
		fmt.Printf(rstStr)

		//var rst []string
		//for _, path := range paths {
		//	if path[0] != '$' {
		//		rst = append(rst, path)
		//		continue
		//	}
		//
		//	res, _ := jsonpath.JsonPathLookup(jsonData, path)
		//	if res == nil {
		//		rst = append(rst, "")
		//	} else {
		//		rst = append(rst, fmt.Sprintf("%v", res))
		//	}
		//}
		//fmt.Printf("%s\n", strings.Join(rst, os.Args[1]))
	}
}
