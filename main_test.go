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
		{"1-2*3", -5},
		{"6/3", 2},
		{"let x = 3 in x+x end", 6},
		{"let x = 1+2+3 in let x = x*4 in x+x end + x end", 54},
		{"let x = 2 in let y = x in let x = 3 in x+y end + x end + x end", 9},
		{"let f(x) = x+2 in f(3) end", 5},
		{"let y = 5 in let f(x) = x+y in let y = 3 in f(y) end end end", 8},
		{"if true then 1 else 2", 1},
		{"let x = 2 in if x == 2 then 2 else 1 end", 2},
		{"let f(x) = if x == 0 then 0 else if x == 1 then 1 else 2 in f(2) end", 2},
		{"let f(x) = if x == 0 then 0 else if x == 1 then 1 else f(x-1) + f(x-2) in f(10) end", 55},
		{"(x -> x * x)(2)", 4},
	}

	for _, test := range tests {
		if got := eval(test.input); got != test.want {
			t.Errorf("calc(%q) = %d want %d", test.input, got, test.want)
		}
	}
}
