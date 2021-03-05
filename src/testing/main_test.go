package main

import "testing"

type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "Negative"},
	{0, "Zero"},
	{20, "big"},
	{800, "huge"},
	{5, "small"},
	{2000, "Enormus"},
}

func TestSize(t *testing.T) {
	for i, test := range tests {
		size := checkSize(test.in)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}
