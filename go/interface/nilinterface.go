//This is because interface values are fat pointers. 
//The first element of this pair is the pointer to the method dispatch table for the implementation of the Bomb interface by the Explodes type, 
//and the second element is the address of the actual Explodes object, which is nil.
package main

type Explodes interface {
	Bang()
	Boom()
}

type Bomb struct {}
func (*Bomb) Bang() {}
func (Bomb) Boom() {}

func main() {
	var bomb *Bomb =nil
	var explodes Explodes = bomb
	println(bomb,explodes)
	if explodes!=nil {
		println("Not nil!")
		explodes.Bang()
		explodes.Boom()
	}else{
		println("nil!")
	}
}