package main

import "fmt"

func main() {

	fmt.Println(testFunc("Obama"))

	intptr := 10
	fmt.Println(testFuncWithPointer(&intptr))

	// 闭包
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func testFunc(name string) (string, int) {
	return name, len(name)
}

func testFuncWithPointer(intPtr *int) int {
	return *intPtr
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
