package durable

import (
	"fmt"
	"testing"
)

func TestInfiniteChannel(t *testing.T) {
	c := make(chan interface{})

	d := Channel(c, nil)

	for i := 0; i < 1000; i++ {
		c <- fmt.Sprintf("%d", i)
	}

	for {
		item := <-d

		fmt.Printf("%#v\n", item)
	}
}
