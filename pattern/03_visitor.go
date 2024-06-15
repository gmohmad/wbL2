package pattern

import "fmt"

/*
Visitor pattern is a behavioral pattern that allows adding further operations to objects without modifying them.
It lets you define a new operation without changing the classes of the elements on which it operates.

Pros:
    - Makes it easier to add new operations
    - Adheres to the open/closed principle
Cons:
    - Can be difficult to implement if the object structure is complex
*/

// Example

// Define 'Visitor' interface
type Visitor interface {
	VisitElementA(*ElementA)
	VisitElementB(*ElementB)
}

// Define 'ElementA' struct
type ElementA struct {
	name string
}

// Implement 'Accept' method for 'ElementA'
func (e *ElementA) Accept(v Visitor) {
	v.VisitElementA(e)
}

// Define 'ElementB' struct
type ElementB struct {
	name string
}

// Implement 'Accept' method for 'ElementB'
func (e *ElementB) Accept(v Visitor) {
	v.VisitElementB(e)
}

// Define 'ConcreteVisitor' struct that implements 'Visitor' interface
type ConcreteVisitor struct{}

// Implement 'VisitElementA' method for 'ConcreteVisitor'
func (v *ConcreteVisitor) VisitElementA(e *ElementA) {
	fmt.Printf("Visiting %s in ElementA\n", e.name)
}

// Implement 'VisitElementB' method for 'ConcreteVisitor'
func (v *ConcreteVisitor) VisitElementB(e *ElementB) {
	fmt.Printf("Visiting %s in ElementB\n", e.name)
}

// Example usage
func main() {
	// Create elements
	elemA := &ElementA{name: "Element A"}
	elemB := &ElementB{name: "Element B"}

	// Create visitor
	visitor := &ConcreteVisitor{}

	// Accept visitor for each element
	elemA.Accept(visitor)
	elemB.Accept(visitor)
}
