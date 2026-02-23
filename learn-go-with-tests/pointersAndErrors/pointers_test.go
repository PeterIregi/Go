package pointersAndErrors

import(
	"testing"
)

func TestWallet(t *testing.T){

	t.Run("deposit", func (t *testing.T){
		wallet :=Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got  != want{
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T){
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

	
//lets see what the compiler says could be the issue 
//withdraw is not a method in our program
