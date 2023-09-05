package main

type ChainHandler struct {
	handlers []Handler
}

func (c *ChainHandler) AddHandler(handler Handler) {
	c.handlers = append(c.handlers, handler)
}

func (c *ChainHandler) Handle(score int) {
	for _, handler := range c.handlers {
		handler.Handle(score)
	}
}
