package wallets

import "testing"

func TestWallet(t *testing.T) {
	assertError := func(t *testing.T, err error, want error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}

		if err != want {
			t.Errorf("got %q, want %q", err, want)
		}
	}

	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)

		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(30))

		assertError(t, err, ErrInsufficientFunds)
	})
}
