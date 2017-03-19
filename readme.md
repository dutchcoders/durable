# Durable Channel

The durable channel allows channels to persist on disk, and being used as normal channels. These channels are limited by storage only.

# Use cases

* sending data with unreliable connections 

# Sample

```
package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
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
        item := <- c
        fmt.Printf("%#v\n", item)
    }
}

```

# Disclaimer

The durable channel is based on the disk queue of [nsqd](https://github.com/nsqio/nsq/blob/master/nsqd/diskqueue.go).


