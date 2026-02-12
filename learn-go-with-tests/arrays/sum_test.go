package arrays


import(
	"testing"
	"reflect"
)

func TestSum(t *testing.T){

	t.Run("collection of five numbers", func(t *testing.T){
		numbers := []int{1, 2, 3, 4, 5}
			

		got := Sum(numbers)
		want := 15 

		if got != want{
			t.Errorf("got %d want %d given %v", got, want, numbers)

		}
	})
	t.Run("collection of any size ", func(t *testing.T){
		numbers := []int{1, 2, 3}
		got :=Sum(numbers)
		want := 6

		if got != want{
			t.Errorf("got  %d want %d given %v", got, want, numbers)
		}
	})
	
}
//lets continue where we left off
//lets test a function sumAll that adds a varryoing number of slices and produces the sum of each slice in the form of a slice
func TestSumAll(t *testing.T){
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
//the test should fail because we don't have a sum all function
//lets make it
//now it says we have no return in our function
//lets fix that
// see you next time