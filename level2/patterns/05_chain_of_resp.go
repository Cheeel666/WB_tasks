package patterns

import "fmt"

// Handler provides a handler interface
type Handler interface {
	SendRequest(message int) string
}

// ConcreteHandlerA implements Handler interface
type ConcreteHandlerA struct {
	next Handler
}

// SendRequest sending request
func (h *ConcreteHandlerA) SendRequest(message int) (res string) {
	if message == 1 {
		res = "I am handler A and i am processing your request"
	} else {
		res = h.next.SendRequest(message)
	}
	return
}

// ConcreteHandlerB implements Handler Interface
type ConcreteHandlerB struct {
	next Handler
}

// SendRequest sending request
func (h *ConcreteHandlerB) SendRequest(message int) (res string) {
	if message == 2 {
		res = "I am handler B and i am processing your request"
	} else {
		res = h.next.SendRequest(message)
	}
	return
}

// ConcreteHandlerC implements Handler interface
type ConcreteHandlerC struct {
	next Handler
}

// SendRequest sending request
func (h *ConcreteHandlerC) SendRequest(message int) (res string) {
	if message == 3 {
		res = "I am handler C and i am processing your request"
	} else {
		res = h.next.SendRequest(message)
	}
	return
}

// ChainOfResponsibilityPattern output example of pattern usage
func ChainOfResponsibilityPattern() {
	handlers := &ConcreteHandlerA{next: &ConcreteHandlerB{next: &ConcreteHandlerC{}}}

	fmt.Println(handlers.SendRequest(2))
}
