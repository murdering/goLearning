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

	fmt.Println("2")

	defer func() {
		fmt.Println("3")

		defer func() {
			fmt.Println("3.1")
		}()

		defer func() {
			fmt.Println("3.2")
		}()
	}()

	defer func() {
		fmt.Println("4")
	}()
}
