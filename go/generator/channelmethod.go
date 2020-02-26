//https://stackoverflow.com/questions/34464146/the-idiomatic-way-to-implement-generators-yield-in-golang-for-recursive-functi
package main

import "fmt"

import "time"

func generate(abort <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			select {
			case ch <- i:
				fmt.Println("Sent", i)
			case <-abort:
				fmt.Println("Aborting")
				return
			}
		}
	}()
	return ch
}

func main() {
	abort := make(chan struct{})
	ch := generate(abort)
	for v := range ch {
		fmt.Println("Processing", v)
		if v == 3 {
			close(abort)
			break
		}
	}
	time.Sleep(time.Second)
}
