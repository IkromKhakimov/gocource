package basics

import "fmt"

func main() {
	// ... Ellipsis
	// func functionName(param1 type1, param2 type2, param

	// fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))
	sequence, total := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(sequence, total)

	sequence2, total2 := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(sequence2, total2)

	numbers := []int{1, 2, 3, 4, 5, 9}

	sequence3, total3 := sum(3, numbers...)
	fmt.Println(sequence3, total3)
}

func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}
