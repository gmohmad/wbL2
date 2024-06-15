package pattern

import "fmt"

/*
State pattern is a behavioral pattern that allows an object to alter its behavior when its internal state changes.
The object will appear to change its class.

Pros:
    - Localizes state-specific behavior and partitions behavior for different states
    - Makes state transitions explicit
Cons:
    - Can result in a large number of classes for different states
*/

// Example

// Define 'State' interface with 'Handle' method
type State interface {
	Handle(context *StateContext)
}

// ConcreteStateA struct
type ConcreteStateA struct{}

// Implement 'Handle' method for 'ConcreteStateA'
func (s *ConcreteStateA) Handle(context *StateContext) {
	fmt.Println("Handling request in State A")
	context.SetState(&ConcreteStateB{})
}

// ConcreteStateB struct
type ConcreteStateB struct{}

// Implement 'Handle' method for 'ConcreteStateB'
func (s *ConcreteStateB) Handle(context *StateContext) {
	fmt.Println("Handling request in State B")
	context.SetState(&ConcreteStateA{})
}

// StateContext struct that uses a state
type StateContext struct {
	state State
}

// SetState method for setting the state
func (c *StateContext) SetState(state State) {
	c.state = state
}

// Request method to handle request based on the current state
func (c *StateContext) Request() {
	c.state.Handle(c)
}

// Example usage
func main() {
	// Create context with initial state
	context := &StateContext{state: &ConcreteStateA{}}

	// Handle request, changing state each time
	context.Request() // Outputs: Handling request in State A
	context.Request() // Outputs: Handling request in State B
	context.Request() // Outputs: Handling request in State A
	context.Request() // Outputs: Handling request in State B
}
