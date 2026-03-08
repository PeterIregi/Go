package mapstests

import(
	"errors"
)


type Dictionary map[string]string

var (
	ErrorNotFound = errors.New("Could not find the word you were looking for")
	ErrorWordExists = errors.New("cannot add word because it already exists")
)

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

//we can do something so that we declare multiple values with on far statement
//lets see now what the compiler tells us
