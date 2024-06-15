package pattern

import "fmt"

/*
Strategy pattern is a behavioral pattern that enables selecting an algorithm's behavior at runtime.
It defines a family of algorithms, encapsulates each one, and makes them interchangeable.

Pros:
    - Allows selecting algorithms at runtime
    - Promotes the open/closed principle
Cons:
    - Clients must be aware of different strategies
*/

// Example

// Define 'Strategy' interface with 'Execute' method
type Strategy interface {
	Execute() string
}

// ConcreteStrategyA struct
type ConcreteStrategyA struct{}

// Implement 'Execute' method for 'ConcreteStrategyA'
func (s *ConcreteStrategyA) Execute() string {
	return "Executing Strategy A"
}

// ConcreteStrategyB struct
type ConcreteStrategyB struct{}

// Implement 'Execute' method for 'ConcreteStrategyB'
func (s *ConcreteStrategyB) Execute() string {
	return "Executing Strategy B"
}

// Context struct that uses a strategy
type Context struct {
	strategy Strategy
}

// SetStrategy method for setting the strategy at runtime
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

// ExecuteStrategy method to execute the strategy's algorithm
func (c *Context) ExecuteStrategy() string {
	return c.strategy.Execute()
}

// Example usage
func main() {
	// Create context
	context := &Context{}

	// Use ConcreteStrategyA
	context.SetStrategy(&ConcreteStrategyA{})
	fmt.Println(context.ExecuteStrategy())

	// Use ConcreteStrategyB
	context.SetStrategy(&ConcreteStrategyB{})
	fmt.Println(context.ExecuteStrategy())
}
