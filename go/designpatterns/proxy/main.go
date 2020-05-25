package main

type AdvancedLifeSupport interface {
	performCPR()
}

type EnagencyCallHandler struct {
}

func (h *EnagencyCallHandler) performCPR() {
	print("do chest compression 30s")
}

//proxy 区别于adapter
type Paramedic struct {
	h *EnagencyCallHandler
}

func (p *Paramedic) performCPR() {
	p.h.performCPR()
}

func main() {
	center := &Paramedic{h: &EnagencyCallHandler{}}
	center.performCPR()
}
