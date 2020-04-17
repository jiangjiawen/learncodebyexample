package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("exit")
	}()
	fmt.Println("hi!")
	defer func() {
		v := recover()
		fmt.Println("panic reconver", v)
	}()
	panic("bye bye")
	fmt.Println("can't be here!")
}
