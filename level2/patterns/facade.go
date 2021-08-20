package patterns

import (
	"fmt"
	"strings"
)

// Паттерн Facade предоставляет высокоуровневый унифицированный интерфейс
// в виде набора имен методов к набору взаимосвязанных классов или объектов
// некоторой подсистемы, что облегчает ее использование.

// Man implements man and facade
type Man struct {
	house *House
	tree  *Tree
	child *Child
}

// House implements subsystem House
type House struct {
}

// BuildHouse - implementation of Build
func (h *House) BuildHouse() string {
	return "Build house"
}

// Tree implements subsystem Tree
type Tree struct {
}

// GrowTree - implementation of grow
func (t *Tree) GrowTree() string {
	return "Tree grow"
}

// Child implements subsystem Child
type Child struct {
}

// BornChild implements born
func (c *Child) BornChild() string {
	return "Child born"
}

// NewMan - create man
func NewMan() *Man {
	return &Man{
		house: &House{},
		tree:  &Tree{},
		child: &Child{},
	}
}

// Todo returns what man must do
func (m *Man) Todo() string {
	result := []string{
		m.house.BuildHouse(),
		m.tree.GrowTree(),
		m.child.BornChild(),
	}
	return strings.Join(result, "\n")
}

// Facade outputs examle
func Facade() {
	man := NewMan()

	fmt.Println("Todo:\n", man.Todo())
}
