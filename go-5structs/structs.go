package structs

import "math"


//Rectangle Definition
type Rectangle struct{
	Width float64
	Height float64
}
func (r Rectangle) Area() float64{
	return r.Width*r.Height
}

//Circle Definition
type Circle struct{
	Radius float64
}
func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}

//Triangle Definition
type Triangle struct{
	Base float64
	Height float64
}
func (t Triangle) Area() float64{
	return t.Base*t.Height/2
}

//Shape Definition
type Shape interface{
	Area() float64
}

//Interface Function
func calculateArea(s Shape) float64{
	return s.Area()
}

//my brain hurts but I am getting it
