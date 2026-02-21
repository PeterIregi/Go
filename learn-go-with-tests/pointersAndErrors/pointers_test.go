//lets learn something new  or old since we have been using it from the beginning
package pointersAndErrors

import(
	"testing"
)

func TestWallet(t *testing.T){
	wallet :=Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	if got  != want{
		t.Errorf("got %d want %d", got, want)
	}
}
//here we want a program that helps us add some coins to ou wallet and check the balance
//lets first run the test and let the compiler tell us what we need before we continue
// lets go define our functions or methods or structs