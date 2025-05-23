package quest

import (
	"fmt"
	"quest/handler"
	"quest/share"
)

type Criteria struct {
	handler.BaseTaskHandler
	Id     uint32
	Config questshare.BaseTaskCriteria
	End    bool
	Reward bool
}

func NewCriteria(id uint32, config questshare.BaseTaskCriteria) *Criteria {
	c := &Criteria{
		Id:     id,
		Config: config,
	}
	// 根据任务类型构建handler
	if builder, ok := builders[questshare.CriteriaType(config.Type)]; ok {
		c.BaseTaskHandler = builder(config)
	} else {
		fmt.Println("quest", "handler not found. id:", id, "%d config.Type:", config.Type)
	}
	return c
}

type TaskHandlerBuilder func(config questshare.BaseTaskCriteria) handler.BaseTaskHandler

var builders = map[questshare.CriteriaType]TaskHandlerBuilder{
	questshare.CriteriaTypeExample1: handler.NewCriteriaExample1,
	questshare.CriteriaTypeExample2: handler.NewCriteriaExample2,
}

func AddBuilder(t questshare.CriteriaType, builder TaskHandlerBuilder) bool {
	if _, ok := builders[t]; ok {
		return false
	}
	builders[t] = builder
	return true
}
