// https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/
// And the since built-in collections (map, slice and array) are references and are mutable,
// copying a struct that contains one of these just copies the pointer to the same underlying memory.
package main

import "fmt"

type S struct {
	A string
	B []string
}

func main() {
	x := S{"x-A", []string{"x-B"}}
	y := x
	y.A = "y-A"
	y.B[0] = "y-B"
	fmt.Println(x, y)
}
