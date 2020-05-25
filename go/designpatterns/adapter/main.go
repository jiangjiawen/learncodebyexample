//https://github.com/bvwells/go-patterns/blob/master/structural/adapter.go

package main

import (
	"fmt"
)

type Target interface {
	Execute()
}

type Adaptee struct{}

func (a *Adaptee) SpecificExecute() {
	fmt.Println("hello")
}

type Adapter struct {
	*Adaptee
}

func (a *Adapter) Execute() {
	a.SpecificExecute()
}

func main() {
	adapter := Adapter{}
	adapter.Execute()
}
