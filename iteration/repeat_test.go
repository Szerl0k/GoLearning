package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Got %s, expected: %s", repeated, expected)
	}
}

func TestStrings(t *testing.T) {

	t.Run("Removing redundant spaces from string", func(t *testing.T) {
		got := removeSpaces("Ala  ma    kota , a  kot ma      Ale")
		want := "Ala ma kota, a kot ma Ale"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Checking if string is a palindrome", func(t *testing.T) {
		got := isPalindrome("TACO CAT")
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("Checking if string is a palindrome", func(t *testing.T) {
		got := isPalindrome("sasasasasaagfgfhbffb")
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("Checking if string is a palindrome", func(t *testing.T) {
		got := isPalindrome("KAYAK")
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 6)
	fmt.Println(repeated)
	//Output: aaaaaa
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
