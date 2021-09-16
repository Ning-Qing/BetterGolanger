package observer

import (
	"fmt"
	"sync"
)

type Handler func(args ...interface{})

type Observer interface {
	Subscribe(topic string, handler Handler)
	Publish(topic string, args ...interface{})
}

var _ Observer = (*Event)(nil)

type Event struct {
	subject map[string][]Handler
	lock    sync.Mutex
}

func NewEvent() *Event {
	return &Event{
		subject: make(map[string][]Handler),
		lock:    sync.Mutex{},
	}
}

func (e *Event) Subscribe(topic string, handler Handler) {
	e.lock.Lock()
	defer e.lock.Unlock()

	_, ok := e.subject[topic]
	if !ok {
		e.subject[topic] = make([]Handler, 0)
	}
  	e.subject[topic] = append(e.subject[topic], handler)
}

func (e *Event) Publish(topic string, args ...interface{}) {
	handlers, ok := e.subject[topic]
	if !ok {
		return
	}

	for _, handler := range handlers {
		go handler(args...)
	}
}
