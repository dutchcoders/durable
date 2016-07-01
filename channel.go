package durable

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type Config struct {
	Name            string
	DataPath        string
	MaxBytesPerFile int64
	MinMsgSize      int32
	MaxMsgSize      int32
	SyncEvery       int64
	SyncTimeout     time.Duration
	Logger          *log.Logger
}

func defaultConfig() *Config {
	return &Config{
		Name:            "",
		DataPath:        "./data",
		MaxBytesPerFile: 102400,
		MinMsgSize:      0,
		MaxMsgSize:      1000,
		SyncEvery:       10,
		SyncTimeout:     time.Second * 10,
		Logger:          log.New(ioutil.Discard, "", 0),
	}
}

type channel struct {
	in     chan interface{}
	out    chan interface{}
	dq     *diskQueue
	config *Config
}

func newChannel(c chan interface{}, config *Config) chan interface{} {
	out := make(chan interface{})

	b := channel{
		in:     c,
		out:    out,
		config: config,
	}

	b.dq = newDiskQueue(config)

	go b.reader()
	go b.writer()

	return out
}

func (b channel) reader() {
	for data := range b.dq.ReadChan() {
		var item interface{}

		if err := json.Unmarshal(data, &item); err != nil {
			b.config.Logger.Printf("Error unmarshalling json object: %s\n", err.Error())
		}

		b.out <- item
	}
}

func (b channel) writer() {
	for {
		item := <-b.in

		if data, err := json.Marshal(item); err != nil {
			b.config.Logger.Printf("Error marshalling json object: %s\n", err.Error())
		} else if err := b.dq.Put(data); err != nil {
			b.config.Logger.Printf("Error putting object: %s\n", err.Error())
		}

	}
}

func Channel(c chan interface{}, config *Config) chan interface{} {
	if config == nil {
		config = defaultConfig()
	}

	return newChannel(c, config)
}
