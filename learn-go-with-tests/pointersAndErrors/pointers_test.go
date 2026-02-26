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

	assertError := func(t testing.TB, got error, want string){
		t.Helper()
		if got == nil {
			t.Error("Wanted an error but didn't get one")
		}
		if got.Error()!= want {
			//Error() hepls us conver the error to a string so that we can compare it 
			//to our want string
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("deposit", func (t *testing.T){
		wallet :=Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("withdraw", func(t *testing.T){
		wallet :=Wallet{balance:Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T){
		startingBalance:= Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		assertError(t, err, "cannot withdraw, insufficient funds")
	})
}
//we can continue now 
//lets make the error make abit more sense and not just be presented as an error
//what will happen now
//now our error message is supposed to be abit more discriptive lets change the cofde 
//so it describes the error encountered