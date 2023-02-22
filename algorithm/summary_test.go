package main

import (
	"fmt"
	"testing"
)

func TestTen2N(t *testing.T) {
	fmt.Println(Ten2N(55, 16))
	fmt.Println(Ten2N(42, 2))
}

func TestTen2NDecimal(t *testing.T) {
	fmt.Println(Ten2NDecimal(0.930908203125, 8))
}

func TestN210(t *testing.T) {
	fmt.Println(N210("1A", 16))
}

func TestNNMoveNinetyDegree(t *testing.T) {
	fmt.Println(NNMoveNinetyDegree([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
}
