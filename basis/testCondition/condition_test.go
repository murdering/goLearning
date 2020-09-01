package testCondition

import "testing"

// 多条件的switch
func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3, 5:
			t.Log("Odd")
		default:
			t.Log("default")
		}
	}
}

// 判断的switch就和if else一样的了
func TestSwitchCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("default")
		}
	}
}
