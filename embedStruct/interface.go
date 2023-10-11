package embedStruct

import "fmt"

type A interface {
	Read(str string)
}

type B struct {
	Str string `json:"str"`
}

func (b *B) Read(str string) {
	fmt.Printf("b is reading %v \n", str)
}

type C struct {
	B
}

type D struct {
	B B
}
