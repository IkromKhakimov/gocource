package intermediate

import "fmt"

func main() {
	var a int = 32
	b := int32(a)
	c := float64(b)
	//d := bool("correct")
	e := 3.14
	f := int(e)
	fmt.Println(f, c)
	g := "Hello @😄"
	var h []byte
	h = []byte(g)
	fmt.Println(h)
	i := []byte{255, 72, 132}
	j := string(i)
	fmt.Println(j)
}
