package patterns

import "fmt"

// Паттерн Builder относится к порождающим паттернам уровня объекта.

// Паттерн Builder определяет процесс поэтапного построения сложного продукта.
// После того как будет построена последняя его часть, продукт можно использовать.

// Builder provides a builder interface
type Builder interface {
	MakeHeader(str string)
	MakeBody(str string)
	MakeFooter(str string)
}

// Director implements a manager
type Director struct {
	builder Builder
}

// Construct tells Builder what to do and in which order
func (d *Director) Construct() {
	d.builder.MakeHeader("header")
	d.builder.MakeBody("body")
	d.builder.MakeFooter("footer")
}

// ConcreteBuilder implements Builder interface
type ConcreteBuilder struct {
	product *Product
}

// MakeHeader makes header
func (b *ConcreteBuilder) MakeHeader(str string) {
	b.product.Content += "<header>" + str + "</header>"
}

// MakeBody makes body
func (b *ConcreteBuilder) MakeBody(str string) {
	b.product.Content += "<body>" + str + "</body"
}

// MakeFooter makes footer
func (b *ConcreteBuilder) MakeFooter(str string) {
	b.product.Content += "<footer>" + str + "</footer>"
}

// Product implementation
type Product struct {
	Content string
}

// Show returns product
func (p *Product) Show() string {
	return p.Content
}

// BuilderPattern implements a builder pattern
func BuilderPattern() {
	product := new(Product)

	director := Director{&ConcreteBuilder{product}}
	director.Construct()

	fmt.Println(product.Show())
}
