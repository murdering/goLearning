package main

import "fmt"

func main() {

	//animals1 := []Animal{Cat{}, Dog{}, Cow{}}
	animals := []Animal{Cat{}, Dog{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	moumou := AnimalWithMou(Cow{})
	fmt.Println(moumou.Mou())

	Test("string")
	Test(123)
	Test(true)

	// advanced
	var d data = 15
	var x interface{} = d

	if n, ok := x.(fmt.Stringer); ok { // 转换为更具体的接口类型
		fmt.Println(n)
	}

	if d2, ok := x.(data); ok { // 转换回原始类型
		fmt.Println(d2)
	}

	//e:=x.(error)           // 错误:main.data is not error
	//fmt.Println(e)

	persons := []person{man{}, man{}}

	for _, person := range persons {
		person.run()
	}
}

type Animal interface {
	Speak() string
}

type AnimalWithMou interface {
	Mou() string
}

type Cat struct{}

func (c Cat) Speak() string {
	return "cat"
}

type Dog struct{}

func (d Dog) Speak() string {
	return "dog"
}

type Cow struct{}

func (cow Cow) Mou() string {
	return "cow"
}

func Test(params interface{}) {
	fmt.Println(params)
}

// advanced
type data int

func (d data) String() string {
	return fmt.Sprintf("data:%d", d)
}

type person interface {
	run()
}

type man struct {
}

func (m man) run() {
	println("man is running")
}
