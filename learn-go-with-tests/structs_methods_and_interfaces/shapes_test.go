package structs_and_methods

import(
	"testing"
)


func TestPerimeter(t *testing.T){
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want{
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T){
//lets make a slice of struct to hold our test cases

	areaTests := []struct{
		name string
		shape Shape
		want float64
	}{
		//inside we will put our shapes
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}
	//now lets loop through our slice of structs and run the tests
	//lets add a name and some descriptions to our items

	for _, tt := range areaTests {
		//this is a different kind of for loop here we are offered the index and the item 
		//at the position we are at in the item we are looping through
		//but since we have no use for the index we will put the underscore
		//in Go we nothig is left unsused
		//first off fo more clarity and so that we can use t.Run 
		//lets make changes to our struct above //lets see the changes
		t.Run(tt.name, func (t *testing.T){
			got:=tt.shape.Area()
			if got != tt.want{
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
	
}

//lets try to improve our tests with the knowledge we have gained so far on structs
//lets see if our tests run now
//before we leave lets add another shape triangle and see how that works out
//lets make the changes in the shapes.go file
//that was it for this series I hope ypu learned as much as I did
