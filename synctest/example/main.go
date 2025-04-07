package main

import (
	"fmt"
	"time"
	"sync"
	"context"
)

type EventProcessor struct {
	events    chan string
	processed map[string]bool
	timedOut  map[string]bool
	timeout   time.Duration
	mu        sync.Mutex
}

func NewEventProcessor(bufferSize int, timeout time.Duration) *EventProcessor {
	return &EventProcessor{
		events:    make(chan string, bufferSize),
		processed: make(map[string]bool),
		timedOut:  make(map[string]bool),
		timeout:   timeout,
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
	ctx, cancel := context.WithTimeout(context.Background(), ep.timeout)
	defer cancel()

	// struct contain 0 bytes.
	doneProcessing := make(chan struct{})

	go func() {
		// it just for looking the behavior of the timeout process in test.
		if event != "slow" {
			ep.mu.Lock()
			ep.processed[event] = true
			ep.mu.Unlock()
			doneProcessing <- struct{}{}
		}
	}()

	select {
	case <-doneProcessing:
		return
	case <-ctx.Done():
		ep.mu.Lock()
		ep.timedOut[event] = true
		ep.mu.Unlock()
		return
	}
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

func (ep *EventProcessor) GetTimedOutEvents() map[string]bool {
	ep.mu.Lock()
	defer ep.mu.Unlock()

	result := make(map[string]bool)
	for event, timedOut := range ep.timedOut {
		result[event] = timedOut
	}
	return result
}

func main() {
	processor := NewEventProcessor(10, 5 * time.Second)
	processor.Start()

	processor.Process("event1")
	processor.Process("event2")
	processor.Process("event3")
	processor.Process("slow")
	
	time.Sleep(3 * time.Second)

	processor.Stop()

	processed := processor.GetProcessedEvents()
	timedOut := processor.GetTimedOutEvents()

	if !processed["event1"] {
		fmt.Println("event1 was not processed")
	}
	if !processed["event2"] {
		fmt.Println("event2 was not processed")
	}
	if !processed["event3"] {
		fmt.Println("event3 was not processed")
	}

	if !timedOut["slow"] {
		fmt.Println("slow was not timed out")
	}

	if len(processed) != 3 {
		fmt.Println("processed events count is not 3")
	}
}
