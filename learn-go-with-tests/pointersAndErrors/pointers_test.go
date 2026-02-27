package pointersAndErrors

import(
	"testing"
)

func TestWallet(t *testing.T){
	assertBalance := func(t testing.TB,wallet Wallet, want Bitcoin){
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want error){
		t.Helper()
		if got == nil {
			t.Error("Wanted an error but didn't get one")
		}
		if got != want {
			//Error() hepls us conver the error to a string so that we can compare it 
			//to our want string
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertNoError := func (t testing.TB, got error){
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	t.Run("deposit", func (t *testing.T){
		wallet :=Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("withdraw", func(t *testing.T){
		wallet :=Wallet{balance:Bitcoin(20)}
		err:= wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		//now in our succeffull withdraw we can assert that there was no error thrown
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T){
		startingBalance:= Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		assertError(t, err, ErrInsufficientFunds)
	})
}
//no error message form the compiler so we are all good
//lets see if our test still passes
//it says there is a type mismatch lets fix that in our test 
//now all the types are matched 
//but what if there should be no error we have't checked for that 
//what if the withdraw was successful 
//lets do that 
//lets try to run it 
//that wraps it up for the series now we move on to something else 