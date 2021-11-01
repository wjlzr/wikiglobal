package router

import (
	"time"
	"wiki_global/src/config"
	"wiki_global/src/controller/attention"
	"wiki_global/src/controller/download"
	"wiki_global/src/controller/geo"
	"wiki_global/src/controller/hotdata"
	"wiki_global/src/controller/latelybrowse"
	"wiki_global/src/controller/oauth"
	"wiki_global/src/controller/region"
	"wiki_global/src/controller/search/company"
	"wiki_global/src/controller/search/officers"
	"wiki_global/src/controller/share"
	"wiki_global/src/controller/user"
	"wiki_global/src/controller/version"
	"wiki_global/src/middleware"
	"wiki_global/src/utils/token"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//路由配置
func RouterEngine(token *token.Token, zapLogger *zap.Logger) *gin.Engine {
	engine := gin.New()

	//gin中间件加载
	engine.Use(middleware.Cors())
	engine.Use(middleware.Secure())
	engine.Use(middleware.Language())
	engine.Use(middleware.CustomIntercept())
	engine.Use(middleware.Ginzap(zapLogger, time.RFC3339, true))
	engine.Use(middleware.RecoveryWithZap(zapLogger, true))

	engine.Use(middleware.UserAuthMiddleware(token, middleware.AllowPathPrefixSkipper(config.Conf().API.AllowPathPrefixSkipper)))

	apiGroup := engine.Group("api/v1")

	//oauth
	oauthGroup(apiGroup.Group("oauth"))
	//用户模块
	userGroup(apiGroup.Group("user"))
	//搜索模块
	searchGroup(apiGroup.Group("search"))
	//搜索模块
	configureGroup(apiGroup.Group("configure"))
	// 通用模块
	// 关注
	apiGroup.POST("/attention/follow", attention.Follow)
	apiGroup.POST("/attention/cancelFollow", attention.CancelFollow)
	apiGroup.POST("/attention/myFollow", attention.MyFollow)
	apiGroup.GET("/attention/countByFollow", attention.CountByFollow)
	// geo
	apiGroup.GET("/geo/getWithIpToLocation", geo.GetWithIpToLocation)
	// download
	apiGroup.GET("/download/getDownload", download.GetDownload)
	// version
	apiGroup.GET("/version/getVersion", version.Version)

	return engine
}

// 认证路由组
func oauthGroup(rg *gin.RouterGroup) {
	//登录
	rg.POST("/login", oauth.Login)
	rg.GET("/smsSend", oauth.SmsSend)
	rg.POST("/register", oauth.Register)
	rg.GET("/validateCode", oauth.ValidateCode)
	rg.POST("/quickLogin", oauth.QuickLogin)
	// 微信登录
	rg.POST("/thirdPartyLogin", oauth.ThirdPartyLogin)
	rg.GET("/sendCode", oauth.SendCode)
	rg.GET("/getWeChatOpenId", oauth.GetWeChatOpenId)
	rg.POST("/validateRegisterPhone", oauth.ValidateRegisterPhone)
	rg.POST("/thirdRegister", oauth.ThirdRegister)
	rg.GET("/getQqOpenId", oauth.GetQqOpenId)
	// apple
	rg.POST("/appleVerify", oauth.AppleVerify)
}

// 用户路由组
func userGroup(rg *gin.RouterGroup) {
	// 用户列表
	rg.GET("/info", user.GetUserInfo)
	// 通过手机号找回密码
	rg.POST("/retrievePasswordByPhone", user.ModifyPassByPhone)
	// 通过旧密码改新密码
	rg.POST("/modifyPassByOld", user.ModifyPassByOld)
	// 校验邮箱是否已经验证过
	rg.POST("/checkMailbox", user.CheckMailbox)
	// 发送邮箱验证码
	rg.POST("/sendEmailCode", user.SendEmailCode)
	// 验证邮箱（验证码）
	rg.POST("/confirmEmailByCode", user.ConfirmEmailByCode)
	// 验证邮箱（链接）
	rg.POST("/ConfirmEmailByLine", user.ConfirmEmailByLine)
}

// 搜索路由组
func searchGroup(rg *gin.RouterGroup) {
	// 公司列表
	rg.POST("/company/list", company.GetList)
	// 公司详情
	rg.GET("/company/info", company.GetInfo)
	// 公司详情-wikibit
	rg.GET("/company/evaluationInfo", company.GetEvaluationInfo)
	// wikibit详情-集团关系
	rg.GET("/company/groupRelationShip", company.GetGroupRelationShip)
	// 公司详情-wikifx
	rg.GET("/company/traderInfo", company.GetTraderInfo)
	// 公司详情-wikifx-点差
	rg.GET("/company/spreadCode", company.GetSpreadCode)
	// 公司详情-wikifx-点差-详情
	rg.GET("/company/wf/spreadHighLow", company.GetSpreadHighLowByCode)
	// 公司详情-wikifx-交易环境
	rg.GET("/company/wf/tradingEnvironment", company.TradingEnvironment)
	// 公司详情-wikifx-交易环境-历史评分
	rg.GET("/company/wf/historyScores", company.HistoryScores)
	// 公司详情-wikifx-实勘列表
	rg.GET("/company/wf/surveyList", company.SurveyList)
	// 公司详情-wikifx-实勘详情
	rg.GET("/company/wf/survey", company.Survey)
	// 公司详情-wikifx-旗舰店广告
	rg.GET("/company/wf/ultimateMember", company.UltimateMember)
	// 公司详情-wikifx-是否显示红色按钮处功能
	rg.GET("/company/wf/researchIsShow", company.ResearchIsShow)
	// 公司详情-wikifx-经纪商是否有历史数据（详情页是否显示蓝色按钮）
	rg.GET("/company/wf/brokerHistoryDataExist", company.BrokerHistoryDataExist)
	// 公司详情-wikifx-epc
	rg.GET("/company/wf/epc", company.Epc)
	// 公司详情-wikifx-提示语
	rg.GET("/company/wf/tips", company.Tips)
	// 公司详情-wikifx-提示语-统计
	rg.POST("/company/wf/traderCount", company.TraderCount)
	// 公司详情-wikifx-账户类型
	rg.GET("/company/wf/traderAccount", company.TraderAccount)
	// 公司详情-wikifx-MT4/5
	rg.GET("/company/wf/mt4", company.MT4)
	// 公司详情-wikifx-查找冒充交易商
	rg.GET("/company/wf/fake", company.Fake)
	// 公司详情-wikifx-研究院数据（饼图数据）
	rg.GET("/company/wf/researchInfo", company.ResearchInfo)
	// 公司详情-wikifx-经纪商历史走势
	rg.POST("/company/wf/brokerHistoryData", company.BrokerHistoryData)
	// 公司详情-wikifx-经纪商历史走势
	rg.GET("/company/wf/exposureList", company.ExposureList)
	// 高管列表
	rg.POST("/officers/list", officers.GetList)
	// 高管详情
	rg.GET("/officers/info", officers.GetInfo)
	// 地区
	rg.GET("/region/info", region.GetInfo)
	// 热门公司或者高管
	rg.POST("/hotdata/list", hotdata.GetList)
	// 获取最近浏览的公司/高管
	rg.POST("/latelybrowse/list", latelybrowse.GetList)
	// 添加浏览的公司/高管
	rg.POST("/latelybrowse/add", latelybrowse.Add)
	// 公司完成推荐
	rg.GET("/company/completionSuggest", company.CompletionSuggest)
	// 高管完成推荐
	rg.GET("/officers/completionSuggest", officers.CompletionSuggest)
	// 修改补充信息-公司
	//rg.POST("/company/update/:uuid", company.ManualUpdate)
	// 修改补充信息-高管
	//rg.POST("/officers/update/:uuid", officers.ManualUpdate)
	// 获取筛选条件聚合
	rg.POST("/company/getCondition", company.GetCondition)
}

// 搜索路由组
func configureGroup(rg *gin.RouterGroup) {
	// 公司列表
	rg.GET("/share/getUrl", share.GetUrl)
}
