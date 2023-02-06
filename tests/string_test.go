package tests

import (
	"strings"
	"testing"
)

// str builder 拼接字符串
func BenchmarkStringBuilder(*testing.B) {
	var strSlice []string
	for i := 0; i < 200; i++ {
		strSlice = append(strSlice, "a")
	}
	var b strings.Builder
	b.Grow(len(strSlice))
	for _, val := range strSlice {
		b.WriteString(val)
	}
}
