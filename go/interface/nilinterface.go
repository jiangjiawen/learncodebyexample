//The call to Bang succeeds because it applies to pointers to a Bomb: there is no need to dereference the pointer to call the method. 
//The Boom method acts on a value and so a call causes pointers to be dereferenced, which causes a panic.
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