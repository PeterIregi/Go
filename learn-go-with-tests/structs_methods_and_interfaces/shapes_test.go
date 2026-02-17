// we will try do deal with structs
package structs_and_methods

import(
	"testing"
)


func TestPerimeter(t *testing.T){
	//lets try to find the perimeter of a rectangle
	//lets use our struct now
	rectangle := Rectangle{10.0, 10.0}
	//and pass it here as our variable
	got := Perimeter(rectangle)
	want := 40.0
	//we are using decimal points instesd of normal numbers because these are not intergers they 
	//are what we call floats thus the decimal

	if got != want{
		t.Errorf("got %.2f want %.2f", got, want)
		//we are using %.2f so that it gives us the float value but cuts it off at two decimal places

	}
}//lets see if our test fails as it should 
//lets make the function for the test now
// lets try to find the area of a rectangle
func TestArea(t *testing.T){
	rectangle := Rectangle{12.0, 6.0}
	got := Area(rectangle)
	want:=72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

//this will test our uncreated function lets see if ti fails coz of that
//lets fix that 

//lets change the functions so they take rectangle as an argument
