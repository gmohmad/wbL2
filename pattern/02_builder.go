package pattern

import "fmt"

/*
Builder pattern is a creational pattern that allows you to construct complex objects step by step.
It separates the construction of a complex object from its representation, allowing the same construction process 
to create different representations.

Pros:
    - Allows step-by-step construction
Cons:
    - Can lead to a lot of boilerplate code
*/

// Example

// Define 'House' struct
type House struct {
	Windows string
	Doors   string
	Roof    string
}

// Define 'HouseBuilder' struct
type HouseBuilder struct {
	windows string
	doors   string
	roof    string
}

// Define 'NewHouseBuilder' function to return a new HouseBuilder
func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{}
}

// Define 'SetWindows' method on 'HouseBuilder' to set windows
func (b *HouseBuilder) SetWindows(windows string) *HouseBuilder {
	b.windows = windows
	return b
}

// Define 'SetDoors' method on 'HouseBuilder' to set doors
func (b *HouseBuilder) SetDoors(doors string) *HouseBuilder {
	b.doors = doors
	return b
}

// Define 'SetRoof' method on 'HouseBuilder' to set roof
func (b *HouseBuilder) SetRoof(roof string) *HouseBuilder {
	b.roof = roof
	return b
}

// Define 'Build' method on 'HouseBuilder' to construct the final 'House' object
func (b *HouseBuilder) Build() *House {
	return &House{
		Windows: b.windows,
		Doors:   b.doors,
		Roof:    b.roof,
	}
}

// Example usage
func main() {

	// Create a builder
	builder := NewHouseBuilder()

	// Create a house using the builder
	house := builder.SetWindows("Double Pane Windows").SetDoors("Wooden Doors").SetRoof("Gabled Roof").Build()

	// Print the house's fields
	fmt.Printf("House built with %s, %s, and %s\n", house.Windows, house.Doors, house.Roof)
}
