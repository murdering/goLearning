package basis

import (
	"fmt"
)

func main() {
	var number3 = [9]string{"a", "b", "c", "d", "e"}
	var slice1 = number3[1:4:6]
	var cap1 = cap(slice1)

	fmt.Println(slice1)
	fmt.Println("slice capacity: ", cap1)

	number3[2] = "bb"
	fmt.Println(number3)
	fmt.Println(slice1)
}
