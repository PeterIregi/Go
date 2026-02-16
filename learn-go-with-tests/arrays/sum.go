
package arrays



func Sum(nums []int)int{
	result:= 0
	for _, num := range nums{
		result += num
	}
	return result
}

func SumAll(numbersToSum ... []int)[]int {
	var sums []int

	for _,numbers := range numbersToSum{
		sums = append(sums,Sum(numbers))
	}
	return sums
}
func SumAllTails(numsToSum ...[]int) []int{
	var sums []int
	for _, numbers := range numsToSum{
		//this gives us the values from the second index to the last
		//since indexing starts at 0
		if len(numbers)==0{
			sums=append(sums,0)
		}else{
			tail:=numbers[1:]
			sums =append(sums, Sum(tail))
		}
		
	}
	return sums
}
//lets see if it passes
//it passed but what if  we had an empty slice mixed in with our aslices
//would our function break
//lets test that
//lets see if this works
//it worked
//that was it for today see you later
