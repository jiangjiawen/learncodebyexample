package main

import "fmt"

type Handler func(s string) string

type Middleware func(next Handler) Handler

func Middlewares(middlewares ...Middleware) Middleware{
	return func(next Handler) Handler{
		for i:=len(middlewares)-1;i>=0;i--{
			next=middlewares[i](next)
		}
		return next
	}
}

func A(next Handler) Handler {
	return func(s string) string { return "(A:" + next(s) + ":A)" }
}

func B(next Handler) Handler {
	return func(s string) string { return "(B:" + next(s) + ":B)" }
}

func C(next Handler) Handler {
	return func(s string) string { return "(C:" + next(s) + ":C)" }
}

func main() {

	var (
		handler = func(s string) string { return "("+s+")" }
		mid     = Middlewares(A, B, C)
	)

	fmt.Println(mid(handler)("value"))
}