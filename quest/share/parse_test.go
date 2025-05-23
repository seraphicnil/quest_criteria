package questshare

import (
	"testing"
)

func TestParseNumber(t *testing.T) {
	type testCase struct {
		In  any
		Out int
	}
	var cases = []testCase{
		{0, 0},
		{uint64(1), 1},
	}
	for i, c := range cases {
		value := ParseNumber[int](c.In)
		if value != c.Out {
			t.Errorf("case #%d failed, expect=%v, actual=%v, in=[%d]", i+1, c.Out, value, c.In)
		}
	}
}

func TestParseString(t *testing.T) {
	type testCase struct {
		In  any
		Out string
	}
	var cases = []testCase{
		{0, "0"},
		{uint64(1), "1"},
		{"3", "3"},
	}
	for i, c := range cases {
		value := ParseString(c.In)
		if value != c.Out {
			t.Errorf("case #%d failed, expect=%v, actual=%v, in=[%d]", i+1, c.Out, value, c.In)
		}
	}
}

func TestParseIntSlice(t *testing.T) {
	type testCase struct {
		In  string
		Out []int
	}
	var cases = []testCase{
		{"0", []int{0}},
		{"1,2,3", []int{1, 2, 3}},
	}
	for i, c := range cases {
		value := ParseIntSlice(c.In)
		containsAll := true
		for _, v := range c.Out {
			if !IsSliceContains(value, v) {
				containsAll = false
				break
			}
		}
		if len(value) != len(c.Out) || !containsAll {
			t.Errorf("case #%d failed, expect=%v, actual=%v, in=[%s]", i+1, c.Out, value, c.In)
		}
	}
}
