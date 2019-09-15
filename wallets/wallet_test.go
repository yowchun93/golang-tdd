package wallets

import "testing"

func TestWallet(t *testing.T) {
	// wallet := Wallet{}
	// wallet.Deposit(10)

	// got := wallet.Balance()
	// want := Bitcoin(10)

	// if got != want {
	// 	// t.Errorf("got %d want %d", got, want)
	// 	t.Errorf("got %s want %s", got, want)
	// }

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			// t.Errorf("got %d want %d", got, want)
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(20)
		wallet.Withdraw(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			// t.Errorf("got %d want %d", got, want)
			t.Errorf("got %s want %s", got, want)
		}
	})

	//  asserting Error is thrown
	t.Run("withdrawing without sufficient fund", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		if err == nil {
			t.Error("wanted an error but did not get one")
		}
	})
}
