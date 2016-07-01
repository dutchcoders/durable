package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dutchcoders/durable"
)

func main() {
	writer := make(chan interface{})
	c := durable.Channel(writer, durable.Config{
		Name:            "",
		DataPath:        "./data",
		MaxBytesPerFile: 102400,
		MinMsgSize:      0,
		MaxMsgSize:      1000,
		SyncEvery:       10,
		SyncTimeout:     time.Second * 10,
		Logger:          log.New(os.Stdout, "", 0),
	})

	for i := 0; i < 10000; i++ {
		writer <- fmt.Sprintf("%d", i)
	}

	for {
		item := <-c
		fmt.Printf("%#v\n", item)
	}
}
