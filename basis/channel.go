package basis.channel

import "fmt"

func main() {
	var ints = make(chan int, 5)
	fmt.Println("ints: ", ints)

	ints <- 5
	var int1, ok = <-ints
	fmt.Println("int: ", int1, "ok: ", ok)

	close(ints)
	fmt.Println("ints: ", ints)
}
