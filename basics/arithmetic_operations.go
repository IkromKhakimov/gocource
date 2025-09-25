package basics

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Subtraction: ", result)

	result = a * b
	fmt.Println("Multiplication: ", result)

	result = a / b
	fmt.Println("Division: ", result)

	result = a % b
	fmt.Println("Modulus: ", result)

	const p float64 = 22 / 7.0
	fmt.Println("p: ", p)

	// Overflow with signed integers
	var maxInt int64 = 922337203685547758
	fmt.Println("Max Int: ", maxInt)

	// maxInt = maxInt + 1
	fmt.Println("Max Int: ", maxInt)

	// Overflow with unsigned integers
	var uMaxUint uint64 = 18446744073709551615
	fmt.Println("Max Uint: ", uMaxUint)

	uMaxUint = uMaxUint + 1

	var smallFloat float64 = 0.1
	fmt.Println("Small Float: ", smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println("Small Float: ", smallFloat)

}
