package main

import (
	"fmt"
	"time"
)

/*
=== Or channel ===

Implement a function that will merge one or more done channels into a single channel if one of its component channels closes.
One option would obviously be to write a select expression that implements this relationship,
however, sometimes the total number of done channels you are working with at runtime is unknown.
In this case, it is more convenient to use a call to a single function, which, having received one or more or channels as input, implements all the functionality.

Function Definition:
var or func(channels ...<- chan interface{}) <- chan interface{}

Example of using the function:
sig := func(after time.Duration) <- chan interface{} {
  c := make(chan interface{})
  go func() {
  defer close(c)
  time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
  sig(2*time.Hour),
  sig(5*time.Minute),
  sig(1*time.Second),
  sig(1*time.Hour),
  sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

// or function takes multiple channels as input and returns a single channel that closes when any of the input channels close.
func or(channels ...<-chan interface{}) <-chan interface{} {
	// Handle edge cases where no channels or only one channel is provided
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	// Create a channel to signal when any of the provided channels close
	orDone := make(chan interface{})
	go func() {
		defer close(orDone) // Ensure orDone is closed when the goroutine exits

		// Handle the case where there are exactly two channels
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			// If there are more than two channels, use recursion to handle them
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...): // Recursively handle the remaining channels
			}
		}
	}()
	return orDone
}

// sig function returns a channel that closes after the specified duration
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)    // Ensure the channel is closed when the goroutine exits
		time.Sleep(after) // Sleep for the specified duration
	}()
	return c
}

func main() {

	start := time.Now()
	// Use the or function to wait for any of the provided channels to close
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v\n", time.Since(start)) // Print the time elapsed since start
}
