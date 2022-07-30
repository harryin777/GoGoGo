package observer

import (
	"fmt"
	"testing"
	"time"
)

//channel 一定在使用之前make
func Test_Sub(t *testing.T) {
	s1 := Subject{}
	o1 := Observer{}
	s1.RegisterChannels = make(map[string]chan string)
	err := s1.Register(&o1)
	if err != nil {
		return
	}
	go func() {
		o1.Subscribe()
		fmt.Println(1)
	}()
	fmt.Println(2)
	err = s1.Publish("hahaha")
	if err != nil {
		return
	}
	time.Sleep(10000)
}
