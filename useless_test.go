package main

import "testing"

func TestAddFive(t *testing.T) {
	fiveMore := addFive(12)
	if fiveMore != 17 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", fiveMore, 17)
	}
}
