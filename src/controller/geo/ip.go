package geo

import (
	"wiki_global/src/common/geo"
	"wiki_global/src/global/i18nresponse"

	"github.com/gin-gonic/gin"
)

//GetWithIpToLocation 根据ip判断是否为大陆
func GetWithIpToLocation(c *gin.Context) {
	var isChina bool
	code := geo.GetCountryCode(c.ClientIP())
	if code == "CN" {
		isChina = true
	}
	i18nresponse.Success(c, "ok", struct {
		IsChina bool `json:"is_china"`
	}{IsChina: isChina})
}
