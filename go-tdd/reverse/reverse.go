package main

func Reverse(s string) string {

	stringRune := []rune(s)
	//turns our string into a slice of runes for more precise representation of special characters
	final := ""
	//empty string where we will put our output

	if len(s) == 0 {
		final = ""
	} else {
		for i := len(stringRune) - 1; i >= 0; i-- {
			final += string(stringRune[i])
		}
	}

	return final
}

//lets see if it works
//it didn't we try again later
//welcome back
//lets try again and see if our test passes or what we could do if it fails again
//lets modify the reverse function
//lets see if the test passes
//we get an error for when the string is empty lets fix that
//lets try it now
//the loop parameters were wrong it should be greater than zero not less than my fault since there is no negative index at least not here
//lets fix it and see
