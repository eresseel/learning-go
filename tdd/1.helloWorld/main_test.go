// Go-ban szokás és ajánlás, hogy a tesztfájlok (pl. main_test.go) ugyanabban a mappában
// és csomagban legyenek, mint a tesztelt kód.
// Ez azért praktikus, mert így a tesztek közvetlenül elérik a nem exportált (private)
// függvényeket és típusokat is.
package main

import (
	"testing" // origin

	"github.com/stretchr/testify/require"
)

// origin
// func TestGetMessage(t *testing.T) {
// 	got := GetMessage()
// 	want := "Hello World"

// 	if got != want {
// 		t.Fatalf("GetMessage() = %q, want %q", got, want)
// 	}
// }

// with testify
func TestGetMessage(t *testing.T) {
	got := GetMessage()
	require.Equal(t, "Hello World", got)
}
