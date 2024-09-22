// concurrent3
// Make the tests pass!

package main_test

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSendAndReceive(t *testing.T) {
	var buf bytes.Buffer

	messages := make(chan string)
	sendAndReceive(&buf, messages)

	got := buf.String()
	want := "Hello World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func sendAndReceive(buf *bytes.Buffer, messages chan string) {
	go func() {
		messages <- "Hello"
		messages <- "World"
		close(messages)
	}()
	
	firstSpace := false

	for greeting := range messages {
		fmt.Fprint(buf, greeting)
		if(!firstSpace) {
			fmt.Fprint(buf, " ")
			firstSpace = true
		}
	}
/*
	// Here we just receive the first message
	// Consider using a for-range loop to iterate over the messages
	_, ok := <-messages
	if !ok {
		fmt.Fprint(buf, "Channel is closed")
	}
	*/
}

// I don't think this is how you solve it because it seems like the status should mean something,
// however, no matter what i do, it keeps staying false and i cannot stop it from writing to the buffer
// again, remember the channel is already closed
