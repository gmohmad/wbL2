package pattern

import "fmt"

/*
Chain of Responsibility pattern is a behavioral pattern that allows passing requests along a chain of handlers. 
When receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.

Pros:
    - Decouples sender and receiver
    - Adds flexibility in assigning responsibilities to objects
Cons:
    - Can result in a complex chain structure
*/

// Example

// Define 'Handler' interface with 'SetNext' and 'Handle' methods
type Handler interface {
	SetNext(Handler) Handler
	Handle(request string)
}

// BaseHandler struct to implement common functionality
type BaseHandler struct {
	next Handler
}

// Implement 'SetNext' method for 'BaseHandler'
func (h *BaseHandler) SetNext(next Handler) Handler {
	h.next = next
	return next
}

// ConcreteHandler1 struct
type ConcreteHandler1 struct {
	BaseHandler
}

// Implement 'Handle' method for 'ConcreteHandler1'
func (h *ConcreteHandler1) Handle(request string) {
	if request == "Request1" {
		fmt.Println("ConcreteHandler1 handled the request")
	} else if h.next != nil {
		h.next.Handle(request)
	}
}

// ConcreteHandler2 struct
type ConcreteHandler2 struct {
	BaseHandler
}

// Implement 'Handle' method for 'ConcreteHandler2'
func (h *ConcreteHandler2) Handle(request string) {
	if request == "Request2" {
		fmt.Println("ConcreteHandler2 handled the request")
	} else if h.next != nil {
		h.next.Handle(request)
	}
}

// ConcreteHandler3 struct
type ConcreteHandler3 struct {
	BaseHandler
}

// Implement 'Handle' method for 'ConcreteHandler3'
func (h *ConcreteHandler3) Handle(request string) {
	if request == "Request3" {
		fmt.Println("ConcreteHandler3 handled the request")
	} else if h.next != nil {
		h.next.Handle(request)
	}
}

// Example usage
func main() {
	// Create handlers
	handler1 := &ConcreteHandler1{}
	handler2 := &ConcreteHandler2{}
	handler3 := &ConcreteHandler3{}

	// Set up chain of responsibility
	handler1.SetNext(handler2).SetNext(handler3)

	// Send requests
	requests := []string{"Request1", "Request2", "Request3", "UnknownRequest"}
	for _, request := range requests {
		handler1.Handle(request)
	}
}
