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
func TestSumAll(t *testing.T){
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T){
	t.Run("make sum of some slices", func (t *testing.T){
		got := SumAllTails([]int{1,2}, []int{0,9})
		want := []int {2, 9}

		if !reflect.DeepEqual(got ,want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("safely sum empty slices", func (t *testing.T){
		got := SumAllTails([]int{}, []int{3,4,5})
		want := []int {0,9}

		if !reflect.DeepEqual(got ,want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	
}
//lets continue with arrays and slices 
//today we will build a function that sums all the numbers in a 
//slice except for the first
//first we build the test for it 
//test should fail we have no such function available
//lets make it
//lets see what the error will be
// the value for [1:]in the empty slice is out of range since the slice is empty
//lets fix our code so that it can accomodate such situations