package main

import (
	"math"
)

type Point struct {
	X, Y float64
}

func (a Point) Add(b Point) Point {
	return Point{a.X + b.X, a.Y + b.Y}
}

func (a Point) Sub(b Point) Point {
	return Point{a.X - b.X, a.Y - b.Y}
}

// Imprecise method which does not guarantee v = v1 when t = 1,
// due to floating-point arithmetic error.
func lerp_v1(v0, v1 float64, t float64) float64 {

	return v0 + t*(v1-v0)
}

// Precise method which guarantees v = v1 when t = 1.
func lerp_v2(v0, v1 float64, t float64) float64 {

	return v0*(1-t) + t*v1
}

var lerp = lerp_v2

func PointLerp(p0, p1 Point, t float64) Point {

	return Point{
		X: lerp(float64(p0.X), float64(p1.X), t),
		Y: lerp(float64(p0.Y), float64(p1.Y), t),
	}
}

func Round(x float64) float64 {
	return math.Floor(x + 0.5)
}
