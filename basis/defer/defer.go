package main

import "fmt"

func main() {
	stackingDefers()
}

func stackingDefers() {
	fmt.Println("0")
	defer func() {
		fmt.Println("1")
	}()
	fmt.Println("1.1")
	defer func() {
		fmt.Println("2")
	}()
	defer func() {
		fmt.Println("3")
	}()
}
