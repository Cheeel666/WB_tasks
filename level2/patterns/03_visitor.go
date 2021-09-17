package patterns

import "fmt"

// Паттерн Visitor относится к поведенческим паттернам уровня объекта.

// Паттерн Visitor позволяет обойти набор элементов (объектов) с разнородными интерфейсами,
// а также позволяет добавить новый метод в класс объекта, при этом, не изменяя сам класс этого объекта.

// Visitor provides a visitor interface
type Visitor interface {
	VisitSushiBar(p *SushiBar) string
	VisitBurgerBar(p *BurgerBar) string
	VisitPizzeria(p *Pizzeria) string
}

// Place provides interface for the place visitor should visit
type Place interface {
	Accept(v Visitor) string
}

// City implements a collection to visit
type City struct {
	places []Place
}

// Add appends Places to collection
func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

// Accept implements visit all places in the city
func (c *City) Accept(v Visitor) string {
	var res string
	for _, p := range c.places {
		res += p.Accept(v)
	}
	return res
}

// People implements Visitor interface
type People struct {
}

// VisitSushiBar implements visit sushi
func (v *People) VisitSushiBar(p *SushiBar) string {
	return p.BuySushi()
}

// VisitPizzeria implements visit pizzeria
func (v *People) VisitPizzeria(p *Pizzeria) string {
	return p.BuyPizza()
}

// VisitBurgerBar implements visit burger
func (v *People) VisitBurgerBar(p *BurgerBar) string {
	return p.BuyBurger()
}

// SushiBar implements the place interface
type SushiBar struct {
}

// Accept implementation
func (s *SushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(s)
}

// BuySushi implementation
func (s *SushiBar) BuySushi() string {
	return "Buy sushi.."
}

// BurgerBar implements the place interface
type BurgerBar struct {
}

//Accept implementation
func (b *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurgerBar(b)
}

// BuyBurger implementation
func (b *BurgerBar) BuyBurger() string {
	return "Buy burger.."
}

// Pizzeria implements the place interface
type Pizzeria struct {
}

// Accept implementation
func (p *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(p)
}

//BuyPizza implementation
func (p *Pizzeria) BuyPizza() string {
	return "Buy pizza"
}

// VisitorPattern outputs the pattern usage
func VisitorPattern() {
	city := new(City)

	city.Add(&SushiBar{})
	city.Add(&Pizzeria{})
	// city.Add(&BurgerBar{})

	fmt.Println(city.Accept(&People{}))
}
