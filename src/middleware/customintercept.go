package middleware

import (
	"github.com/gin-gonic/gin"
)

// 自定义参数拦截
func CustomIntercept() gin.HandlerFunc {

	return func(c *gin.Context) {

		lang := c.Request.Header.Get("Language")
		if lang == "en" {
			c.Request.Header.Set("country-code", "840")
		} else {
			c.Request.Header.Set("country-code", "156")
		}
		c.Next()
	}
}
