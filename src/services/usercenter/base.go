package usercenter

import "github.com/gin-gonic/gin"

const (
	userName                  = "gsw"
	password                  = "2E6CAA66096F1D4BED0FE21EE5468FE2"
	getToken                  = "api/Permission/Login"                                         //获取token
	sendCode                  = "PersonCenter/usercenter/wikiglobal/sendcode"                  //发送验证码
	register                  = "PersonCenter/usercenter/wikiglobal/register"                  //用户账号注册
	login                     = "PersonCenter/usercenter/wikiglobal/login"                     //账号密码登录
	getUserInfo               = "PersonCenter/usercenter/wikiglobal/getuser"                   //获取用户信息
	quickLogin                = "PersonCenter/usercenter/wikiglobal/quicklogin"                //快捷登录
	validateCode              = "PersonCenter/usercenter/wikiglobal/validatecode"              //验证验证码
	registerValidateUserPhone = "PersonCenter/usercenter/wikiglobal/registervalidateuserphone" //验证手机号是否注册过
	modifyPassByPhone         = "PersonCenter/usercenter/wikiglobal/modifypassbyphone"         //通过手机号找回密码
	modifyPassByOld           = "PersonCenter/usercenter/wikiglobal/modifypassbyold"           //通过旧密码改新密码
	checkMailbox              = "PersonCenter/usercenter/wikiglobal/verificationemailpass"     //校验验证码是否验证
	sendEmail                 = "PersonCenter/usercenter/wikiglobal/sendemail"                 //发送邮箱验证码
	confirmEmailByCode        = "PersonCenter/usercenter/wikiglobal/comfirmemailbycode"        //验证邮箱（验证码）
	confirmEmailByline        = "PersonCenter/usercenter/wikiglobal/comfirmemailbyline"        //验证邮箱（链接）
	thirdPartyLogin           = "PersonCenter/usercenter/wikiglobal/thirdpartylogin"           //第三方登录
	validateRegisterPhone     = "PersonCenter/usercenter/wikiglobal/validateregisterphone"     //第三方登录 验证手机号是否注册或者第三方是否绑定这个手机号
	thirdPartyRegister        = "PersonCenter/usercenter/wikiglobal/thirdpartyregister"        //第三方注册
)

var (
	authorization   string
	c               *gin.Context
	applicationType = 61
)

type tokenRequest struct {
	UserName string
	Password string
}

// tokenRequest
type tokenResponse struct {
	Status      bool   `json:"status"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

// ValidateUserPhoneResponse 验证手机号response new
type ValidateUserPhoneResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"Success"`
	Data    struct {
		Result  string `json:"result"`
		Succeed bool   `json:"succeed"`
		Message string `json:"message"`
	} `json:"Data"`
}

// 发送验证码request
type sendCodeRequest struct {
	AreaCode        string `json:"areaCode"`
	Phone           string `json:"phone"`
	LanguageCode    string `json:"languageCode"`
	UserId          string `json:"userId"`
	SmsBusinessType int    `json:"smsBusinessType"`
	ApplicationType int    `json:"applicationType"`
}

// SendCodeResponse 发送验证码response
type SendCodeResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"Success"`
	Data    struct {
		Result struct {
			Requestid string `json:"requestid"`
		} `json:"Result"`
		Succeed bool   `json:"succeed"`
		Message string `json:"message"`
	} `json:"Data"`
}

// CurrencyWithUserResponse 带有用户信息通用返回的参数
type CurrencyWithUserResponse struct {
	Code    int                          `json:"code"`
	Success bool                         `json:"Success"`
	Msg     string                       `json:"msg"`
	Data    CurrencyWithUserResponseData `json:"data"`
}

type CurrencyWithUserResponseData struct {
	Succeed bool                           `json:"succeed"`
	Message string                         `json:"message"`
	Result  CurrencyWithUserResponseResult `json:"result"`
}

type CurrencyWithUserResponseResult struct {
	UserId               string `json:"userId"`
	Nickname             string `json:"nickname"`
	Nick                 string `json:"nick"`
	Avatar               string `json:"avatar"`
	Sex                  int    `json:"sex"`
	Areaflag             string `json:"areaflag"`
	Areacode             string `json:"areacode"`
	Phone                string `json:"phone"`
	Password             string `json:"password"`
	Email                string `json:"email"`
	Shoppingaddresscount int    `json:"shoppingaddresscount"`
	Realname             string `json:"realname"`
	Isphonecomfirm       bool   `json:"isphonecomfirm"`
	Isemailcomfirm       bool   `json:"isemailcomfirm"`
}

// ValidateCodeResponseOld 验证短信验证码Response old
type ValidateCodeResponseOld struct {
	RequestId string `json:"RequestId"`
	Timestamp string `json:"Timestamp"`
	Content   struct {
		Result struct {
			Succeed bool   `json:"succeed"`
			Message string `json:"message"`
		} `json:"result"`
	} `json:"Content"`
}

// 验证短信验证码Response 用通用CurrencyResponse
//type ValidateCodeResponse struct {
//	Code    int    `json:"code"`
//	Msg     string `json:"msg"`
//	Success bool   `json:"Success"`
//	Data    struct {
//		Succeed bool   `json:"succeed"`
//		Message string `json:"message"`
//	} `json:"Data"`
//}

// CurrencyResponse 通用response new
type CurrencyResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"Success"`
	Data    struct {
		Succeed bool   `json:"succeed"`
		Message string `json:"message"`
	} `json:"Data"`
}

// ValidateRegisterPhoneResponse 验证手机号是否注册或者第三方是否绑定这个手机号response
type ValidateRegisterPhoneResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"Success"`
	Data    struct {
		Result struct {
			ErrorType string `json:"errorType"`
		} `json:"Result"`
		Succeed bool   `json:"succeed"`
		Message string `json:"message"`
	} `json:"Data"`
}

// EvaluationInfoResponse 受评方详情response
type EvaluationInfoResponse struct {
	Code    int                        `json:"code"`
	Msg     string                     `json:"msg"`
	Success bool                       `json:"Success"`
	Data    EvaluationInfoDataResponse `json:"Data"`
}

type EvaluationInfoDataResponse struct {
	Succeed bool        `json:"succeed"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// SpreadCodResponse 交易商点差
type SpreadCodResponse struct {
	Code    int                        `json:"code"`
	Msg     string                     `json:"msg"`
	Success bool                       `json:"Success"`
	Data    EvaluationInfoDataResponse `json:"Data"`
}
