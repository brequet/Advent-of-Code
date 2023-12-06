package main

import "testing"

type partATest struct {
	input    string
	expected int
}

var partATests = []partATest{
	{"(())", 0},
	{"()()", 0},
	{"(((", 3},
	{"(()(()(", 3},
	{"))(((((", 3},
	{")", -1},
	{"))", -2},
	{"))(", -1},
	{")())())", -3},
}

func TestPartA(t *testing.T) {
	for _, tt := range partATests {
		t.Run(tt.input, func(t *testing.T) {
			result := partA(tt.input)
			if result != tt.expected {
				t.Errorf("partA(%s) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
