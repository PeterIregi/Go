//today we will learn about maps
//they allow us to store item in a similar way to how they ae stored in dictionaries
//with the word as the j=key and the definition as the value ..
// so in key value pairs
//lets try to make a dictionary with this knowledge to  make us understand  better
package mapstests

import (
	"testing"
)

func TestSearch(t *testing.T){
	dictionary := map[string] string{"test": "this is just a test"}
	//lets just keep going we will explain the map [string] syntax as we get further 

	got := Search(dictionary, "test")
	want := "this is just a test"

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}
// map[string].. is how we declare a map with strings inside
//we state the {"test":which is the key type and the "this is just a test" is the value type}
//lets make our search function