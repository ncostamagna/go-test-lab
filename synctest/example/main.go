package main

import (
	"fmt"
	"time"
	"sync"
)

type EventProcessor struct {
	events    chan string
	processed map[string]bool
	mu        sync.Mutex
}

func NewEventProcessor(bufferSize int) *EventProcessor {
	return &EventProcessor{
		events:    make(chan string, bufferSize),
		processed: make(map[string]bool),
	}
}

func (ep *EventProcessor) Start() {
	go func() {
		for event := range ep.events {
			go ep.handleEvent(event) // no guarantee, possible solution is to use sync.WaitGroup
		}
	}()
}

func (ep *EventProcessor) Stop() {
	close(ep.events)
}

func (ep *EventProcessor) Process(event string) {
	ep.events <- event
}

func (ep *EventProcessor) handleEvent(event string) {
	ep.mu.Lock()
	ep.processed[event] = true
	ep.mu.Unlock()
}

func (ep *EventProcessor) GetProcessedEvents() map[string]bool {
	ep.mu.Lock()
	defer ep.mu.Unlock()

	result := make(map[string]bool)
	for event, processed := range ep.processed {
		result[event] = processed
	}
	return result
}

func main() {
	processor := NewEventProcessor(10)
	processor.Start()

	processor.Process("event1")
	processor.Process("event2")
	processor.Process("event3")

	time.Sleep(1 * time.Second)

	processor.Stop()

	processed := processor.GetProcessedEvents()

	if !processed["event1"] {
		fmt.Println("event1 was not processed")
	}
	if !processed["event2"] {
		fmt.Println("event2 was not processed")
	}
	if !processed["event3"] {
		fmt.Println("event3 was not processed")
	}

	if len(processed) != 3 {
		fmt.Println("processed events count is not 3")
	}
}
