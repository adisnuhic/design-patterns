/*
	Imagine you have a smart home with various electronic devices, such as lights and fans, and you want to control them through a
	smart home control panel. The control panel has buttons for turning on and off the lights and fans in different rooms.
	To implement this functionality, we use the Command pattern. The Command pattern allows us to encapsulate the actions (commands)
	we want to perform on the electronic devices as objects.
	Each command represents a specific action, such as turning on a light or a fan.
	These commands are then passed to the smart home control panel.
*/

package main

import "fmt"

// Command is the interface for different device control commands.
type Command interface {
	Execute()
}

// Light is the receiver that represents a light.
type Light struct {
	location string
}

func (l *Light) TurnOn() {
	fmt.Printf("Light in %s turned ON.\n", l.location)
}

func (l *Light) TurnOff() {
	fmt.Printf("Light in %s turned OFF.\n", l.location)
}

// Fan is the receiver that represents a fan.
type Fan struct {
	location string
}

func (f *Fan) TurnOn() {
	fmt.Printf("Fan in %s turned ON.\n", f.location)
}

func (f *Fan) TurnOff() {
	fmt.Printf("Fan in %s turned OFF.\n", f.location)
}

// LightOnCommand is a concrete command that turns on the light.
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.TurnOn()
}

// LightOffCommand is a concrete command that turns off the light.
type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.TurnOff()
}

// FanOnCommand is a concrete command that turns on the fan.
type FanOnCommand struct {
	fan *Fan
}

func (c *FanOnCommand) Execute() {
	c.fan.TurnOn()
}

// FanOffCommand is a concrete command that turns off the fan.
type FanOffCommand struct {
	fan *Fan
}

func (c *FanOffCommand) Execute() {
	c.fan.TurnOff()
}

// Invoker represents the smart home control panel.
type Invoker struct {
	commandStack []Command
}

func (i *Invoker) AddCommand(command Command) {
	i.commandStack = append(i.commandStack, command)
}

func (i *Invoker) ExecuteCommands() {
	for _, command := range i.commandStack {
		command.Execute()
	}
}

func main() {
	// Create the electronic devices (receivers)
	livingRoomLight := &Light{location: "Living Room"}
	bedroomFan := &Fan{location: "Bedroom"}

	// Create the concrete commands
	lightOnCommand := &LightOnCommand{light: livingRoomLight}
	lightOffCommand := &LightOffCommand{light: livingRoomLight}
	fanOnCommand := &FanOnCommand{fan: bedroomFan}
	fanOffCommand := &FanOffCommand{fan: bedroomFan}

	// Create the smart home control panel (invoker)
	controlPanel := &Invoker{}

	// Add commands to the control panel's stack
	controlPanel.AddCommand(lightOnCommand)
	controlPanel.AddCommand(fanOnCommand)
	controlPanel.AddCommand(lightOffCommand)
	controlPanel.AddCommand(fanOffCommand)

	// Execute the commands
	controlPanel.ExecuteCommands()
}
