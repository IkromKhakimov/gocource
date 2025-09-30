package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//reader := bufio.NewReader(strings.NewReader("Hello, bufio packageee!\nHow are you doing?"))
	//
	//// Reading byte slice
	//data := make([]byte, 20)
	//n, err := reader.Read(data)
	//if err != nil {
	//	fmt.Println("Error reading:", err)
	//	return
	//}
	//fmt.Printf("Read %d bytes: %s\n", n, data[:n])
	//
	//line, err := reader.ReadString('\n')
	//if err != nil {
	//	fmt.Println("Error reading:", err)
	//	return
	//}
	//fmt.Println("Read string:", line)

	writer := bufio.NewWriter(os.Stdout)

	// Writing byte slice
	data := []byte("Hello, bufio package!\n")
	nm, err := writer.Write(data)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", nm)

	// Flush the buffer to ensure all data in written to os.Stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	// Writing string
	str := "This is a string\n"
	nm, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes.\n", nm)

	// Flush the buffer
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}
}
