package main

import "fmt"

type Handler func(s string) string

type Middleware func(next Handler) Handler

func buildChain(h Handler,m ...Middleware) Handler {
	if len(m)==0{
		return h
	}
	return m[0](buildChain(h,m[1:cap(m)]...))
}

func PrintaMiddleware(h Handler) Handler{
	return func(s string) string{
		fmt.Println("do a")
		return h(s)
	}
}

func PrintbMiddleware(h Handler) Handler{
	return func(s string) string{
		fmt.Println("do b")
		return h(s)
	}
}

func PrintcMiddleware(h Handler) Handler{
	return func(s string) string{
		fmt.Println("do c")
		return h(s)
	}
}

func main(){
	var chainMiddlewares = []Middleware{
		PrintaMiddleware,
		PrintbMiddleware,
		PrintcMiddleware,
	}
	var handler = func(s string) string { return "do D" }
	

	fmt.Println(buildChain(handler, chainMiddlewares...)("d"))
}

