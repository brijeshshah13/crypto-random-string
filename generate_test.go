package cryptorandomstring

import "testing"

func TestGenerate(t *testing.T) {
	if got := Generate(); got == "" {
		t.Errorf("Expected random string is empty")
	}
}
