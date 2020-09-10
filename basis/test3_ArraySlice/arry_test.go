package ArraySliceTest

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	var arr1 = [4]int{1, 2, 3, 4}
	var arr2 = [...]int{1, 2, 3, 4, 5}
	t.Log(arr)
	t.Log(arr1)
	t.Log(arr2)
}

func TestArrayLoop(t *testing.T) {
	var arr2 = [...]int{1, 2, 3, 4, 5}

	// 传统for循环
	for i := 0; i < len(arr2); i++ {
		t.Log(i)
	}

	// for range
	for idx, num := range arr2 {
		t.Log(idx, num)
	}

	// for range 忽略索引下标
	for _, num := range arr2 {
		t.Log(num)
	}
}

// 数组截取
func TestArraySection(t *testing.T) {
	var arr2 = [...]int{1, 2, 3, 4, 5}
	var arr2Sec1 = arr2[:3]
	var arr2Sec2 = arr2[1:3]
	t.Log("arr2Sec1: ", arr2Sec1)
	t.Log("arr2Sec2: ", arr2Sec2)
}
