package intermediate

import (
	"fmt"
	"time"
)

func main() {
	tz, _ := time.LoadLocation("Asia/Tashkent")
	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current Unix Time:", unixTime)
	t := time.Unix(unixTime, 0)
	fmt.Println(t)
	fmt.Println(now.In(tz))
	fmt.Println()
}
