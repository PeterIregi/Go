package mapstests


func Search(dictionary map[string]string, word string) string{
	return dictionary[word]
	//here we see some new syntax
	//to get an item from a map we can use the key 
	//like we used the index in arraysand others
}

//the test should compile but fail
//it should pass
//we continue later
