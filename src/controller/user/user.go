package user

import (
	"wiki_global/src/common/convert"
	"wiki_global/src/global/i18nresponse"
	"wiki_global/src/models/mysql/oauth"
	"wiki_global/src/services/usercenter"
	"wiki_global/src/utils"

	"github.com/gin-gonic/gin"
)

// GetUserInfo 获取用户详情
func GetUserInfo(c *gin.Context) {

	var req oauth.GetUserInfoRequest
	req.CountryCode = c.Request.FormValue("countryCode")
	req.ApplicationType = convert.StrToInt(c.Request.FormValue("applicationType"))
	if req.CountryCode == "" || req.ApplicationType == 0 {
		i18nresponse.Error(c, "1010004")
		return
	}
	req.UserId = convert.Int64ToString(utils.GetUserId(c))
	result, err := usercenter.GetUserInfo(req)
	if err != nil || result.Succeed != true || result.Message != "success" {
		if result.Message == "" {
			i18nresponse.Error(c, "1010006") // 获取用户信息失败
			return
		} else {
			i18nresponse.Error(c, result.Message)
			return
		}
	}

	i18nresponse.Success(c, "ok", result.Result)
}

// ModifyPassByPhone 通过手机号找回密码
func ModifyPassByPhone(c *gin.Context) {

	var req oauth.ModifyPassByPhoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	//req.UserId = convert.Int64ToString(utils.GetUserId(c))
	result, err := usercenter.ModifyPassByPhone(req)
	if err != nil || result.Data.Succeed != true || result.Data.Message != "success" {
		if result.Data.Message == "" {
			i18nresponse.Error(c, "1010007") // 找回密码失败
			return
		} else {
			i18nresponse.Error(c, result.Data.Message)
			return
		}
	}

	i18nresponse.Success(c, "ok", result.Data)
}

// ModifyPassByOld 通过旧密码改新密码
func ModifyPassByOld(c *gin.Context) {

	var req oauth.ModifyPassByOldRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	req.UserId = convert.Int64ToString(utils.GetUserId(c))
	result, err := usercenter.ModifyPassByOld(req)
	if err != nil || result.Data.Succeed != true || result.Data.Message != "success" {
		if result.Data.Message == "" {
			i18nresponse.Error(c, "1010008") // 修改密码失败
			return
		} else {
			i18nresponse.Error(c, result.Data.Message)
			return
		}
	}

	i18nresponse.Success(c, "ok", result.Data)
}

// CheckMailbox 校验邮箱是否已经验证过
func CheckMailbox(c *gin.Context) {

	var req oauth.CheckMailboxRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	req.UserId = convert.Int64ToString(utils.GetUserId(c))
	result, err := usercenter.CheckMailbox(req)
	if err != nil {
		if result.Data.Message == "" {
			i18nresponse.Error(c, "1010009") // 邮箱验证失败
			return
		} else {
			i18nresponse.Error(c, result.Data.Message)
			return
		}
	}

	i18nresponse.Success(c, "ok", result.Data)
}

// SendEmailCode 校验邮箱是否已经验证过
func SendEmailCode(c *gin.Context) {

	var req oauth.SendEmailCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	req.UserId = convert.Int64ToString(utils.GetUserId(c))
	result, err := usercenter.SendEmailCode(req)
	if err != nil || result.Data.Succeed != true || result.Data.Message != "success" {
		if result.Data.Message == "" {
			i18nresponse.Error(c, "1010010") // 发送验证码失败
			return
		} else {
			i18nresponse.Error(c, result.Data.Message)
			return
		}
	}

	i18nresponse.Success(c, "ok", result.Data)
}

// ConfirmEmailByCode 验证邮箱（验证码）
func ConfirmEmailByCode(c *gin.Context) {

	var req oauth.ConfirmEmailByCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	req.UserId = convert.Int64ToString(utils.GetUserId(c))
	result, err := usercenter.ConfirmEmailByCode(req)
	if err != nil || result.Data.Succeed != true || result.Data.Message != "success" {
		if result.Data.Message == "" {
			i18nresponse.Error(c, "1010009") // 邮箱验证失败
			return
		} else {
			i18nresponse.Error(c, result.Data.Message)
			return
		}
	}

	i18nresponse.Success(c, "ok", result.Data)
}

// ConfirmEmailByLine 验证邮箱（链接）
func ConfirmEmailByLine(c *gin.Context) {

	var req oauth.ConfirmEmailByLineRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	req.UserId = convert.Int64ToString(utils.GetUserId(c))
	result, err := usercenter.ConfirmEmailByLine(req)
	if err != nil || result.Data.Succeed != true || result.Data.Message != "success" {
		if result.Data.Message == "" {
			i18nresponse.Error(c, "1010009") // 邮箱验证失败
			return
		} else {
			i18nresponse.Error(c, result.Data.Message)
			return
		}
	}

	i18nresponse.Success(c, "ok", result.Data)
}
