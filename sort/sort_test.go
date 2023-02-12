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

func TestBubbleSort(t *testing.T) {
	fmt.Println(BubbleSort([]int{10, 4, 3, 6, 2, 1, 0}))
}

func TestChoseSort(t *testing.T) {
	val := int(^uint(0) >> 1)
	min := ^val
	fmt.Printf("%#v ; %T \n", val, val)
	fmt.Printf("%#v ; %T \n", min, min)
	fmt.Println(SelectionSort([]int{10, 4, 3, 6, 2, 1, 0}))
}

func TestInsertSort(t *testing.T) {
	fmt.Println(InsertSort([]int{10, 4, 3, 6, 2, 1, 0}))
}

func TestShellSort(t *testing.T) {
	fmt.Println(ShellSort([]int{10, 4, 3, 6, 2, 1, 0}))
}

func TestHeapSort(t *testing.T) {
	fmt.Println(HeapSort([]int{10, 4, 3, 6, 2, 1, 0}))
}

func TestCountSort(t *testing.T) {
	fmt.Println(CountSort([]int{10, 4, 3, 6, 2, 1, 0}))
}
