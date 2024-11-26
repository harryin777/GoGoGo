package main

import "fmt"

type user struct {
	name string
}

var u = &user{
	name: "111",
}

func Set(status int) func() {
	update := func(info *user) {
		if status == 1 {
			info.name = "1"
		} else {
			info.name = "2"
		}
	}

	return func() {
		update(u)
	}
}

func e(update func(*user)) {
	update(u)
}

func main() {
	Set(1)()

	fmt.Println(u.name)
}
