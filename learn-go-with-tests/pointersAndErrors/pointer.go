package pointersAndErrors

import(
	"fmt"
	"errors"
)

type Bitcoin int

var ErrInsufficientFunds = errors.New("Cannot withdraw, insufficient funds")

func (b Bitcoin) String()string{
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct{
	balance Bitcoin
}
func (w *Wallet) Deposit(amount Bitcoin){
	w.balance+= amount
}
func (w *Wallet) Balance()Bitcoin{
	return w.balance
}
func (w *Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance{
		return ErrInsufficientFunds
		//now we can return our varriable here
		//this lets us create a new error with a message of our choosing
	}
	w.balance -= amount
	return nil
}
//lets define our error as a variable because in go we can do that and for reusability
//lets try to use it in our test also
