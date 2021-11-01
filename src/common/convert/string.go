package convert

import "strconv"

func IntToString(e int) string {
	return strconv.Itoa(e)
}

func Int64ToString(e int64) string {
	return strconv.FormatInt(e, 10)
}
