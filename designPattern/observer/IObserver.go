package observer

import "fmt"

type IObserver interface {
	Subscribe()
}

type Observer struct {
	ObserverName  string      `json:"observerName"`
	ListenChannel chan string `json:"listenChannel"`
}

func (observer Observer) Subscribe() {
	msg := <-observer.ListenChannel
	fmt.Println(msg)
}
