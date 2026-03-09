package mapstests

import (
	"testing"
)

func TestSearch(t *testing.T){

	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func (t *testing.T){
		got, _ := dictionary.Search("test")
		
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	t.Run("unkown word", func (t *testing.T){
		_, got := dictionary.Search("unknown")

		if got == nil{
			t.Fatal("expected to get an error")

		}
		assertError(t, got, ErrorNotFound)
	})
}
func TestAdd(t *testing.T){
	t.Run("new word", func (t *testing.T){
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func (t *testing.T){
		
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")


		assertError(t, err, ErrorWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
	

}

func TestUpdate(t *testing.T){
	t.Run("existing word", func(t *testing.T){
		//this is for if the word is already existing in the map and we need to change the definition

		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T){
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		//it needs to be empty for it to be a new word 

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
	
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string){
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings (t testing.TB,got, want string){
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got , want error){
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
} 
//what we wanted to add something to our dictionary instead of just searching
//lets simplify you code abit
//what if the word we are looking for already exists
//lets try to test it first
//lets satisfy what the compiler wants so that the test runs even if it fails
//later we will find out why our test passed even when that was not our intention
//when we left oy test wasa still passing even when it wasn't supposed to
//lets fix the test
//lets see what the compiler tells us
//it says it got nil instead of the error message now lets fix the function
//we have no update method or function and so it should complain
//lets go make a function for it 
//there should be some complaints from the compiler
//we will deal with this later