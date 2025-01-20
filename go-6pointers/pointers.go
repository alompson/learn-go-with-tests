package pointers

import "fmt"
import "errors"

type Bitcoin int

type Stringer interface{
	String() string
}

func (b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) GetBalance() Bitcoin{
	return w.balance
}

func (w *Wallet) Deposit(value Bitcoin){
	w.balance += value
	//fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(value Bitcoin) error{
	if(value > w.balance) {
		return ErrInsufficientFunds

	}
	w.balance -= value
	return nil
}