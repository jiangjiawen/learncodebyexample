//function curry
//https://blog.usejournal.com/function-currying-in-go-a88672d6ebcf
package main

import "fmt"

func multiply(a int) func(int) int {
	return func(i int) int {
		return a * i
	}
}

func subtract(a int) func(int) int {
	return func(i int) int {
		return i - a
	}
}

func main() {
	var in = 1
	m := multiply(4)
	s := subtract(10)
	sm := func(i int) int { return m(s(i)) }
	ms := func(i int) int { return s(m(i)) }
	fmt.Printf("%v %v\n", ms(in), sm(in))
}
