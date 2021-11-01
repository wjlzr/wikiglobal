package utils

import "github.com/gin-gonic/gin"

// 获取客户端
func CountryCode(c *gin.Context) string {
	return c.Request.Header.Get("country-code")
}
