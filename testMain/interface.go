package main

type TTT interface {
	Say(str string)
}

func NewImpl() TTT {
	return &Impl{
		a: "123",
	}
}
