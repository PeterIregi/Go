package arrays


import(
	"testing"
)

func TestSum(t *testing.T){

	t.Run("collection of five numbers", func(t *testing.T){
		//lets create an array of size 5 since arrays have to have a defined size
		numbers := []int{1, 2, 3, 4, 5}
			//this is an array of intergers with the size of 5

		got := Sum(numbers)
		want := 15
	//our function should return an int 

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
//the test should fail
//the error is that we cannot use the type []int a slice of int when our
//function only takes arrays of size five [5]int
//lets change that so that it doesn't  cause any issues
//that means in the test and in the function
//we want the function to take in slices so that we are not held back by the size of the arrays
