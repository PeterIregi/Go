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

	checkArea := func(t *testing.T, shape Shape, want float64){
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangles", func (t *testing.T){
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)

		
	})
	t.Run("circles", func (t *testing.T){
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)

	})
		
	
}
//since there is some repeating code here lets try to reduce that
//lets see what ur error will be
//it says shape is undefined lets define it then
//we will do more with structs methods and interfaces later including incoporating them in our tests later