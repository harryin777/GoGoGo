package observer

type ISubject interface {
	Register(observer IObserver) error
	CancelRegister(observer IObserver) error
	Publish() error
}

type Subject struct {
	RegisterChannels map[string]chan string
}

func (s *Subject) Register(observer *Observer) error {
	observer.ListenChannel = make(chan string)
	s.RegisterChannels[observer.ObserverName] = observer.ListenChannel
	return nil
}

func (s *Subject) CancelRegister(observer Observer) (res bool, err error) {
	if _, open := <-observer.ListenChannel; open {
		return false, nil
	}

	delete(s.RegisterChannels, observer.ObserverName)
	return true, nil
}

func (s *Subject) Publish(msg string) error {
	for _, val := range s.RegisterChannels {
		go func() {
			val <- msg
		}()
	}
	return nil
}
