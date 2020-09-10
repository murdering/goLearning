package test1_Opening

import (
	"fmt"
	"testing"
)

// 菲波拉切数列
func TestFiboList(t *testing.T) {
	var a = 1 // 第一种赋值方式
	var b = 1

	/*
		var (
			a = 2
			b = 3
		)
	*/

	fmt.Print(a)
	for i := 0; i < 10; i++ {
		fmt.Print(" ", b)
		temp := a // 第二种赋值方式
		a = b
		b = a + temp
	}
}

// 交换参数
func TestExchangeVariables(t *testing.T) {
	var a = 1
	var b = 2
	/*
		var temp = a
		a = b
		b = temp
	*/
	a, b = b, a
	t.Log(a, b)
}
