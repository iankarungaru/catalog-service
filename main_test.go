package main

import "testing"

func TestApplyDiscount(t *testing.T) {
	result := ApplyDiscount(100, 10)
	expected := 90.0
	if result != expected {
		t.Errorf("expected %.2f, got %.2f", expected, result)
	}
}