package MapTest

import "testing"

// map初始化，不同的赋值方式
func TestMapInit(t *testing.T) {
	var m1 = map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("length of m1 is %d", len(m1))

	var m2 = map[int]int{}
	m2[4] = 16
	t.Logf("length of m2 is %d", len(m2))

	var m3 = make(map[int]int, 10)
	t.Logf("length of m3 is %d", len(m3))
}

// golang的map, key不存在的时候不会抛异常
func TestMapKeyNotExists(t *testing.T) {
	var m1 = map[int]int{}

	// m1[3] = 3333
	m1[0] = 0

	if v, exists := m1[3]; exists {
		t.Logf("m1 is an existed key with value: %d", v)
	} else {
		t.Logf("3 is not an existed key")
	}
}

// map遍历
func TestMapLoop(t *testing.T) {
	var m1 = map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
