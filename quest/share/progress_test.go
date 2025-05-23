package questshare

import (
	"testing"
)

func TestProgress(t *testing.T) {
	var p Progress
	var targetValue = int64(1)
	type testCase struct {
		In1 int64
		In2 ProgressSetType
		Out bool
	}
	var cases = []testCase{
		{0, ProgressSetTypeSet, false},
		{-1, ProgressSetTypeSet, true},
		{2, ProgressSetTypeAccumulate, true},
		{3, ProgressSetTypeAccumulate, true},
		{4, ProgressSetTypeHighest, false},
		{6, ProgressSetTypeHighest, true},
		{0, ProgressSetTypeContinue, true},
		{1, ProgressSetTypeContinue, true},
	}
	for i, c := range cases {
		value := p.Set(c.In1, c.In2)
		if value != c.Out {
			t.Errorf("case #%d failed, expect=%t, actual=%t, in1=[%d], in2=[%d]", i+1, c.Out, value, c.In1, c.In2)
		}
	}
	if p.Value() != targetValue {
		t.Errorf("final value check failed, expect=%d, actual=%d", targetValue, p.Value())
	}
}
