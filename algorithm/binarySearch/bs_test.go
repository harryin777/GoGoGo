package binarySearch

import (
	"fmt"
	"testing"
)

func TestLeftBound(t *testing.T) {
	//fmt.Println(LeftBound([]int{0, 1, 2, 3, 5, 6, 7}, 8))
	//fmt.Println(leftBound([]int{0, 1, 2, 3, 5, 6, 7}, 8))
	fmt.Println(RightBound([]int{1, 2, 3, 3, 3, 5, 6, 7}, 3))
	fmt.Println(rightBound([]int{1, 2, 3, 3, 3, 5, 6, 7}, 3))
}

func TestTripleOne(t *testing.T) {
	TripleOne(2)
}

func TestNormalBSVSBS(t *testing.T) {
	fmt.Println(BS([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 3))
	fmt.Println(BS2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 3))
}
