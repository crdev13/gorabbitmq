package main

import (
	"time"

	setup "github.com/crdev13/gorabbitmq/example-1"
)

func main() {
	go setup.RunProducer()

	time.Sleep(3 * time.Second)
	setup.RunConsumer()
}
