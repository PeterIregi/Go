package pointersAndErrors

import(
	"fmt"
	"errors"
)

type Bitcoin int

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
		return errors.New("cannot withdraw, insufficient funds")
		//this lets us create a new error with a message of our choosing
	}
	w.balance -= amount
	return nil
}
//our method for withdraw needs to return an error currently it has no return value
//we can import the error package and do some more interesting things with our error in terms of the message lets see
//lets make our test more readable y creating a test helper for our error check
//we will learn more about error handling bit by bit 