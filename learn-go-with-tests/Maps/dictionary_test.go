package mapstests

import (
	"testing"
)

func TestSearch(t *testing.T){

	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func (t *testing.T){
		got, _ := dictionary.Search("test")
		//the _ means the method should return two things but we only need one so 
		//_just means ignore the other one and give me the one i have requested
		//this is because in Go we cannot declare a variable and not use it
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	//now we expect some complaints from the compiler cause we only return one value
	//from our method
	//lets just make sure
	t.Run("unkown word", func (t *testing.T){
		_, err := dictionary.Search("unknown")
		//now we have no use for the first value hence the _ at the beginning
		want := "could not find the word you are looking for"

		if err == nil{
			//this just says if there is no error message of which should be wrong
			t.Fatal("expected to get an error")

		}
		assertStrings(t, err.Error(), want)
	})
}
func assertStrings (t testing.TB,got, want string){
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
//lets continue
//lets first modify our code so that we reduce repetition in future
//now lets make a dictionary type to wrap our map and add a search method to it but first the test
//lets go define dictionary in our dictionary .go
//lets try to fix that