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
	bookName = " 孤"
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
			part2 := []rune(part)[0:2]
			bookName2 := []rune(bookName)[0:2]
			fmt.Println(string(part2))
			fmt.Println(string(bookName2))
			fmt.Println(string(part2) == bookName)
			part = strings.Trim(part, "\n")
			if strings.HasPrefix(part, bookName) {
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
	//fmt.Print(string(bytes))
}
