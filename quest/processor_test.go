package quest

import (
	"quest/handler"
	questshare "quest/share"
)

type CriteriaProcessorExample struct {
	CriteriaList map[uint32]*Criteria
}

func (p *CriteriaProcessorExample) GetListByType(criteriaType questshare.CriteriaType) []*Criteria {
	result := make([]*Criteria, 0)
	for _, criteria := range p.CriteriaList {
		if criteria.GetCriteriaType() == criteriaType {
			result = append(result, criteria)
		}
	}
	return result
}
func (p *CriteriaProcessorExample) CanUpdate(criteria *Criteria, args ...interface{}) bool {
	return criteria.CanUpdate(args...)
}
func (p *CriteriaProcessorExample) SetProgress(id uint32, changeValue int64) int64 {
	criteria, ok := p.CriteriaList[id]
	if ok == false {
		return 0
	}
	count := criteria.BaseTaskHandler.Count()
	if count.Set(changeValue, criteria.GetProgressSetType()) {
		return count.Value()
	}
	return count.Value()
}
func (p *CriteriaProcessorExample) Complete(id uint32) {
	criteria, ok := p.CriteriaList[id]
	if ok == false {
		return
	}
	count := criteria.Count()
	if count.Value() != criteria.Config.Target {
		return
	}
	criteria.End = true
}

func (p *CriteriaProcessorExample) AddCriteria(c *Criteria) {
	p.CriteriaList[c.Id] = c
}

// CriteriaTest 累计，指定模式玩N局，见策划表
type CriteriaTest struct {
	questshare.CriteriaType
	questshare.ProgressSetType
	count questshare.Progress // 数量
}

func NewCriteriaTest(config questshare.BaseTaskCriteria) handler.BaseTaskHandler {
	return &CriteriaTest{
		CriteriaType:    questshare.CriteriaTypeExample2,
		ProgressSetType: questshare.ProgressSetTypeSet,
		count:           0,
	}
}

func (c *CriteriaTest) GetCriteriaType() questshare.CriteriaType {
	return c.CriteriaType
}
func (c *CriteriaTest) GetProgressSetType() questshare.ProgressSetType {
	return c.ProgressSetType
}
func (c *CriteriaTest) CanUpdate(args ...any) bool {
	// 获取参数
	if len(args) < 1 {
		return false
	}
	count := questshare.ParseNumber[int64](args[0])

	// 判断参数

	c.count.Set(count, c.ProgressSetType)
	return true
}

func (c *CriteriaTest) Count() questshare.Progress {
	return c.count
}

var processor = &CriteriaProcessorExample{
	CriteriaList: map[uint32]*Criteria{},
}

func ExampleNewCriteria() {
	// 添加任务
	testCriteriaId := uint32(11)                   //read from config
	testCriteriaType := questshare.CriteriaType(3) // user define it !
	criteriaTest := NewCriteria(testCriteriaId, questshare.BaseTaskCriteria{
		Type:   int32(testCriteriaType),
		Target: 10,  // read from config
		Param:  nil, //read from config
	})
	AddBuilder(testCriteriaType, NewCriteriaTest)
	processor.AddCriteria(criteriaTest)
}

func ExampleUpdate() {
	//user define for update and notify
	cb := func(dirtyCriteria []*Criteria) {
	}
	Update(processor, 3, cb, 11, 13)
}
