package mode

import "wiki_global/src/config"

//获取redis值，根据模式检查
func GetRedisKey(key string, isLine bool) string {
	if config.Conf().Application.Mode == "dev" || config.Conf().Application.Mode == "test" {
		if isLine {
			return key + "_test_"
		}
		return key + "_test"
	}
	return key
}
