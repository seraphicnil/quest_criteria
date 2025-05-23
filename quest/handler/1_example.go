package handler

import (
	"quest/share"
)

// CriteriaExample1 累计，指定模式玩N局，见策划表
type CriteriaExample1 struct {
	questshare.CriteriaType
	questshare.ProgressSetType
	count questshare.Progress // 数量
	Param map[int32]int32
}

func NewCriteriaExample1(config questshare.BaseTaskCriteria) BaseTaskHandler {
	c := &CriteriaExample1{
		CriteriaType:    questshare.CriteriaTypeExample1,
		ProgressSetType: questshare.ProgressSetTypeAccumulate,
		count:           0,
		Param:           config.Param.(map[int32]int32),
	}
	return c
}

func (c *CriteriaExample1) GetCriteriaType() questshare.CriteriaType {
	return c.CriteriaType
}
func (c *CriteriaExample1) GetProgressSetType() questshare.ProgressSetType {
	return c.ProgressSetType
}
func (c *CriteriaExample1) CanUpdate(args ...any) bool {
	// 获取参数
	if len(args) < 3 {
		return false
	}

	// 判断参数
	key := args[1].(int32)
	val := args[2].(int32)
	if c.Param[key] != val {
		return false
	}

	count := questshare.ParseNumber[int64](args[0])
	c.count.Set(count, c.ProgressSetType)
	return true
}

func (c *CriteriaExample1) Count() questshare.Progress {
	return c.count
}

// CriteriaExample2 累计，指定模式玩N局，见策划表
type CriteriaExample2 struct {
	questshare.CriteriaType
	questshare.ProgressSetType
	count  questshare.Progress // 数量
	config questshare.BaseTaskCriteria
}

func NewCriteriaExample2(config questshare.BaseTaskCriteria) BaseTaskHandler {
	return &CriteriaExample2{
		CriteriaType:    questshare.CriteriaTypeExample2,
		ProgressSetType: questshare.ProgressSetTypeSet,
		count:           0,
		config:          config,
	}
}

func (c *CriteriaExample2) GetCriteriaType() questshare.CriteriaType {
	return c.CriteriaType
}
func (c *CriteriaExample2) GetProgressSetType() questshare.ProgressSetType {
	return c.ProgressSetType
}
func (c *CriteriaExample2) CanUpdate(args ...any) bool {
	// 获取参数
	if len(args) < 2 {
		return false
	}
	// 判断参数
	if c.config.Param != args[1] {
		return false
	}

	count := questshare.ParseNumber[int64](args[0])
	c.count.Set(count, c.ProgressSetType)
	return true
}

func (c *CriteriaExample2) Count() questshare.Progress {
	return c.count
}
