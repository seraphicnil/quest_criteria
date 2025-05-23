package quest

import (
	questshare "quest/share"
	"testing"
)

func TestEvents(t *testing.T) {
	c1 := questshare.BaseTaskCriteria{
		Type:   int32(questshare.CriteriaTypeExample1),
		Target: 10,
		Param:  map[int32]int32{1: 2},
	}
	NewCriteria(uint32(1), c1)
	c2 := questshare.BaseTaskCriteria{
		Type:   int32(questshare.CriteriaTypeExample2),
		Target: 10,
		Param:  3,
	}
	NewCriteria(uint32(1), c2)
}
