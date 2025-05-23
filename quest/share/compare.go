package questshare

import (
	"strings"
)

// IsAny 表明这个配置参数不做限制，可以取任意值
//
//	例：
//	1. param="" => true
//	2. param="0" => true
//	3. param="Any" => true
func IsAny(param string) bool {
	if param == "" || param == "0" || param == "Any" {
		return true
	}
	return false
}

// IsStrContains 判断配置中逗号分割的字符串参数是否包含目标值
//
//	例：param="some,str,in,here"
//	1. arg="so" => false
//	2. arg="some" => true
func IsStrContains(param, arg string) bool {
	strParam := strings.Split(param, ",")
	for _, str := range strParam {
		str = strings.TrimSpace(str)
		if str == "" {
			continue
		}
		if str == arg {
			return true
		}
	}
	return false
}

// IsStrAnyContains 判断配置中逗号分割的字符串参数是否包含Any或者目标值
//
//	例：param="Any" 或 param="some,str,in,here"
//	1. arg="some" => true
//	2. arg="str" => true
func IsStrAnyContains(param, arg string) bool {
	strParam := strings.Split(param, ",")
	for _, str := range strParam {
		str = strings.TrimSpace(str)
		if str == "" {
			continue
		}
		if str == "Any" || str == arg {
			return true
		}
	}
	return false
}

// IsSliceContains 判断Slice类型的变量是否包含某个元素
//
//	例：elems=[]share.GameCategory{share.GameCategoryRank,share.GameCategoryDuel}
//	1. v=share.GameCategoryRank => true
//	2. v=share.GameCategoryRoom => false
func IsSliceContains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

// IsNumbersMustContains 判断NumberSlice类型的变量必须包含某个元素
//
//	例：elems=[]share.GameCategory{share.GameCategoryRank,share.GameCategoryDuel}
//	1. v=share.GameCategoryAny => false
//	2. v=share.GameCategoryRank => true
//	3. v=share.GameCategoryRoom => false
func IsNumbersMustContains[T Number](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

// IsNumbersAnyContains 判断NumberSlice类型的变量是否包含0或者某个元素
//
//	例：elems=[]share.GameCategory{share.GameCategoryRank,share.GameCategoryDuel}
//	1. v=share.GameCategoryAny => true
//	2. v=share.GameCategoryRank => true
//	3. v=share.GameCategoryRoom => false
func IsNumbersAnyContains[T Number](elems []T, v T) bool {
	for _, s := range elems {
		if s == T(0) || s == v {
			return true
		}
	}
	return false
}
