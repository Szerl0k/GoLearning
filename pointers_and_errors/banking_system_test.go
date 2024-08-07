package pointers_and_errors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	/*
		assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
			t.Helper()
			got := wallet.Balance()

			if got != want {
				t.Errorf("got %s want %s", got, want)
			}
		}
	*/
	/*
		assertError := func(t testing.TB, got, want error) {
			t.Helper()
			if got == nil {
				t.Fatal("wanted an error but didn't get one")
			}

			if got.Error() != want.Error() {
				t.Errorf("got %q, want %q", got, want)
			}
		}
	*/
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		//fmt.Printf("address of balance in the test is %p \n\n", &wallet.balance)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw inefficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)

	})

}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
