package questshare

type Progress int64

func (v *Progress) Value() int64 {
	return int64(*v)
}

func (v *Progress) Set(changeValue int64, setType ProgressSetType) bool {
	var newValue int64
	switch setType {
	case ProgressSetTypeSet:
		// 无论什么情况，都直接覆盖
		newValue = changeValue
	case ProgressSetTypeAccumulate:
		// 一直累加
		newValue = v.Value() + changeValue
	case ProgressSetTypeHighest:
		// 取最高值
		newValue = v.Value()
		if newValue < changeValue {
			newValue = changeValue
		}
	case ProgressSetTypeContinue:
		// 非零值，一直累加
		// 零值或负值，还原为零
		if changeValue > 0 {
			newValue = v.Value() + changeValue
		} else {
			newValue = 0
		}
	}
	// 如果值没有变化，返回false代表无变化
	if newValue == v.Value() {
		return false
	}
	// 保底策略，保证进度最小值为零
	if newValue < 0 {
		newValue = 0
	}
	*v = Progress(newValue)
	return true
}
