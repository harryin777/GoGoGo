package observer

import "reflect"

type ISubject interface {
	Register(observer IObserver) error
	CancelRegister(observer IObserver) error
	Publish() error
}

type Subject struct {
	RegisterChannels []chan string `json:"registerChannels"`
}

func (s *Subject) Register(observer *Observer) error {
	observer.ListenChannel = make(chan string)
	s.RegisterChannels = append(s.RegisterChannels, observer.ListenChannel)
	return nil
}

func (s *Subject) CancelRegister(observer Observer) error {
	for index, val := range s.RegisterChannels {
		if reflect.DeepEqual(val, observer.ListenChannel) {
			s.RegisterChannels = append(s.RegisterChannels[:index], s.RegisterChannels[index+1:]...)
		}
	}
	return nil
}

func (s *Subject) Publish(msg string) error {
	for i := 0; i < len(s.RegisterChannels); i++ {
		//这里一定要用go协程去发布消息，因为一旦有一个channel没有被监听到，所有其他的channel都不会被监听
		go func(i int) {
			s.RegisterChannels[i] <- msg
		}(i)

	}
	return nil
}
