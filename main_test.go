package security

import "testing"

func TestSHA1(t *testing.T) {
	var expected = "a9993e364706816aba3e25717850c26c9cd0d89d"
	if id := SHA1("abc"); id != expected {
		t.Errorf("Expected %v to equal %v", id, expected)
	}
}
