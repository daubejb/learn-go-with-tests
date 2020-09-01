package perimeter

// Rectangle - represents a rectangle by specifying its width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter - returns the perimeter given a width and height
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area - returns the area given a width and a height
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
