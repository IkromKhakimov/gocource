package intermediate

import "fmt"

func main() {
	// var prt *int

	var ptr *int
	var a int = 10
	ptr = &a
	*ptr = 20

	fmt.Println(a)
	fmt.Println(ptr)
	// fmt.Println(*ptr)
	//if ptr == nil {
	//	fmt.Println("Pointer is nil")
	//}

	modifyValue(ptr)
	fmt.Println(a)
}

func modifyValue(ptr *int) {
	*ptr++
}
