package main

import "fmt"

import "runtime"

func giveString(astring chan string) {
	astring <- "Message"
}

func main() {
	fmt.Println(runtime.NumGoroutine())
	astring := make(chan string)
	fmt.Println(runtime.NumGoroutine())
	go giveString(astring)
	fmt.Println(runtime.NumGoroutine())
	var result = <-astring
	fmt.Println(result)
	fmt.Println(runtime.NumGoroutine())
}
