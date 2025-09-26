package intermediate

import "fmt"

func main() {
	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, \rGo!"
	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of rawmessage variable is ", len(rawMessage))

	fmt.Println("Length of rawmessage variable is ", message[0])

	greeting := "Hello"
	name := "Alice"
	fmt.Println(greeting + name)

	str1 := "Apple"
	str2 := "banana"
	fmt.Println(str1 < str2)

	for i, char := range message {
		fmt.Printf("Character at index %d is %c", i, char)
	}
}
