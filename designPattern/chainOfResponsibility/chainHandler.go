package main

import "fmt"

type ChainHandler struct {
	handlers []Handler
}

func (c ChainHandler) AddHandler(handler Handler) {
	c.handlers = append(c.handlers, handler)
}

func (c ChainHandler) Handle(score int) {
	for _, handler := range c.handlers {
		handler.Handle(score)
	}
}

type ee struct {
	aa []int
}

func (e ee) AddHandler(a int) {
	e.aa = append(e.aa, a)
}

func main() {
	e := ee{}
	e.AddHandler(1)
	fmt.Print(e.aa)
}
