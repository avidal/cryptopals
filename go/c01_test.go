package cryptopals

import (
    "testing"
)

func TestHexToBase64(t *testing.T) {
    inp := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

    got, err := HexToBase64(inp)
    if err != nil {
        t.Errorf("got error running test")
    }

    if want != got {
        t.Errorf("wanted %v, got %v", want, got)
    }
}
