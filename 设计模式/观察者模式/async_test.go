package observer

import (
	"fmt"
	"testing"
	"time"
)

func sub1(args ...interface{}) {
	time.Sleep(1 * time.Microsecond)
	fmt.Printf("sub1, %s %s\n", args[0], args[1])
}

func sub2(args ...interface{}) {
	fmt.Printf("sub2, %s %s\n", args[0], args[1])
}
func TestAsyncEvent(t *testing.T) {
	bus := NewEvent()
	bus.Subscribe("topic:1", sub1)
	bus.Subscribe("topic:1", sub2)
	bus.Publish("topic:1", "test1", "test2")
	bus.Publish("topic:1", "testA", "testB")
	time.Sleep(1 * time.Second)
}