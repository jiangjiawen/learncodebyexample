//https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ slice gotchas
//https://medium.com/@Jarema./golang-slice-append-gotcha-e9020ff37374
package main

import "fmt"

func doAssign(value []string) {
	fmt.Printf("value=%v\n", value)
	value2 := value[:]
	value2 = append(value2, "b")
	fmt.Printf("value=%v,value2=%v\n", value, value2)
	value2[0] = "z"
	fmt.Printf("value=%v,value2=%v\n", value, value2)
}

func doCopy(value []string) {
	fmt.Printf("value=%v\n", value)
	value2 := make([]string, len(value), len(value)+1)
	copy(value2, value)
	value2 = append(value2, "b")
	fmt.Printf("value=%v,copyvalue2=%v\n", value, value2)
	value2[0] = "z"
	fmt.Printf("value=%v,copyvalue2=%v\n", value, value2)
}

func main() {
	slice1 := []string{"a"}
	// doAssign(slice1)
	doCopy(slice1)

	slice10 := make([]string, 1, 10)
	slice10[0] = "a"
	// doAssign(slice10)
	doCopy(slice10)
}
