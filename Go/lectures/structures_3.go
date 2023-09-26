package main

import "math"

type Circle struct {
	radius float64
}

type Rectangle struct {
	width float64
	height float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64{
	return r.height*r.width
}

func (c Circle) Area() float64{
	return math.Pi*math.Pow(c.radius,2)
}
