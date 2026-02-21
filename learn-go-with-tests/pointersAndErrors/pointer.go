package pointersAndErrors


type Wallet struct{
	balance int
}
func (w Wallet) Deposit(amount int){
	w.balance += amount
}
func (w Wallet) Balance()int{
	return w.balance
}
//lets run it now
//we need to add methods to it 
//now it fails because we have not met the terms stated in the test lets try to meet them
//lets try it now
//it still fails we will learn more on why later