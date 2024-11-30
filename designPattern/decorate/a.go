package main

import "fmt"

type msg interface {
	ReadMsg(string) string
}

type A struct {
	B msg
}

type MsgWrapper func(msg) msg

func (a *A) ReadMsg(str string) string {
	fmt.Printf("this is a : %v \n", str)
	if a.B != nil {
		a.B.ReadMsg(str)
	}
	return str
}

func NewAWrapper() MsgWrapper {
	return func(m msg) msg {
		return &A{
			B: m,
		}
	}
}

func NewBWrapper() MsgWrapper {
	return func(m msg) msg {
		return &B{
			C: m,
		}
	}
}

func main() {
	c := &C{}
	t := NewAWrapper()(NewBWrapper()(c))
	t.ReadMsg("1")
}
