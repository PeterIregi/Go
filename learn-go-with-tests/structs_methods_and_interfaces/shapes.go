package structs_and_methods

import(
	"math"
)


type Rectangle struct {
	Width float64
	Height float64
}
func (r Rectangle)Area()float64{
	return r.Width * r.Height
}

type Circle struct{
	Radius float64
}
func (c Circle) Area()float64{
	return math.Pi * c.Radius *c.Radius
}
//this interface will accept any shape that has the method Area that gives a float64 output
//that is the two shapes we have currently lets see if this works now
type Shape interface {
	Area()float64
}

func Perimeter(rectangle Rectangle) float64{
	return 2*(rectangle.Width + rectangle.Height)

}
// lets make some changes to our test file so that we can learn about interfaces
