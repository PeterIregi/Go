package main
 
import(
	"testing"
	
)

//Today we will try and see how errors works in go
func TestDivide(t *testing.T){
	result, err := Divide(10, 2)
//displays the error we got if any
	if err != nil{
		t.Fatalf("did not expect an error, got %v", err)
	}
	if result != 5{
		t.Errorf("expexted 5, Got %d", result)
	}
}
func TestDivideByZero(t *testing.T){
	_, err := Divide(10, 0)
	//expects an error and displays message if none is thrown
	if err == nil{
		t.Fatalf("expected an error, but got nil")
	}

}
//should fail because we have not made the divide function lets see then create the functio
//it fails cause divide is undefinedlest make it
//was starting to get worried
//lets test for diividing by zero
//lets test it
//we will learn how to deal with this one later 