package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	lineParts := strings.Split(scanner.Text(), "~")
	//	fmt.Println(lineParts)
	//}
	//
	//fmt.Println("gggg")
	buf := bufio.NewReader(os.Stdin)
	for {
		str, err := buf.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println(str)
	}
}
