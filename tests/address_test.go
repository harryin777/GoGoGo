package tests

import (
	"fmt"
	"testing"
)

func Test_Address(t *testing.T) {
	fmt.Println(*receiveData())
}

func changeData() (res *string) {
	text := "123"
	res = &text
	return
}

func receiveData() (res *string) {
	res = changeData()
	return
}

func Test_Address2(t *testing.T) {
	var data2 *string
	data2 = changeData2()
	fmt.Println(*data2)
}

func changeData2() *string {
	data := "hello"
	return &data
}
