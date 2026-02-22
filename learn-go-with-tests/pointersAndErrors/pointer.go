package pointersAndErrors


type Wallet struct{
	balance int
}
func (w *Wallet) Deposit(amount int){
	w.balance+= amount
}
func (w *Wallet) Balance()int{
	return w.balance
}
//lets run it now
//we need to add methods to it 
//now it fails because we have not met the terms stated in the test lets try to meet them
//lets try it now
//it still fails we will learn more on why later
//we left it off here  
// why does it say got 0 instead of 10
//that is because as of now out w.balance refers to a copy of the w.balance hence the one in our function is still 0
//to change that lets do some modifications to our code
//lets see if if it passes now

// *Wallet is a pointer to the walletwe created so when we chage w. balance the balance now changes since it 
//is not creating a copy of the w.balance but pointing to the w.balane 
//this looks familiar we have been using it since the beginning in all our tests as *testing.T
//slowly things are starting to get pieced together the pieces are all starting to fall into place
//we will continue later 