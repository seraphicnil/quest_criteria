package questshare

import (
	"fmt"
	"strconv"
	"strings"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// ParseNumber 将数字转为定义好的其他类型
//
//	例：var count int64 = ParseNumber[int64](args[0])
//	var gameMode share.GameMode = ParseNumber[share.GameMode](args[1])
func ParseNumber[T Number](v any) T {
	val := Str2Int64(fmt.Sprintf("%d", v))
	return T(val)
}

// ParseString 将任意类型转为String类型
//
//	例：0 -> "0", 1.2 -> "1.2", "sum" -> "sum"
func ParseString(v any) string {
	switch val := v.(type) {
	case string:
		return val
	default:
		return fmt.Sprintf("%v", v)
	}
}

// ParseIntSlice 将配置里的数字数组字段转化为判断用的slice字段
//
//	例：param="1,2" => []int{1,2}
func ParseIntSlice(param string) []int {
	paramStr := strings.Split(param, ",")
	var result []int
	for _, str := range paramStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("quest", "parseGameResultTypeSlice param: ", param)
			continue
		}
		result = append(result, num)
	}
	return result
}
