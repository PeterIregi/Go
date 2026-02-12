//lets write just enough code for the test to at least fail with something else
package arrays

// our function should take in an array of size five and give us an int

func Sum(nums []int)int{
	result:= 0
	for _, num := range nums{
		result += num
	}
	return result
}

//lets see
// it failed 
//now lets write enough code to make it pass
//it should pass now
// what  if we want something without the size defined 
//we should use slices lets try tpo change our tests first
//lets see if it passes
//we will continue with more tests later
// goodbye for now!
func SumAll(numbersToSum ... []int)[]int {
	//numbersToSum ... helps us to state that we want a varying ni=umber of arguments we dont know the number that will be given
	var sums []int

	for _,numbers := range numbersToSum{
		//lets append the values instread of adding them
		sums = append(sums,Sum(numbers))
	}
	return sums
}
// we cannot compare slices  like this  != we will have to use another way to use another method
//we will use reflect.DeepEqual() 
//lets see if that solves our issue
//lets see if it passes
//lets try again