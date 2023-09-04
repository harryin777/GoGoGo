package main

func main() {
	a := getInfo()
	_ = a
}

func getInfo() (res *int) {
	var i int
	i = 1
	res = &i
	return
}
