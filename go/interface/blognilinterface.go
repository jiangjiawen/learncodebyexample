package main

type Explodes interface {
    Bang()
    Boom()
}

// Type Bomb implements Explodes
type Bomb struct {}
func (*Bomb) Bang() {}
func (Bomb) Boom() {}

func main() {
    var bomb *Bomb = nil
    var explodes Explodes = bomb
    println(bomb, explodes) // '0x0 (0x1084fe0,0x0)'
    if explodes != nil {
        println("Not nil!") // 'Not nil!' What are we doing here?!?!
        explodes.Bang()     // works fine
        explodes.Boom()     // panic: value method main.Bomb.Boom called using nil *Bomb pointer
    } else {
        println("nil!")     // why don't we end up here?
    }
}