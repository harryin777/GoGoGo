package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

const (
	path     = "F:\\documents\\My Clippings.txt"
	bookName = " 孤独小说家"
)

func main() {
	f, err := os.OpenFile(path, os.O_RDONLY, 0666) //打开文件
	if err != nil {
		panic(err)
	}
	b := bufio.NewReader(f)
	bytes, err := io.ReadAll(b)
	if err != nil {
		panic(err)
	}
	content := string(bytes)
	contentParts := strings.Split(content, "==========")
	if len(bookName) != 0 {
		fmt.Println(bookName)
	}
	re := regexp.MustCompile("\\r\\n\\r\\n(.*)")
	for _, part := range contentParts {
		if len(bookName) != 0 {
			nameList := []rune(bookName)[1:2]
			part2 := []rune(part)[1 : len(nameList)+1]
			var nameStr string
			var linePrefix string
			for _, name := range nameList {
				nameStr = nameStr + string(name)
			}
			for _, val := range part2 {
				linePrefix = linePrefix + string(val)
			}

			if nameStr == linePrefix {
				match := re.FindStringSubmatch(part)
				if len(match) > 1 {
					fmt.Println(match[1])
				}
			}
		} else {
			match := re.FindStringSubmatch(part)
			if len(match) > 1 {
				fmt.Println(match[1])
			}
		}
	}
}
