package intermediate

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("User env var:", user)
	fmt.Println("Home env var:", home)

	err := os.Setenv("FRUIT", "APPLE")
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
	}

	fmt.Println("FRUIT env var:", os.Getenv("FRUIT"))

	//for _, e := range os.Environ() {
	//	kvpair := strings.SplitN(e, "=", 1)
	//	fmt.Println("Key:", kvpair[0], "Value:", kvpair[0])
	//}

	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println("Error unsetting environment variable:", err)
	}

	fmt.Println("FRUIT env var:", os.Getenv("FRUIT"))

	str := "a=b=c=d"

	fmt.Println(strings.SplitN(str, "=", -1))
	fmt.Println(strings.SplitN(str, "=", 0))
	fmt.Println(strings.SplitN(str, "=", 1))
	fmt.Println(strings.SplitN(str, "=", 2))
	fmt.Println(strings.SplitN(str, "=", 3))
	fmt.Println(strings.SplitN(str, "=", 4))
	fmt.Println(strings.SplitN(str, "=", 5))
}
