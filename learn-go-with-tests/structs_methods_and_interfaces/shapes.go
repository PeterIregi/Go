package structs_and_methods

import(
	"math"
)

//this is the syntax for declaring structs
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

func Perimeter(rectangle Rectangle) float64{
	//see we are stating rectagle as a type 
	return 2*(rectangle.Width + rectangle.Height)
	//to refer to the values in our struct we use the . followed by the value name

}

//we also stopped using  the .2f in the error state ment coz g //offers the decimal 
//to the precise decimal as compared to the f and in this case preciseness is what we are after
//lets see now
//we still cannot pass circle as aruments in our function
//the test will fail but for other reasons 
//we also cannot re declare this function
//but we can declare it as a method in our structs lets change the test first
//lets make the necessary changes
//now it should fail because of the out put 
//now lets fix the out put 
//for us to find the area of a circle we need to borrow pi 
//from the math package 
//hence we need to import it to use the pi constant
//it should work now
//thats all for today