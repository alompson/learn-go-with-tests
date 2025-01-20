package pointers

import "testing"
//import "fmt"



func checkValue(t testing.TB, wallet Wallet, want Bitcoin){
	t.Helper()
	got := wallet.GetBalance()
	if got != want{
		t.Errorf("got %s but want %s", got, want)
	}
}

func TestWallet(t *testing.T){
	assertError := func(t testing.TB, got, want error) {
		t.Helper()
	
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
	
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T){
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
	
		want := Bitcoin(10)	
		checkValue(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T){
		wallet:= Wallet{balance: Bitcoin(10)}

		wallet.Withdraw(Bitcoin(5))

		want := Bitcoin(5)

		checkValue(t, wallet, want)

	})

	t.Run("Withdraw withour enough funds", func(t *testing.T){
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		
		checkValue(t, wallet, startingBalance)

		assertError(t, err, ErrInsufficientFunds)
	})
}