package perimeter

import "math"

// Shape - represents a two dimensional object
type Shape interface {
	Area() float64
}

// Rectangle - represents a rectangle by specifying its width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Area - returns the area of a rectangle given a Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle - a circle as represented by radius
type Circle struct {
	Radius float64
}

// Area - calculates the area of a cicle given a radius
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter - returns the perimeter given a width and height
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// // Area - returns the area given a width and a height
// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Width * rectangle.Height
// }
