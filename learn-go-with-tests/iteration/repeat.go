package iteration


func Repeat(ch string, t int) string{
	result:=""
	for i:=0;i<t;i++{
		//basicaly this says start with i being 0(i:=0)
		//while i is less than 5 continue looping
		//and after each repetition increase i by one (i++)
		result+=ch
		//each time we repeat or loop  add ch which is our character to result
		//which means result starts as an empty string and our character is added 
		//and re added to the resulting string till the loop is no longer running then it stops and returns the result with the ch added five times
		//to the string that was empty
		//lets see if our test confirms that
	}
	return result
}

//it passed so we did what the test required lets add some more requirements to our tests
//now it should pass //

//see you later