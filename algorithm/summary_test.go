package main

import (
	"fmt"
	"testing"
)

func TestTen2N(t *testing.T) {
	fmt.Println(DecimalToAny(55, 32))
	fmt.Println(N210("1n", 32))
}

func TestTen2NDecimal(t *testing.T) {
	fmt.Println(Ten2NDecimal(0.930908203125, 8))
}

func TestN210(t *testing.T) {
	fmt.Println(N210("1n", 10))
}

func TestNNMoveNinetyDegree(t *testing.T) {
	fmt.Println(NNMoveNinetyDegree([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
}

func TestSquareNums(t *testing.T) {
	SquareNums(200)
}

func TestIptoIn(t *testing.T) {
	fmt.Println(IptoIn("100.101.1.5"))
}

func TestShuiXianHua(t *testing.T) {
	for i := 0; i < 100000; i++ {

	}
}
