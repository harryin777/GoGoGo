package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort1(t *testing.T) {
	fmt.Println(MergeSort1([]int{10, 4, 3, 6, 2, 1}))
}

func TestQuickSort(t *testing.T) {
	fmt.Println(QuickSort([]int{10, 4, 3, 6, 2, 1, 0}, 0, 5))
	fmt.Println(string([]byte{1}))
}
