package basic

import (
	"testing"
)

func TestByteStrTobase64(t *testing.T) {
	var input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	var expected = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if ByteStrTobase64(input) != expected {
		t.Errorf(`basic.ByteStrTobase64("%s") != "%s"`, input, expected)
	}
}
