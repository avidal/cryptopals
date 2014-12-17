package tests

import (
	"../cryptopals"
	"testing"
)

func TestFixedXOR(t *testing.T) {
	var tests = []struct {
		a    string
		b    string
		want string
	}{
		{
			"1c0111001f010100061a024b53535009181c",
			"686974207468652062756c6c277320657965",
			"746865206b696420646f6e277420706c6179",
		},
	}

	for _, tt := range tests {
		got, err := cryptopals.FixedXORString(tt.a, tt.b)
		if err != nil {
			t.Errorf("got err %v", err)
		}
		if string(got) != tt.want {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
