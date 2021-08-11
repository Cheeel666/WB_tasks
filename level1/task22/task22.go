package task22

import (
	"fmt"
	"sort"
)

type Car struct {
	Name    string
	Mileage int
}

type byMileage []Car

func (a byMileage) Len() int           { return len(a) }
func (a byMileage) Less(i, j int) bool { return a[i].Mileage < a[j].Mileage }
func (a byMileage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func Run() {
	slice := []int{5, 4, 1, 2, 0, 1, 4, 10, 8}

	autopark := []Car{
		{"bmw", 1000},
		{"mercedes", 12000},
		{"audi", 200},
		{"lexus lx 450", 100},
	}

	fmt.Println("before sort", slice)
	sort.Ints(slice)
	fmt.Println("after sort: ", slice)

	sort.Sort(byMileage(autopark))
	fmt.Println(autopark)
}
