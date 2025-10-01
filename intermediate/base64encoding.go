package intermediate

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("He~lo, Base64 Encoding")

	// Encode Base64
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	// Decode from Base4
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(string(decoded))

	// Url safe, avoid '/'
}
