package patterns

import "fmt"

// Command provides a command interface
type Command interface {
	Execute() string
}

// ToggleOnCommand implements the command interface
type ToggleOnCommand struct {
	receiver *Receiver
}

// Execute command
func (t *ToggleOnCommand) Execute() string {
	return t.receiver.ToggleOn()
}

// ToggleOffCommand implements the command interface
type ToggleOffCommand struct {
	receiver *Receiver
}

// Execute command
func (t *ToggleOffCommand) Execute() string {
	return t.receiver.ToggleOff()
}

// Receiver implementation
type Receiver struct {
}

// ToggleOn implementation
func (r *Receiver) ToggleOn() string {
	return "Toggling on.."
}

// ToggleOff Implementation
func (r *Receiver) ToggleOff() string {
	return "Toggling off.."
}

// Invoker implementation
type Invoker struct {
	commands []Command
}

// StoreCommand adds command
func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

// UnstoreCommand removes command
func (i *Invoker) UnstoreCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

// Execute all stored commands
func (i *Invoker) Execute() string {
	var res string
	for _, command := range i.commands {
		res += command.Execute() + "\n"
	}
	return res
}

// CommandPattern outputs the example of pattern command
func CommandPattern() {
	invoker := &Invoker{}
	receiver := &Receiver{}

	invoker.StoreCommand(&ToggleOffCommand{receiver: receiver})
	invoker.StoreCommand(&ToggleOnCommand{receiver: receiver})

	fmt.Println(invoker.Execute())

}
