package main

import (
	"testing"
	"testing/synctest"
	"time"
)

func TestEventProcessorTimeout(t *testing.T) {
	// this Run creates the bubble
	synctest.Run(func() {
		processor := NewEventProcessor(10, 100 * time.Minute)
		processor.Start()

		processor.Process("event1")
		processor.Process("event2")
		processor.Process("event3")
		processor.Process("slow")

		time.Sleep(90 * time.Minute)

		processor.Stop()

		//synctest.Wait() // wait for the all goroutines to finish in the bubble

		processed := processor.GetProcessedEvents()
		timedOut := processor.GetTimedOutEvents()

		if !processed["event1"] {
			t.Errorf("event1 was not processed")
		}
		if !processed["event2"] {
			t.Errorf("event2 was not processed")
		}
		if !processed["event3"] {
			t.Errorf("event3 was not processed")
		}
		if !timedOut["slow"] {
			t.Errorf("slow was not timed out")
		}


		if len(processed) != 3 {
			t.Errorf("processed events count is not 3")
		}
	})
}