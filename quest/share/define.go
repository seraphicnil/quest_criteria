package questshare

type ProgressSetType int32

const (
	ProgressSetTypeSet        ProgressSetType = 1 // 无论什么情况，都直接覆盖
	ProgressSetTypeAccumulate ProgressSetType = 2 // 一直累加
	ProgressSetTypeHighest    ProgressSetType = 3 // 取最高值
	ProgressSetTypeContinue   ProgressSetType = 4 // 非零值，一直累加 // 零值或负值，还原为零
)

type CriteriaType int32

const (
	CriteriaTypeExample1 CriteriaType = 1
	CriteriaTypeExample2 CriteriaType = 2
)

type BaseTaskCriteria struct {
	Type   int32
	Target int64
	Param  interface{}
}
