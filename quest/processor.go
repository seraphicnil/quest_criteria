package quest

import (
	"quest/share"
)

type CriteriaProcessor interface {
	GetListByType(criteriaType questshare.CriteriaType) []*Criteria
	CanUpdate(criteria *Criteria, args ...interface{}) bool
	SetProgress(id uint32, changeValue int64) int64
	Complete(id uint32)
}

// Update 更新任务进度，由Account的通用UpdateTask接口调用
func Update(processor CriteriaProcessor, criteriaType questshare.CriteriaType, cb func(dirtyCriteria []*Criteria), args ...interface{}) {
	criteriaList := processor.GetListByType(criteriaType)
	for _, criteria := range criteriaList {
		if !processor.CanUpdate(criteria, args...) {
			continue
		}
		count := criteria.Count()
		final := processor.SetProgress(criteria.Id, count.Value())
		if final >= criteria.Config.Target {
			processor.Complete(criteria.Id)
		}
	}
	cb(criteriaList)
}
