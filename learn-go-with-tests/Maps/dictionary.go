package mapstests

import(
	"errors"
)


type Dictionary map[string]string
//we have declared our type dictionary of map
//now on to the function

//we would have used pointers if we were actualy changing something in our dictionary but since we are just displaying it 
//.we will leave it as is
func (d Dictionary)Search(word string) (string, error){
	definition, ok := d[word]
	//it will give us the value in the key [word] and ok  will be nil 
	if  !ok {
		//this says what to do if ok is not nil 
		return "", errors.New("could not find the word you are looking for")
	}
	
	return definition, nil
}

//now what if we wantede to search fo a word that  is not in our dictionary what happens 
//first the test
//lets make a test for unknown words or words not in our dictionary first
//does this solve any thing since now we have two return values for our method
//lets import the errors package first 
//lets try it now//
//the problem was the string we were returning as the error
//later