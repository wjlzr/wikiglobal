package usercenter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/k0kubun/pp"
	"io"
	"io/ioutil"
	"net/http"
	"wiki_global/src/config"
	"wiki_global/src/models/mysql/oauth"
	"wiki_global/src/utils/log"

	"go.uber.org/zap"
)

//init token
func passiveInit() error {
	resp, err := http.Get(config.Conf().UserCenter.TestUrl + fmt.Sprintf(getToken+"?username=%s&password=%s", userName, password))
	if err != nil {
		log.Logger().Error("UserCenter init http err：", zap.Error(err))
		return err
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Logger().Error("UserCenter init ioutil.ReadAll err：", zap.Error(err))
		return err
	}
	var t tokenResponse
	_ = json.Unmarshal(bs, &t)
	if t.Status == false || t.AccessToken == "" {
		log.Logger().Info("UserCenter init 授权失败 err：", zap.Error(err))
		return errors.New("授权失败")
	}

	authorization = t.TokenType + " " + t.AccessToken
	return nil
}

// http请求
func request(method, url string, body io.Reader) (request *http.Request, err error) {

	// 签名流程
	if resp := passiveInit(); resp != nil {
		return nil, errors.New("签名流程失败")
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Logger().Error("UserCenter request http err：", zap.Error(err))
		return request, err
	}
	req.Header.Add("Authorization", authorization)
	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

//返回参数统一处理
func responseHandle(request *http.Request) []byte {
	client := &http.Client{}

	resp, _ := client.Do(request)
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("调用中台接口：%+v \n", resp.Request)
	fmt.Printf("用户中台返回值：%s \n", string(content))
	return content
}

// ValidatePhone 验证手机号是否注册过
func ValidatePhone(code, phone string) (success bool, err error) {

	request, err := request(http.MethodGet, config.Conf().UserCenter.User+fmt.Sprintf(registerValidateUserPhone+"?areaCode=%s&phoneNumber=%s", code, phone), nil)
	if err != nil {
		log.Logger().Error("UserCenter ValidatePhone 请求 err：", zap.Error(err))
		return false, err
	}

	content := responseHandle(request)
	var v ValidateUserPhoneResponse
	_ = json.Unmarshal(content, &v)
	if v.Code != 200 || v.Success != true {
		log.Logger().Info("UserCenter init 手机号验证失败 response：", zap.Any("response", v))
		return false, errors.New(v.Msg)
	}
	return v.Data.Succeed, nil
}

// SendCode 发送验证码
func SendCode(areaCode, phone, languageCode, userId string, smsBusinessType int) (s SendCodeResponse, err error) {

	jsonStr, _ := json.Marshal(sendCodeRequest{AreaCode: areaCode, Phone: phone, LanguageCode: languageCode, UserId: userId, SmsBusinessType: smsBusinessType, ApplicationType: applicationType})
	request, err := request(http.MethodPost, config.Conf().UserCenter.User+sendCode, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Logger().Error("UserCenter SendCode 请求 err：", zap.Error(err))
		return SendCodeResponse{}, err
	}

	content := responseHandle(request)
	_ = json.Unmarshal(content, &s)
	if s.Code != 200 || s.Success != true {
		log.Logger().Info("UserCenter SendCode 发送验证码Error response：", zap.Any("response", s))
		return SendCodeResponse{}, errors.New(s.Msg)
	}
	return s, nil
}

// Register 用户注册
func Register(r oauth.RegisterRequest) (result CurrencyWithUserResponseData, err error) {
	pp.Println("普通用户注册参数")
	pp.Println(r)
	jsonStr, _ := json.Marshal(r)
	request, err := request(http.MethodPost, config.Conf().UserCenter.User+register, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Logger().Error("UserCenter Register 请求 err：", zap.Error(err))
		return CurrencyWithUserResponseData{}, err
	}

	content := responseHandle(request)
	var res CurrencyWithUserResponse
	_ = json.Unmarshal(content, &res)
	if res.Code != 200 || res.Success != true {
		log.Logger().Info("UserCenter Register 用户注册Error response：", zap.Any("response", res))
		return CurrencyWithUserResponseData{}, errors.New(res.Msg)
	}
	return res.Data, nil
}

// ThirdRegister 第三方用户注册
func ThirdRegister(r oauth.ThirdRegisterRequest) (result CurrencyWithUserResponseData, err error) {

	pp.Println("第三方用户注册参数")
	pp.Println(r)

	jsonStr, _ := json.Marshal(r)
	request, err := request(http.MethodPost, config.Conf().UserCenter.User+thirdPartyRegister, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Logger().Error("UserCenter ThirdRegister 请求 err：", zap.Error(err))
		return CurrencyWithUserResponseData{}, err
	}

	content := responseHandle(request)
	var res CurrencyWithUserResponse
	_ = json.Unmarshal(content, &res)
	if res.Code != 200 || res.Success != true {
		log.Logger().Info("UserCenter ThirdRegister 用户第三方注册Error response：", zap.Any("response", res))
		return CurrencyWithUserResponseData{}, errors.New(res.Msg)
	}
	return res.Data, nil
}

// ValidateCode 验证短信验证码
func ValidateCode(req oauth.ValidateCodeRequest) (res CurrencyResponse, err error) {

	request, err := request(http.MethodGet, config.Conf().UserCenter.User+fmt.Sprintf(validateCode+"?areaCode=%s&phoneNumber=%s&smscode=%s&userId=%s&applicationType=%d", req.AreaCode, req.PhoneNumber, req.Smscode, req.UserId, applicationType), nil)
	if err != nil {
		log.Logger().Error("UserCenter ValidateCode 请求 err：", zap.Error(err))
		return CurrencyResponse{}, err
	}

	content := responseHandle(request)
	_ = json.Unmarshal(content, &res)
	return res, nil
}

// QuickLogin 快捷登录
func QuickLogin(q oauth.QuickLoginRequest) (result CurrencyWithUserResponseData, err error) {
	q.ApplicationType = applicationType
	jsonStr, _ := json.Marshal(q)
	request, err := request(http.MethodPost, config.Conf().UserCenter.User+quickLogin, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Logger().Error("UserCenter QuickLogin 请求 err：", zap.Error(err))
		return CurrencyWithUserResponseData{}, err
	}

	content := responseHandle(request)
	var resp CurrencyWithUserResponse
	_ = json.Unmarshal(content, &resp)
	if resp.Code != 200 || resp.Success != true {
		log.Logger().Info("UserCenter QuickLogin 快捷登录Error response：", zap.Any("response", resp))
		return CurrencyWithUserResponseData{}, errors.New(resp.Msg)
	}
	return resp.Data, nil
}

// Login 账号密码登录
func Login(l oauth.LoginRequest) (result CurrencyWithUserResponseData, err error) {
	l.ApplicationType = applicationType
	jsonStr, _ := json.Marshal(l)
	request, err := request(http.MethodPost, config.Conf().UserCenter.User+login, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Logger().Error("UserCenter Login 请求 err：", zap.Error(err))
		return CurrencyWithUserResponseData{}, err
	}

	content := responseHandle(request)
	var resp CurrencyWithUserResponse
	_ = json.Unmarshal(content, &resp)
	if resp.Code != 200 || resp.Success != true {
		log.Logger().Info("UserCenter Login 普通登录Error response：", zap.Any("response", resp))
		return CurrencyWithUserResponseData{}, errors.New(resp.Msg)
	}
	return resp.Data, nil
}

// GetUserInfo 获取用户详情
func GetUserInfo(req oauth.GetUserInfoRequest) (result CurrencyWithUserResponseData, err error) {
	request, err := request(http.MethodGet, config.Conf().UserCenter.User+fmt.Sprintf(getUserInfo+"?userId=%s&countryCode=%s&applicationType=%d", req.UserId, req.CountryCode, applicationType), nil)
	if err != nil {
		log.Logger().Error("UserCenter GetUserInfo 请求 err：", zap.Error(err))
		return CurrencyWithUserResponseData{}, err
	}

	content := responseHandle(request)
	var resp CurrencyWithUserResponse
	_ = json.Unmarshal(content, &resp)
	if resp.Code != 200 || resp.Success != true {
		log.Logger().Info("UserCenter GetUserInfo 获取用户详情Error response：", zap.Any("response", resp))
		return CurrencyWithUserResponseData{}, errors.New(resp.Msg)
	}
	return resp.Data, nil
}

// ModifyPassByPhone 通过手机号找回密码
func ModifyPassByPhone(m oauth.ModifyPassByPhoneRequest) (result CurrencyResponse, err error) {
	jsonStr, _ := json.Marshal(m)
	request, err := request(http.MethodPost, config.Conf().UserCenter.User+modifyPassByPhone, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Logger().Error("UserCenter ModifyPassByPhone 请求 err：", zap.Error(err))
		return CurrencyResponse{}, err
	}

	content := responseHandle(request)
	var resp CurrencyResponse
	_ = json.Unmarshal(content, &resp)
	if resp.Code != 200 || resp.Success != true {
		log.Logger().Info("UserCenter ModifyPassByPhone 通过手机号找回密码Error response：", zap.Any("response", resp))
		return CurrencyResponse{}, errors.New(resp.Msg)
	}
	return resp, nil
}

// ModifyPassByOld 通过旧密码改新密码
func ModifyPassByOld(m oauth.ModifyPassByOldRequest) (result CurrencyResponse, err error) {
	jsonStr, _ := json.Marshal(m)
	request, err := request(http.MethodPost, config.Conf().UserCenter.User+modifyPassByOld, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Logger().Error("UserCenter ModifyPassByOld 请求 err：", zap.Error(err))
		return CurrencyResponse{}, err
	}

	content := responseHandle(request)
	var resp CurrencyResponse
	_ = json.Unmarshal(content, &resp)
	if resp.Code != 200 || resp.Success != true {
		log.Logger().Info("UserCenter ModifyPassByOld 通过旧密码改新密码Err response：", zap.Any("response", resp))
		return CurrencyResponse{}, errors.New(resp.Msg)
	}
	return resp, nil
}

// CheckMailbox 校验邮箱是否验证
func CheckMailbox(req oauth.CheckMailboxRequest) (resp ValidateUserPhoneResponse, err error) {
	request, err := request(http.MethodGet, config.Conf().UserCenter.User+fmt.Sprintf(checkMailbox+"?userId=%s&email=%s&applicationType=%d", req.UserId, req.Email, applicationType), nil)
	if err != nil {
		log.Logger().Error("UserCenter CheckMailbox 请求 err：", zap.Error(err))
		return ValidateUserPhoneResponse{}, err
	}

	content := responseHandle(request)
	_ = json.Unmarshal(content, &resp)
	if resp.Code != 200 || resp.Success != true {
		log.Logger().Info("UserCenter CheckMailbox 校验邮箱Error response：", zap.Any("response", resp))
		return ValidateUserPhoneResponse{}, errors.New(resp.Msg)
	}
	return resp, nil
}

// SendEmailCode 发送邮箱验证码
func SendEmailCode(m oauth.SendEmailCodeRequest) (result CurrencyResponse, err error) {
	m.ApplicationType = applicationType
	jsonStr, _ := json.Marshal(m)
	request, err := request(http.MethodPost, config.Conf().UserCenter.User+sendEmail, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Logger().Error("UserCenter SendEmailCode 请求 err：", zap.Error(err))
		return CurrencyResponse{}, err
	}

	content := responseHandle(request)
	var resp CurrencyResponse
	_ = json.Unmarshal(content, &resp)
	if resp.Code != 200 || resp.Success != true {
		log.Logger().Info("UserCenter SendEmailCode 发送邮箱验证码Err response：", zap.Any("response", resp))
		return CurrencyResponse{}, errors.New(resp.Msg)
	}
	return resp, nil
}

// ConfirmEmailByCode 验证邮箱（验证码）
func ConfirmEmailByCode(m oauth.ConfirmEmailByCodeRequest) (result CurrencyResponse, err error) {
	m.ApplicationType = applicationType
	jsonStr, _ := json.Marshal(m)
	request, err := request(http.MethodPost, config.Conf().UserCenter.User+confirmEmailByCode, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Logger().Error("UserCenter ConfirmEmailByCode 请求 err：", zap.Error(err))
		return CurrencyResponse{}, err
	}

	content := responseHandle(request)
	var resp CurrencyResponse
	_ = json.Unmarshal(content, &resp)
	if resp.Code != 200 || resp.Success != true {
		log.Logger().Info("UserCenter ConfirmEmailByCode 验证邮箱（验证码）Err response：", zap.Any("response", resp))
		return CurrencyResponse{}, errors.New(resp.Msg)
	}
	return resp, nil
}

// ConfirmEmailByLine 验证邮箱（链接）
func ConfirmEmailByLine(m oauth.ConfirmEmailByLineRequest) (result CurrencyResponse, err error) {
	m.ApplicationType = applicationType
	jsonStr, _ := json.Marshal(m)
	request, err := request(http.MethodPost, config.Conf().UserCenter.User+confirmEmailByline, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Logger().Error("UserCenter ConfirmEmailByLine 请求 err：", zap.Error(err))
		return CurrencyResponse{}, err
	}

	content := responseHandle(request)
	var resp CurrencyResponse
	_ = json.Unmarshal(content, &resp)
	if resp.Code != 200 || resp.Success != true {
		log.Logger().Info("UserCenter ConfirmEmailByLine 验证邮箱（链接）Err response：", zap.Any("response", resp))
		return CurrencyResponse{}, errors.New(resp.Msg)
	}
	return resp, nil
}

// ThirdPartyLogin 第三方登录
func ThirdPartyLogin(m oauth.ThirdPartyLoginRequest) (result CurrencyWithUserResponseData, err error) {
	m.ApplicationType = applicationType
	jsonStr, _ := json.Marshal(m)
	request, err := request(http.MethodPost, config.Conf().UserCenter.User+thirdPartyLogin, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Logger().Error("UserCenter ThirdPartyLogin 请求 err：", zap.Error(err))
		return CurrencyWithUserResponseData{}, err
	}

	content := responseHandle(request)
	var resp CurrencyWithUserResponse
	_ = json.Unmarshal(content, &resp)
	if resp.Code != 200 || resp.Success != true {
		log.Logger().Info("UserCenter ThirdPartyLogin 第三方登录 Err response：", zap.Any("response", resp))
		return CurrencyWithUserResponseData{}, errors.New(resp.Msg)
	}
	return resp.Data, nil
}

// ValidateRegisterPhone 验证手机号是否注册或者第三方是否绑定这个手机号
func ValidateRegisterPhone(m oauth.ValidateRegisterPhoneRequest) (result ValidateRegisterPhoneResponse, err error) {
	m.ApplicationType = applicationType
	jsonStr, _ := json.Marshal(m)
	request, err := request(http.MethodPost, config.Conf().UserCenter.User+validateRegisterPhone, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Logger().Error("UserCenter ThirdPartyLogin 请求 err：", zap.Error(err))
		return ValidateRegisterPhoneResponse{}, err
	}

	content := responseHandle(request)
	_ = json.Unmarshal(content, &result)
	if result.Code != 200 || result.Success != true {
		log.Logger().Info("UserCenter ThirdPartyLogin 第三方登录 Err response：", zap.Any("response", result))
		return ValidateRegisterPhoneResponse{}, errors.New(result.Msg)
	}
	return result, nil
}
