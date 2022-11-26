package base91

import (
	"testing"
)

func TestDecode(t *testing.T) {
	decoded := Decode("fPNKd")
	if string(decoded) != "test" {
		t.Errorf("Expected %s, got %s", "test", decoded)
	}
}

func TestEncode(t *testing.T) {
	encoded := Encode([]byte("test"))
	if encoded != "fPNKd" {
		t.Errorf("Expected %s, got %s", "fPNKd", encoded)
	}
}
