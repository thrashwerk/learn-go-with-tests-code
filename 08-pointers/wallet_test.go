package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Shitcoin(10))

		assertBalance(t, wallet, Shitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Shitcoin(20)}

		err := wallet.Withdraw(Shitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Shitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Shitcoin(20)
		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Shitcoin(100))

		assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, w Wallet, want Shitcoin) {
	t.Helper()

	got := w.Balance()

	if got != want {
		t.Errorf("got: %s; want: %s", got, want)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()

	if got == nil {
		t.Fatal("expected an error but didn't get one")
	}

	if got.Error() != want {
		t.Errorf("got: %q; want: %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("got an error but didn't expect one")
	}
}
