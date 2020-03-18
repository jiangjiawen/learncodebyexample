package main

func read(i interface{}) {
	println(i)
}

func main() {
	var i int8 = 1
	//a pointer to information about the type stored
	//a pointer to the associated data
	read(i)
}
