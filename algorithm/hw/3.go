package main

import "fmt"

func main3() {
	var a, b int
	fmt.Scan(&a, &b)

	if a*2%b != 0 {
		fmt.Println(-1)
		return
	}

	if (a*2/b-b)%2 == 0 {
		fmt.Println(-1)
		return
	}

	num := (a*2/b + 1 - b) / 2
	for i := 0; i < b; i++ {
		fmt.Printf("%d ", i+num)
	}
}
