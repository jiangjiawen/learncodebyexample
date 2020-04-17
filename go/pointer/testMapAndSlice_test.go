//https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go
// map is reference type
//推荐使用指针持有者类型来代替这个术语 来自go101 感觉是混乱。
package testMapAndSliceTest

import "testing"

func fn(m map[int]int) {
	m = make(map[int]int)
}

func fn2(m map[int]int) {
	m[2] = 2
}

func TestMain(t *testing.T) {
	var m map[int]int
	fn(m)
	t.Log(m == nil)
	m = make(map[int]int)
	t.Log(m == nil)
	fn2(m)
	t.Log(m)
}
