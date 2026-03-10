package mapstests

import(
	//"errors"
)


type Dictionary map[string]string

const (
	ErrorNotFound = DictionaryErr("Could not find the word you were looking for")
	ErrorWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exits")
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

func (d Dictionary) Update (word, definition string) error{
	_, err := d.Search(word)

	switch err{
	case ErrorNotFound:
		return ErrWordDoesNotExist
	case nil:
		 d[word] = definition
	default:
		return err
	}


	return nil
}
func (d Dictionary) Delete(word string) error{
	_, err := d.Search(word)

	switch err{
	case ErrorNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil

}
//we can do something so that we declare multiple values with on far statement
//lets see now what the compiler tells us
//lets make some modifications here
//should work the same way 
//lets see why it fails
//it did not update the definition as was expected lets fix that
//lets make some modifications to our add test to accomodate for the new update method
// it should run now but still fail.. or should it
//the test fails but thats coz the return isn't what its supposed to be
//lets see if it passes now
//lets try to create a delete function ...but first the test
//it should compile but fail

//try to make it pass
//what if the word we are trying to delete does not exist
//lets see if it passes