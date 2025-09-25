package basics

import "fmt"

func main() {
	// panic(int)

	process2(10)

	process2(-3)
}

func process2(input int) {

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0 {
		fmt.Println("Before Panic")
		panic("input must be a non-negative number")
		// fmt.Println("After Panic")
		//defer fmt.Println("Deferred 3")
	}
	fmt.Println("Processing input:", input)
}
