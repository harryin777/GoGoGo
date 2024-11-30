package main

import (
	"fmt"
)

type B struct {
	C msg
}

func (b *B) ReadMsg(str string) string {
	fmt.Printf("this is b : %v \n", str)

	return str
}

type C struct {
}

func (c *C) ReadMsg(str string) string {
	fmt.Printf("this is c : %v \n", str)

	return str
}
