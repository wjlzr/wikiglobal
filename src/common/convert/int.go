package convert

import "strconv"

//字符串转换成int64
func StrToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

//字符串转换成int32
func StrToInt32(str string) int32 {
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0
	}
	return int32(i)
}

//字符串转换int
func StrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}
