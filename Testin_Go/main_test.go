package main

import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 +2 to equeal 4")
	}
}

func TestTableCaculate(t *testing.T) {
	var tests = []struct {
		input   int
		expeted int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{25, 27},
		{7, 9},
	}
	for _, test := range tests {
		if output := Calculate(test.input); output != test.expeted {
			t.Error("Test failed {} inputted, {} expected, recived {}", test.input, test.expeted, output)
		}
	}

}
