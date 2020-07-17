package main

import "fmt"

func main() {
	// a := 6
	if a := 5; a < 1 {
		fmt.Println("OK", a)
	} else if a > 5 {
		fmt.Println("OjbK", a)
	} else {
		fmt.Println("not ok", a)
	}

	i := 2
	switch i {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	case 4, 5, 6:
		fmt.Println("4, 5, 6")
	default:
		fmt.Println("Default")
	}

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum: ", sum)

	myfunc()
}

func myfunc() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}
