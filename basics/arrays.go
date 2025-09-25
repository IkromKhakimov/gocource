package basics

import "fmt"

func main() {

	// var arrayName [size]elementType

	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	numbers[0] = 9
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Oranage", "Grapes"}
	fmt.Println("Fruits array: ", fruits)

	fmt.Println("Third element:", fruits[2])

	//originalArray := [3]int{1, 2, 3}
	//copiedArray := originalArray
	//
	//copiedArray[0] = 100
	//
	//fmt.Println("Original array:", originalArray)
	//fmt.Println("Copied array:", copiedArray)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index,", i, ":", numbers[i])
	}

	for _, v := range numbers {
		fmt.Printf("Value %d\n", v)
	}

	a, _ := someFunction()
	fmt.Println(a)
	// fmt.Println(b)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray
	copiedArray[0] = 100

	fmt.Println("Original array:", originalArray)
	fmt.Println("Copied array:", *copiedArray)

}

func someFunction() (int, int) {
	return 1, 2
}
