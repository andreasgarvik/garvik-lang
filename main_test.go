package main

import (
	"testing"
)

func TestEvaluator(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"1+1", 2},
		{"1+2+3", 6},
		{"1*2+3", 5},
		{"1*(2+3)", 5},
		{"1+2*3", 7},
		{"(1+2)*3", 9},
		{"let x = 3 in x+x end", 6},
		{"let x = 1+2+3 in let x = x*4 in x+x end + x end", 54},
		{"let x = 2 in let y = x in let x = 3 in x+y end + x end + x end", 9},
	}

	for _, test := range tests {
		if got := eval(test.input); got != test.want {
			t.Errorf("calc(%q) = %d want %d", test.input, got, test.want)
		}
	}
}
