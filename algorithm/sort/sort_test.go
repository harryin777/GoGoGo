package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort1(t *testing.T) {
	fmt.Println(MergeSort1([]int{10, 4, 3, 6, 2, 1}))
	fmt.Println(Merge1([]int{10, 4, 3, 6, 2, 1}))
}

func TestQuickSort(t *testing.T) {
	//fmt.Println(QuickSort([]int{4, 2, 1, 3}, 0, 3))
	//fmt.Println(QuickSort([]int{1, 3, 2}, 0, 2))
	fmt.Println(QuickSort([]int{10, 4, 3, 6, 2, 1, 0}, 0, 6))
	fmt.Println(Quick2([]int{10, 4, 3, 6, 2, 1, 0}, 0, 6))
}

func TestBubbleSort(t *testing.T) {
	fmt.Println(BubbleSort([]int{10, 4, 3, 6, 2, 1, 0}))
}

func TestSelectionSort(t *testing.T) {
	fmt.Println(SelectionSort([]int{10, 4, 3, 6, 2, 1, 0}))
	fmt.Println(Selection2([]int{10, 4, 3, 6, 2, 1, 0}))

}

func TestInsertSort(t *testing.T) {
	fmt.Println(InsertSort([]int{10, 4, 3, 6, 2, 1, 0}))
	fmt.Println(Insert2([]int{10, 4, 3, 6, 2, 1, 0}))
}

func TestShellSort(t *testing.T) {
	fmt.Println(ShellSort([]int{10, 4, 3, 6, 2, 1, 0}))
}

func TestHeapSort(t *testing.T) {
	fmt.Println(HeapSort([]int{10, 4, 3, 6, 2, 7, 0}))
	fmt.Println(HeapSort2([]int{10, 4, 3, 6, 2, 7, 0}))
}

func TestCountSort(t *testing.T) {
	fmt.Println(CountSort([]int{10, 4, 3, 6, 2, 1, 0}))
	fmt.Println(CountSort2([]int{10, 4, 3, 6, 2, 1, 0}))
	fmt.Println(CountSort3([]int{10, 4, 3, 6, 2, 1, 0}))
}

func TestMultiMergeSort(t *testing.T) {
	MultiMergeSort()
}
