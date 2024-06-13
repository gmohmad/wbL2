package pattern

import "fmt"

/*
Command pattern is a behavioral pattern that turns a request into a standalone object that contains all information 
about the request.

Pros:
	- Lets parameterize clients with different requests
    - Decouples the sender and receiver
    - Supports undo and redo operations
Cons:
    - Can result in a lot of small classes
*/

// Example

// Define 'Command' interface with 'Execute' method
type Command interface {
	Execute()
}

// Define 'Light' struct
type Light struct {
	isOn bool
}

// Define 'TurnOn' method for 'Light'
func (l *Light) TurnOn() {
	l.isOn = true
	fmt.Println("The light is on")
}

// Define 'TurnOff' method for 'Light'
func (l *Light) TurnOff() {
	l.isOn = false
	fmt.Println("The light is off")
}

// Define 'TurnOnCommand' struct
type TurnOnCommand struct {
	light *Light
}

// Implement 'Execute' method for 'TurnOnCommand'
func (c *TurnOnCommand) Execute() {
	c.light.TurnOn()
}

// Define 'TurnOffCommand' struct
type TurnOffCommand struct {
	light *Light
}

// Implement 'Execute' method for 'TurnOffCommand'
func (c *TurnOffCommand) Execute() {
	c.light.TurnOff()
}

// Define 'RemoteControl' struct with a 'command' field
type RemoteControl struct {
	command Command
}

// Define 'SetCommand' method for 'RemoteControl'
func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

// Define 'PressButton' method for 'RemoteControl'
func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

// Example usage
func main() {
	// Create a light
	light := &Light{}

	// Create commands
	turnOn := &TurnOnCommand{light: light}
	turnOff := &TurnOffCommand{light: light}

	// Create remote control
	remote := &RemoteControl{}

	// Set and execute turn on command
	remote.SetCommand(turnOn)
	remote.PressButton()

	// Set and execute turn off command
	remote.SetCommand(turnOff)
	remote.PressButton()
}
