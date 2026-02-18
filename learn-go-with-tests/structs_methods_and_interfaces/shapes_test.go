// we will try do deal with structs
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
	t.Run("rectangles", func (t *testing.T){
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		want:=72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("circles", func (t *testing.T){
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g ", got, want)
		}
	})
		
	
}
//lets continue
//we want to build a function that finds the area of a circle
//should fail coz of some obvious reasons
//circle is undefined and cannot be passed to our function
//lets try and fix that
//now it should fail because we have no method in both rectangle and circle called Area
//lets see