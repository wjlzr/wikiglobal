package dsl

import "encoding/json"

func Got(src interface{}) string {
	data, _ := json.Marshal(src)
	got := string(data)
	return got
}
