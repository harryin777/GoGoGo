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
	bookName = "窄门"
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
	re := regexp.MustCompile("(.*)\\r\\n(.*)\\r\\n\\r\\n(.*)")
	for _, part := range contentParts {
		if len(bookName) != 0 {
			match := re.FindStringSubmatch(part)
			if len(match) > 1 {
				if strings.Contains(match[0], bookName) {
					fmt.Println(match[3])
					fmt.Println()
				}
			}
		}
	}
}
