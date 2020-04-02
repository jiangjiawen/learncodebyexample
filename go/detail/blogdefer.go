//来自蔡超谈软件 公众号 推荐他在极客时间的go课程
package main

import "fmt"

func doPrint() func() {
	fmt.Println("inner")
	return func() {
		fmt.Println("innerinner")
	}
}

func main() {
	defer doPrint()()
	fmt.Println("main")
}
