package main

type Handler interface {
	Handle(score int)
}

type HandlerOne struct {
	//Handler
}

func (h *HandlerOne) Handle(score int) {
	if score <= 10 {
		println("HandlerOne")
		return
	}
}

type HandlerTwo struct {
	//Handler
}

func (h *HandlerTwo) Handle(score int) {
	if score > 10 && score <= 20 {
		println("HandlerTwo")
		return
	}
}

type HandlerThree struct {
	//Handler
}

func (h *HandlerThree) Handle(score int) {
	if score > 20 {
		println("HandlerThree")
		return
	}
}

func main() {
	c := ChainHandler{}
	c.AddHandler(&HandlerOne{})
	c.AddHandler(&HandlerTwo{})
	c.AddHandler(&HandlerThree{})
	c.Handle(10)
}
