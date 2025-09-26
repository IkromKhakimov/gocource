package basics

import "fmt"

func main() {
	process3()
	fmt.Println("Returned from process")
}

func process3() {
	defer func() {
		// if r := recover(); r != nil {
		r := recover()
		if r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start Process")
	panic("Something went wrong!")
	fmt.Println("End Process")
}
