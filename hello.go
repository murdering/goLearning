package main

import "fmt"

func main() {
	functionLearn()
}

func functionLearn() {
	var power int
	power = 100
	fmt.Println("power is", power)

	power1 := getPower()
	fmt.Println("power1 is", power1)

	var power2 = getPower()
	fmt.Println("power2 is", power2)

	var name, count = getPowerWithName("leo")
	fmt.Println("name is", name, count)
}

func getPower() int {
	return 10000
}

func getPowerWithName(name string) (string, int) {
	return name, len(name)
}
