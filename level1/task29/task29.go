package task29

import (
	"fmt"
	"math"
	"math/big"
)

func mult(a, b int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(a), big.NewInt(b))
}

func div(a, b int64) *big.Float {
	return new(big.Float).Quo(big.NewFloat(float64(a)), big.NewFloat(float64(b)))
}

func sub(a, b int64) *big.Int {
	return new(big.Int).Sub(big.NewInt(a), big.NewInt(b))
}

func add(a, b int64) *big.Int {
	return new(big.Int).Add(big.NewInt(a), big.NewInt(b))
}

func Run() {
	var a int64 = int64(math.Pow(2, 22))
	var b int64 = int64(math.Pow(2, 20)) + 1
	var choice int
	fmt.Println("a = ", a, ", b = ", b, "; \nInput your choice:\n1)Multipy\n2)Division\n3)Addition\n4)Subtration")
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		fmt.Println("Result:", mult(a, b))
	case 2:
		fmt.Println("Result:", div(a, b))
	case 3:
		fmt.Println("Result:", add(a, b))
	case 4:
		fmt.Println("Result:", sub(a, b))
	}
}
