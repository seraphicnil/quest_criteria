package handler

import (
	questshare "quest/share"
	"testing"
)

func TestCriteriaExample1(t *testing.T) {
	config := questshare.BaseTaskCriteria{
		Target: 3,
		Param:  map[int32]int32{1: 2},
	}
	criteria := NewCriteriaExample1(config)
	type testCase struct {
		Out bool
	}
	var cases = []testCase{}
	for i, c := range cases {
		value := criteria.CanUpdate(1)
		if value != c.Out {
			t.Errorf("case #%d failed, expect=%t, actual=%t", i+1, c.Out, value)
		}
	}
	count := criteria.Count()
	if config.Target != count.Value() {
		t.Errorf("progress check failed, expect=%d, actual=%d", config.Target, count)
	}
}
