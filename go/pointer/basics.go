package main

import "fmt"

func main() {
	p0 := new(int)
	fmt.Println(p0)
	fmt.Println(*p0)
	fmt.Println(&*p0)

	x:=*p0
	p1,p2:=&x,&x
	fmt.Println(p1,p2)
	fmt.Println(p1==p2)
	fmt.Println(p0==p1)

	p3:=&*p0
	fmt.Println(p0==p3)
	*p0,*p1=123,789
	fmt.Println(*p2,x,*p3)
	fmt.Printf("%T,%T\n",*p0,x)
	fmt.Printf("%T,%T\n",p0,p1)
}