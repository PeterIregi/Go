package mapstests

import(
	//"errors"
)


type Dictionary map[string]string

const (
	ErrorNotFound = DictionaryErr("Could not find the word you were looking for")
	ErrorWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error()string{
	return string(e)
}

func (d Dictionary)Search(word string) (string, error){
	definition, ok := d[word]
	 
	if  !ok { 
		return "", ErrorNotFound
	}
	
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error{
	_, err := d. Search(word)

	switch err{
	case ErrorNotFound:
		d[word]=definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update (word, definition string){

	d[word] = definition
}

//we can do something so that we declare multiple values with on far statement
//lets see now what the compiler tells us
//lets make some modifications here
//should work the same way 
//lets see why it fails
//it did not update the definition as was expected lets fix that
//lets make some modifications to our add test to accomodate for the new update method
