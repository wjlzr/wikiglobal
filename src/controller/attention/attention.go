package attention

import (
	"wiki_global/src/common/convert"
	"wiki_global/src/global/i18nresponse"
	"wiki_global/src/global/response"
	"wiki_global/src/models/mysql/attention"
	"wiki_global/src/utils"

	"github.com/gin-gonic/gin"
)

// Follow 关注操作
func Follow(c *gin.Context) {

	var receiver attention.Attention
	if err := c.ShouldBindJSON(&receiver); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	result, err := receiver.Follow(c)
	if err != nil {
		i18nresponse.Error(c, "1010011")
		return
	}

	i18nresponse.Success(c, "ok", result)
}

// CancelFollow 取消关注
func CancelFollow(c *gin.Context) {

	var receiver attention.Attention
	if err := c.ShouldBindJSON(&receiver); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	result, err := receiver.CancelFollow(c)
	if err != nil {
		i18nresponse.Error(c, "1010012")
		return
	}

	i18nresponse.Success(c, "ok", result)
}

// MyFollow 我的关注列表
func MyFollow(c *gin.Context) {

	var myFollowRequest attention.MyFollowRequest
	if err := c.ShouldBindJSON(&myFollowRequest); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	result, count, err := myFollowRequest.MyFollow(c)
	if err != nil {
		i18nresponse.Error(c, "1010013")
		return
	}

	response.PageOK(c, result, count, myFollowRequest.PageIndex, myFollowRequest.PageSize, "ok")
}

// CountByFollow 我的关注公司/高管数量
func CountByFollow(c *gin.Context) {

	var a attention.Attention
	a.UserId = convert.Int64ToString(utils.GetUserId(c))
	count, err := a.CountByFollow()
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", count)
}
