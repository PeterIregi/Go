package structs_and_methods


//this is the syntax for declaring structs
type Rectangle struct {
	Width float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64{
	//see we are stating rectagle as a type 
	return 2*(rectangle.Width + rectangle.Height)
	//to refer to the values in our struct we use the . followed by the value name

}

func Area(rectangle Rectangle) float64{
	return rectangle.Width * rectangle.Height
}
//should pass now
//what if we wanted to find the area of a rectangle only and didn't want any 
//other shape
// we could rename the function to AreaOfRectangle 

//or we could create a struct called rectangle and pass it to our function
//lets try that
//lets change the test first
//it should pass as it did before
//small doses enough for today