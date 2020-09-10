// 1. xxxx_test。以_test结尾
// 2. func TestXXX。以Test开头。
package test1_Opening

import "testing"

func TestFirst(t *testing.T) {
	t.Log("my first try!")
}
