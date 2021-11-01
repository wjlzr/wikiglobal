package attention

import (
	"github.com/k0kubun/pp"
	"wiki_global/src/common/convert"
	"wiki_global/src/config"
	orm "wiki_global/src/db/mysql"
	"wiki_global/src/models/es"
	"wiki_global/src/models/es/base"
	"wiki_global/src/utils"
	"wiki_global/src/utils/log"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type Attention struct {
	ID       int64  `json:"id"`
	UserId   string `json:"user_id"`
	Type     int    `json:"type"`
	Uuid     string `json:"uuid"`
	CreateAt int64  `json:"create_at" xorm:"created"`
	UpdateAt int64  `json:"update_at" xorm:"updated"`
}

type MyFollowRequest struct {
	UserId    string `json:"user_id"`
	Type      int    `json:"type"`
	PageIndex int    `json:"page_index"`
	PageSize  int    `json:"page_size"`
}

type CountFollowResponse struct {
	CompanyCount  int64 `json:"company_count"`
	OfficersCount int64 `json:"officers_count"`
}

// Follow 关注操作
func (receiver *Attention) Follow(c *gin.Context) (val bool, err error) {

	receiver.UserId = convert.Int64ToString(utils.GetUserId(c))
	receiver.Uuid = convert.IntToString(convert.StrToInt(receiver.Uuid) - config.Conf().Encryption.SecretKey)

	_, err = orm.Engine.Insert(receiver)
	if err != nil {
		log.Logger().Error("attention Follow insert err: ", zap.Error(err))
		return false, err
	}

	return true, nil
}

// CancelFollow 取消关注操作
func (receiver *Attention) CancelFollow(c *gin.Context) (val bool, err error) {

	pp.Println("uuid:", receiver.Uuid)
	pp.Println("解密后的uuid:", convert.IntToString(convert.StrToInt(receiver.Uuid)-config.Conf().Encryption.SecretKey))
	receiver.UserId = convert.Int64ToString(utils.GetUserId(c))
	receiver.Uuid = convert.IntToString(convert.StrToInt(receiver.Uuid) - config.Conf().Encryption.SecretKey)
	pp.Println(receiver)
	_, err = orm.Engine.Delete(receiver)
	if err != nil {
		log.Logger().Error("attention CancelFollow Delete err: ", zap.Error(err))
		return false, err
	}

	return val, nil
}

// MyFollow 我的关注列表
func (receiver *MyFollowRequest) MyFollow(c *gin.Context) (info interface{}, count int, err error) {

	var a []string

	receiver.UserId = convert.Int64ToString(utils.GetUserId(c))
	if err = orm.Engine.Table("attention").Cols("uuid").Where("user_id = ? and type = ?", receiver.UserId, receiver.Type).OrderBy("create_at desc").Find(&a); err != nil {
		log.Logger().Error("attention MyFollow Find err: ", zap.Error(err))
		return info, 0, err
	}

	if len(a) != 0 {

		condition := make(map[string][]string)
		condition["uuid"] = a
		// 获取组装好的条件
		filter, boolQuery := es.SearchListRequest{Screen: condition, PageIndex: receiver.PageIndex, PageSize: receiver.PageSize}.FilterConditionByBrief()

		info, count, err = base.FindCompanyOrOfficersByUuid(c, a, receiver.Type, filter, boolQuery)
	}

	if err != nil {
		return nil, 0, err
	}

	return
}

// QueryInfo 查询
func (receiver Attention) QueryInfo(c *gin.Context) (attention []Attention, count int64, err error) {

	if count, err = orm.Engine.Cols("id", "user_id", "type", "uuid", "create_at", "update_at").Where("user_id = ? and uuid = ? and type = ?", receiver.UserId, convert.IntToString(convert.StrToInt(receiver.Uuid)-config.Conf().Encryption.SecretKey), receiver.Type).FindAndCount(&attention); err != nil {
		log.Logger().Error("attention QueryInfo FindAndCount err: ", zap.Error(err))
		return attention, 0, err
	}

	return attention, count, nil
}

func (receiver Attention) CountByFollow() (countFollowResponse CountFollowResponse, err error) {

	if countFollowResponse.CompanyCount, err = orm.Engine.Where("user_id = ? and type = 1", receiver.UserId).Count(&receiver); err != nil {
		log.Logger().Error("attention CountByFollow Count1 err: ", zap.Error(err))
		return CountFollowResponse{}, err
	}
	if countFollowResponse.OfficersCount, err = orm.Engine.Where("user_id = ? and type = 2", receiver.UserId).Count(&receiver); err != nil {
		log.Logger().Error("attention CountByFollow Count2 err: ", zap.Error(err))
		return CountFollowResponse{}, err
	}

	return countFollowResponse, nil
}
