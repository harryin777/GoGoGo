package bitMap

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	bm := NewBitMap(100000000)
	count := 8888889
	bm.Add(uint(count))
	fmt.Println(bm.IsExist(uint(count)))
}
