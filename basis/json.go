package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	ExampleUnmarshal()
}

func ExampleUnmarshal() {
	var jsonBlob = []byte(`[
	{"Name": "Platypus", "order": "Monotremata"},
    {"Name": "Platypus1", "order": "Monotremata2"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
	// Output:
	// [{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
}
