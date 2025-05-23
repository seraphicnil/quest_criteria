package questshare

import (
	"testing"
)

func TestIsAny(t *testing.T) {
	type testCase struct {
		In  string
		Out bool
	}
	var cases = []testCase{
		{"", true},
		{"0", true},
		{"Any", true},
		{"any", false},
	}
	for i, c := range cases {
		value := IsAny(c.In)
		if value != c.Out {
			t.Errorf("case #%d failed, expect=%t, actual=%t, in=[%s]", i+1, c.Out, value, c.In)
		}
	}
}

func TestIsStrContains(t *testing.T) {
	type testCase struct {
		In1 string
		In2 string
		Out bool
	}
	var cases = []testCase{
		{"", "", false},
		{"1", "1", true},
		{"some,str,in,here", "", false},
		{"some,str,in,here", "so", false},
		{"some,str,in,here", "some", true},
		{"some,str,in,here", "in", true},
		{"some,str,in,here", "here", true},
	}
	for i, c := range cases {
		value := IsStrContains(c.In1, c.In2)
		if value != c.Out {
			t.Errorf("case #%d failed, expect=%t, actual=%t, in1=[%s], in2=[%s]", i+1, c.Out, value, c.In1, c.In2)
		}
	}
}

func TestIsStrAnyContains(t *testing.T) {
	type testCase struct {
		In1 string
		In2 string
		Out bool
	}
	var cases = []testCase{
		{"", "", false},
		{"1", "1", true},
		{"some,str,in,here", "", false},
		{"some,str,in,here", "so", false},
		{"some,str,in,here", "some", true},
		{"some,str,in,here", "in", true},
		{"some,str,in,here", "here", true},
		{"Any", "here", true},
	}
	for i, c := range cases {
		value := IsStrAnyContains(c.In1, c.In2)
		if value != c.Out {
			t.Errorf("case #%d failed, expect=%t, actual=%t, in1=[%s], in2=[%s]", i+1, c.Out, value, c.In1, c.In2)
		}
	}
}

func TestIsSliceContains(t *testing.T) {
	type testCase struct {
		In1 []string
		In2 string
		Out bool
	}
	var cases = []testCase{
		{[]string{"some", "str", "in", "here"}, "here", true},
		{[]string{"some", "str", "in", "here"}, "out", false},
	}
	for i, c := range cases {
		value := IsSliceContains(c.In1, c.In2)
		if value != c.Out {
			t.Errorf("case #%d failed, expect=%t, actual=%t, in1=[%s], in2=[%s]", i+1, c.Out, value, c.In1, c.In2)
		}
	}
}
