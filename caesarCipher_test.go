package caesarCipher

import "testing"

func TestCaesarCipher(t *testing.T) {
	expected := "efgh"
	got := CaesarCipher("abcd")
	if expected != got {
		t.Errorf("got '%s', expected '%s'", string(got), string(expected))
	}
}

func TestCaesarCipher2(t *testing.T) {
	expected := "abcd"
	got := CaesarCipher("å,. ")
	if expected != got {
		t.Errorf("got '%s', expected '%s'", string(got), string(expected))
	}
}
