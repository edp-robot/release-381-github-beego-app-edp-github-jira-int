package controllers

import "testing"

func TestHello(t *testing.T) {
	msg := getMessage()
	if msg != "Hello, World!" {
		t.Errorf("expected \"Hello, World!\", but got %v", msg)
	}
}
