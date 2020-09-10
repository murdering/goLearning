package ArraySliceTest

import "testing"

func TestSliceInit(t *testing.T) {
	var slice0 []int
	t.Log(len(slice0), cap(slice0))
	slice0 = append(slice0, 1)
	t.Log(len(slice0), cap(slice0))

	var slice1 = []int{1, 2, 3, 4}
	t.Log(len(slice1), cap(slice1))

	var slice2 = make([]int, 3, 5)
	t.Log(len(slice2), cap(slice2))

	// out of range
	// t.Log(slice1[0],slice1[1],slice1[2],slice1[3],slice1[4])

	// length就+1了
	slice2 = append(slice2, 1)
	t.Log(slice1[0], slice1[1], slice1[2], slice1[3])
	t.Log(len(slice2), cap(slice2))
}

// 切片length和capacity的变化, length是依量增加的，capacity是翻倍增加
func TestSliceGrowing(t *testing.T) {
	var slice []int
	for i := 0; i < 10; i++ {
		slice = append(slice, i)
		t.Log(len(slice), cap(slice))
	}
}

// 切片共享空间
func TestSliceShare(t *testing.T) {
	var months = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	var q2 = months[3:6]
	t.Log(q2, len(q2), cap(q2))

	var summer = months[5:8]
	t.Log(summer, len(summer), cap(summer))

	//修改summer值，共享都变了，引用类型
	summer[0] = "Unknown"
	t.Log(q2)
	t.Log(months)
}
