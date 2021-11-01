package middleware

import (
	"net/http"
	"wiki_global/src/common/constant"
	"wiki_global/src/utils"
	"wiki_global/src/utils/token"

	"github.com/gin-gonic/gin"
)

//用户授权中间件
func UserAuthMiddleware(token *token.Token, skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		//跳过路由不检测授权
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		if t := utils.GetToken(c); t != "" {
			if userId := utils.ParseUserID(token, c, t); userId > 0 {
				c.Set(constant.UserId, userId)
				c.Next()
				return
			} else {
				utils.ResponseErrorJson(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
				return
			}
		}

		utils.ResponseErrorJson(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	}
}
