package hotdata

import (
	"fmt"
	orm "wiki_global/src/db/mysql"
	"wiki_global/src/models/es"
	"wiki_global/src/models/es/base"
	"wiki_global/src/utils/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HotData struct {
	ID               int64  `json:"id" xorm:"id"`
	Uuid             string `json:"uuid"`
	Name             string `json:"name"`
	CompanyNumber    string `json:"company_number"`
	Type             int    `json:"type"`
	Status           int64  `json:"status"`
	AreaCode         string `json:"area_code"`
	JurisdictionCode string `json:"jurisdiction_code"`
	CreateAt         int64  `json:"create_at" xorm:"created"`
	UpdateAt         int64  `json:"update_at" xorm:"updated"`
}

type HotDataRequest struct {
	Type      int    `json:"type"`
	AreaCode  string `json:"area_code"`
	PageIndex int    `json:"page_index"`
	PageSize  int    `json:"page_size"`
}

func (h *HotDataRequest) QueryList(c *gin.Context) (info interface{}, count int, err error) {

	var uuid []string
	table := orm.Engine.Table("hot_data")
	table = table.Where("type = ?", h.Type)
	if h.AreaCode != "" {
		table = table.Where("area_code = ?", h.AreaCode)
	}

	if err = table.Cols("uuid").Find(&uuid); err != nil {
		log.Logger().Error("hotdata QueryList Find err: ", zap.Error(err))
		return nil, 0, err
	}

	if len(uuid) != 0 {
		condition := make(map[string][]string)
		condition["uuid"] = uuid
		// 获取组装好的条件
		filter, boolQuery := es.SearchListRequest{Screen: condition, PageIndex: h.PageIndex, PageSize: h.PageSize}.FilterConditionByBrief()

		info, count, err = base.FindCompanyOrOfficersByUuid(c, uuid, h.Type, filter, boolQuery)
		fmt.Printf("data：%+v \n", info)
	}

	if err != nil {
		return nil, 0, err
	}

	return
}
