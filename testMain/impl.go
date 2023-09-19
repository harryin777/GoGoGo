package main

import "fmt"

type Impl struct {
	a string
}

func (i *Impl) Say(str string) {
	fmt.Printf("here is say : %v, a : %v \n", str, i.a)
}
