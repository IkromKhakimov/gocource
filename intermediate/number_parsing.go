package intermediate

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("Parsed Integer:", num)
	fmt.Println("Parsed Integer:", num+1)

	numistr, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}

	fmt.Println("Parsed Integer:", numistr)

	floatstr := "3.14"
	floatval, err := strconv.ParseFloat(floatstr, 64)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Printf("Parsed float: %.2f\n", floatval)

	binaryStr := "1010" // 0 + 2 + 0 + 8 = 10
	decimal, err := strconv.ParseInt(binaryStr, 2, 6)
	if err != nil {
		fmt.Println("Error parsing binary value:", err)
	}
	fmt.Println("Parsed binary:", decimal)

	hexStr := "FF" // 0 + 2 + 0 + 8 = 10
	hex, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Error parsing binary value:", err)
	}
	fmt.Println("Parsed binary:", hex)

	invalidNum := "456abc" // 0 + 2 + 0 + 8 = 10
	invalidParse, err := strconv.Atoi(invalidNum)
	if err != nil {
		fmt.Println("Error parsing value:", err)
	}
	fmt.Println("Parsed invalid number:", invalidParse)

}
