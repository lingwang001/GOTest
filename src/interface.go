package main

import "fmt"

type general interface {
	area() float64
}

type rect struct {
	w float64
	l float64
}

type zeng struct {
	b float64
}

func (r rect) area() float64 {
	return r.w * r.l
}

func (z zeng) area() float64 {
	return z.b * 4
}

func gene(g general) float64 {
	return g.area()
}

func main() {

	var a1 rect
	var a2 zeng

	a1.l = 2
	a1.w = 4

	a2.b = 4

	fmt.Println(gene(a1))
	fmt.Println(gene(a2))

}
