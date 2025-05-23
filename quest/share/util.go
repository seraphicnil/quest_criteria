package questshare

import "strconv"

func Str2Int64(s string) int64 {
	ret, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}

	return ret
}
