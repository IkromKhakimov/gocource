package intermediate

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file.", file)
	}
	defer file.Close()

	// write data to file
	data := []byte("Hello World!\n\n\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", file)
	}

	fmt.Println("Data has been written to file")

	file, err = os.Create("writeString.txt")
	if err != nil {
		fmt.Println("Error writing to file:", file)
	}
	defer file.Close()

	_, err = file.WriteString("Hello Go!\n\n\n")
	if err != nil {
		fmt.Println("Error writing to file:", file)
	}
}
