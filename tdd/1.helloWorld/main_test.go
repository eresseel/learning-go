// Go-ban szokás és ajánlás, hogy a tesztfájlok (pl. main_test.go) ugyanabban a mappában
// és csomagban legyenek, mint a tesztelt kód.
// Ez azért praktikus, mert így a tesztek közvetlenül elérik a nem exportált (private)
// függvényeket és típusokat is.
package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
