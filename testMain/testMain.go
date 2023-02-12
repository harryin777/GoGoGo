package main

import (
	"fmt"
)

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		fmt.Printf("here : %v \n", i)
		ch <- vs[i](name)
	}
	fmt.Println(len(vs))
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}

func main() {
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})

	fmt.Println(ret)

}
