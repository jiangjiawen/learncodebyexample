//https://medium.com/a-journey-with-go/go-understand-the-empty-interface-2d9fc1e5ec72
package main_test

import (
	"testing"
)

var x MultipleFieldStructure

type MultipleFieldStructure struct {
	a int
	b string
	c float32
	d float64
	e int32
	f bool
	g uint64
	h *string
	i uint16
}

//go:noinline
func emptyInterface(i interface{}) {
	s := i.(MultipleFieldStructure)
	x = s
}

//go:noinline
func typed(s MultipleFieldStructure) {
	x = s
}

func BenchmarkWithType(b *testing.B) {
	s := MultipleFieldStructure{a: 1, h: new(string)}
	for i := 0; i < b.N; i++ {
		typed(s)
	}
}

func BenchmarkWithEmptyInterface(b *testing.B) {
	s := MultipleFieldStructure{a: 1, h: new(string)}
	for i := 0; i < b.N; i++ {
		emptyInterface(s)
	}
}
