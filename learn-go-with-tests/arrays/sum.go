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
