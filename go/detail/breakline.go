//https://gfw.go101.org/article/details.html
package main

import (
	"fmt"
)

func False() bool {
	return false
}

func main(){
	switch False()
	{
	case true:fmt.Println("true")
	case false:fmt.Println("false")
	}
	switch False(){
	case true:fmt.Println("true")
	case false:fmt.Println("false")
	}
}