package tests

import "fmt"

func a() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover from panic")
		}
	}()
	fmt.Println("this is a")
	b()
}

func b() {
	fmt.Println("this is b")
	c()
}

func c() {
	fmt.Println("this is c")
	panic("c test")
}
