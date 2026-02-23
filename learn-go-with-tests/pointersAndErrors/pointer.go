package pointersAndErrors

import(
	"fmt"
)

type Bitcoin int
//in Go we can create a type from an existing type 
//now we will use  stringer  so that we can  manipulate how it 
//displays when we print it out in format strings and use  %s
func (b Bitcoin) String()string{
	return fmt.Sprintf("%d BTC", b)
	//we have to import the fmt package in order to use fmt.
	//now lets modify our tests so that it prints out what we want
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
func (w *Wallet) Withdraw(amount Bitcoin){
	w.balance-= amount
}

//now lets see if our wallet could contain bitcoins 
//since we are currently just using int
//lets make the test fail so we see if it prints out what we want
//stringer works
//what we wanted to withdraw
//lets make some test for that 
//lets write enough code so that it compiles
//lets see if it passes
//later we deal with a situationwhere the withdrawn amount is larger than the balance 
