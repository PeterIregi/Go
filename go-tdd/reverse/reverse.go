package main

func Reverse(s string) string {
	result := []rune(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {

		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

//lets see if it works
