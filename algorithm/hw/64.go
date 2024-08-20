package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
输入两个字符串s和，都只包含小写字母，1en(S)<=1/
的有效子字符串。
1en(L)<=599999。刂斤S是否疋L
．判定规贝刂：s中的每个字符在L中都能找到（可以不连续），且s在中字符的前后顺序与s中
顺序要保持一致。
例如：
ace
abcde

注意这个题，如果不完全匹配的话不是输出 -1 而是最后一个匹配的字符在t中的位置
*/

func main64() {
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	s := buf.Text()
	buf.Scan()
	t := buf.Text()
	cal64(s, t)
}

func cal64(s, t string) {
	i, j, lashMatch := 0, 0, -1
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			lashMatch = j
			i++
			j++
		} else {
			j++
		}
	}

	fmt.Println(lashMatch)
}
