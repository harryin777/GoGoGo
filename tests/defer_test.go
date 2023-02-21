package tests

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	fmt.Println(ff1())
	fmt.Println(ff2())
}

func ff1() int {
	var i int
	defer func() {
		i++
		fmt.Printf("defer 1: %v \n", i)
	}()

	defer func() {
		i++
		fmt.Printf("defer 2: %v \n", i)
	}()

	return i
}

func ff2() (i int) {
	defer func() {
		i++
		fmt.Printf("defer 1: %v \n", i)
	}()

	defer func() {
		i++
		fmt.Printf("defer 2: %v \n", i)
	}()

	return i
}

func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func TestClosePackage(t *testing.T) {
	in := Increase()
	fmt.Println(in())
	fmt.Println(in())
}
