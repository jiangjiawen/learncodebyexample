// this is from this source:
// https://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/
// There are three conditions:
//(1) Variable must not be modified uses value;
//(2) Variable is a large struct uses pointer;
//(3) Variable is a map or slice uses values;
// especially, Passing by value often is cheaper!!!
package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func changeNameByPassByvalue(p Person) {
	p.firstName = "Bob"
}

func changeNameByPassByPointer(p *Person) {
	p.firstName = "Bob"
}

func main() {
	person := Person{
		firstName: "Alice",
		lastName:  "Dow",
	}

	changeNameByPassByvalue(person)

	fmt.Println(person)

	changeNameByPassByPointer(&person)

	fmt.Println(person)
}
