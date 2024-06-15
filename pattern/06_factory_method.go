package pattern

import "fmt"

/*
Factory Method pattern is a creational pattern that provides an interface for creating objects in a superclass, and
allows subclasses to alter the type of objects that will be created.

Pros:
    - Provides a way to encapsulate object creation
    - Promotes code reuse and decouples client code from specific implementations
Cons:
    - Can increase the complexity of the code with many classes and methods
*/

// Example

// Define 'Product' interface
type Product interface {
	Use() string
}

// ConcreteProductA struct
type ConcreteProductA struct{}

// Implement 'Use' method for 'ConcreteProductA'
func (p *ConcreteProductA) Use() string {
	return "Using Product A"
}

// ConcreteProductB struct
type ConcreteProductB struct{}

// Implement 'Use' method for 'ConcreteProductB'
func (p *ConcreteProductB) Use() string {
	return "Using Product B"
}

// Creator interface defines the factory method
type Creator interface {
	CreateProduct() Product
}

// ConcreteCreatorA struct
type ConcreteCreatorA struct{}

// Implement 'CreateProduct' method for 'ConcreteCreatorA'
func (c *ConcreteCreatorA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// ConcreteCreatorB struct
type ConcreteCreatorB struct{}

// Implement 'CreateProduct' method for 'ConcreteCreatorB'
func (c *ConcreteCreatorB) CreateProduct() Product {
	return &ConcreteProductB{}
}

// Example usage
func main() {
	var creator Creator

	// Create ConcreteProductA using ConcreteCreatorA
	creator = &ConcreteCreatorA{}
	productA := creator.CreateProduct()
	fmt.Println(productA.Use())

	// Create ConcreteProductB using ConcreteCreatorB
	creator = &ConcreteCreatorB{}
	productB := creator.CreateProduct()
	fmt.Println(productB.Use())
}
