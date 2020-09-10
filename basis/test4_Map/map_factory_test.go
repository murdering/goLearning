package MapTest

import "testing"

// value可以是方法
func TestMapFactory(t *testing.T) {
	var m = map[int]func(in int) int{}
	m[1] = func(in int) int {
		return in
	}

	m[2] = func(in int) int {
		return in * in
	}

	m[3] = func(in int) int {
		return in * in * in
	}

	t.Log(m[1](2), m[2](2), m[3](2))
}
