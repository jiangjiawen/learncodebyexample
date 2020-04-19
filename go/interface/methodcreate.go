package main

import "fmt"

type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

func (p *Gopher) code() {
	fmt.Println("I am coding %s language\n", p.language)
}

func (p *Gopher) debug() {
	fmt.Println("I am debugging %s language\n", p.language)
}

func main() {
	var c coder = &Gopher{"Go"}
	c.code()
	c.debug()
	// var anotherC coder = Gopher("php")
	// anotherC.code()
	// anotherC.debug()
}
