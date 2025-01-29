package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	X float64
	Y float64
}

func (r *Rectangle) Perimeter() float64 {
	return r.X*2 + r.Y*2
}

func (r *Rectangle) Area() float64 {
	return r.X * r.Y
}

type Circle struct {
	R float64
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.R, 2)
}

type Triangle struct {
	B float64
	H float64
}

func (t *Triangle) Area() float64 {
	return t.B / 2 * t.H
}
