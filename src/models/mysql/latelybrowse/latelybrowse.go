package latelybrowse

import (
	orm "wiki_global/src/db/mysql"
	"wiki_global/src/models/es"
	"wiki_global/src/models/es/base"
	"wiki_global/src/utils/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LatelyBrowse struct {
	ID               int64  `json:"id"`
	UserId           string `json:"user_id"`
	Uuid             string `json:"uuid"`
	Name             string `json:"name"`
	CompanyNumber    string `json:"company_number"`
	Status           int64  `xorm:"default:1" json:"status"`
	Type             int64  `json:"type"`
	JurisdictionCode string `json:"jurisdiction_code"`
	IsSystem         int64  `xorm:"default:2" json:"is_system"`
	CreateAt         int64  `json:"create_at" xorm:"created"`
	UpdateAt         int64  `json:"update_at" xorm:"updated"`
}

type LatelyBrowseRequest struct {
	ID        int64  `json:"id"`
	UserId    string `json:"user_id"`
	Type      int    `json:"type"`
	IsSystem  int64  `json:"is_system"`
	PageIndex int    `json:"page_index"`
	PageSize  int    `json:"page_size"`
}

func (l *LatelyBrowseRequest) QueryList(c *gin.Context) (info interface{}, count int, err error) {

	var uuid []string
	if err = orm.Engine.Table("lately_browse").Cols("uuid").Where("type = ? and is_system = ? and status = ? and user_id = ?", l.Type, l.IsSystem, 1, l.UserId).OrderBy("create_at desc").Find(&uuid); err != nil {
		log.Logger().Error("latelybrowse QueryList Find err: ", zap.Error(err))
		return nil, 0, err
	}

	if len(uuid) != 0 {
		condition := make(map[string][]string)
		condition["uuid"] = uuid
		// 获取组装好的条件
		filter, boolQuery := es.SearchListRequest{Screen: condition, PageIndex: l.PageIndex, PageSize: l.PageSize}.FilterConditionByBrief()

		info, count, err = base.FindCompanyOrOfficersByUuid(c, uuid, l.Type, filter, boolQuery)
	}

	if err != nil {
		return nil, 0, err
	}

	return
}

// 插入
func (l *LatelyBrowse) Insert() (val bool, err error) {

	if val, err = l.QueryInfo(); err != nil || val == false {
		return false, nil
	}

	if _, err = orm.Engine.Insert(l); err != nil {
		log.Logger().Error("latelybrowse Insert Insert err: ", zap.Error(err))
		return false, err
	}

	return true, nil
}

// 获取单个详情
func (l *LatelyBrowse) QueryInfo() (val bool, err error) {
	var id []string
	if err = orm.Engine.Table("lately_browse").Cols("id").Where("uuid = ? and user_id = ? and type =?", l.Uuid, l.UserId, l.Type).Find(&id); err != nil {
		log.Logger().Error("latelybrowse QueryInfo Find err: ", zap.Error(err))
		return false, err
	}

	if len(id) == 0 {
		return true, nil
	}

	return false, nil
}
