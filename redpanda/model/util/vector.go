package main

import "math"

type Vector struct {
	x float64
	y float64
}

func (v Vector) size() (vectorSize float64) {
	x2 := math.Pow(v.x, 2)
	y2 := math.Pow(v.y, 2)
	vectorSize = math.Pow(x2+y2, 0.5)
	return
}

func (v1 *Vector) add(v2 *Vector) (result Vector) {
	result.x = v1.x + v2.x
	result.y = v1.y + v2.y
	return
}

func (v1 *Vector) subtract(v2 Vector) (result Vector) {
	var negV2 = v2.multiply(-1)
	result = v1.add(&negV2)
	return
}

func (v Vector) multiply(factor float64) (result Vector) {
	x := factor * v.x
	y := factor * v.y
	result = Vector{x, y}
	return
}

func (v1 Vector) distance(v2 Vector) (result float64) {
	result = v1.subtract(v2).size()
	return
}
