package main

import "fmt"

func main() {
	//var variable = 1
	//switch variable {
	//case 1:
	//case 2:
	//default:
	//}

	fruit := "apple"

	switch fruit {
	case "apple":
		fmt.Println("It's an apple")
	case "banana":
		fmt.Println("It's a banana")
	default:
		fmt.Println("Unknown Fruit!")
	}

	// Multiple Conditions
	day := "Monday"

	switch day {
	case "Monday", "Thuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday.")
	case "Sunday":
		fmt.Println("It's a weekend")
	default:
		fmt.Println("Invalid day.")
	}

	number := 15

	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 19")
	default:
		fmt.Println("Number is 20 or more")
	}

	num := 2

	switch {
	case num > 1:
		fmt.Println("Greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Not Two")
	}

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown Type")
	}
}
