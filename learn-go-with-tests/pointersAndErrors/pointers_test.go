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

	assertError := func(t testing.TB, err error){
		t.Helper()
		if err == nil {
			t.Error("Wanted an error but didn't get one")
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

		assertError(t, err)
	})
}
//lets see how we would deal with the case where we want to withdraw more than waht is in our balance 
//it should give us an error message right lets try that 
//lets see what the compiler says 
//its says our withdraw method returns no value yet we are asking for an error
//lets fix that
//lets close it of at that 
	
