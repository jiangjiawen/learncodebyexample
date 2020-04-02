package main

import "fmt"

func doSth(sliceint []int) {
	fmt.Printf("before one=%v\n", sliceint)
	sliceanother := sliceint[:]
	sliceanother = append(sliceanother, 100)
	fmt.Printf("after append one=%v, another=%v\n", sliceint, sliceanother)
	sliceanother[0]=99
	fmt.Printf("after assign one=%v, another=%v\n", sliceint, sliceanother)
}

func doCopy(sliceint []int){
	fmt.Printf("before one=%v\n", sliceint)
	sliceanother := make([]int, len(sliceint), len(sliceint)+1)
	copy(sliceanother, sliceint)
	sliceanother = append(sliceanother, 100)
	fmt.Printf("after append one=%v, another=%v\n", sliceint, sliceanother)
	sliceanother[0]=99
	fmt.Printf("after assign one=%v, another=%v\n", sliceint, sliceanother)
}

func main() {
	slice10 := make([]int,1,10)
	doSth(slice10)
	slice10 = make([]int,1,10)
	doCopy(slice10)
}