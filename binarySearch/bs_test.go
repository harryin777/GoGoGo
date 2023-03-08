package binarySearch

import (
	"fmt"
	"testing"
)

func TestLeftBound(t *testing.T) {
	fmt.Println(LeftBound([]int{1, 2, 3, 3, 5, 6, 7}, 3))
	fmt.Println(RightBound([]int{1, 2, 3, 3, 5, 6, 7}, 3))
}
