package middleglobalsearchname

import (
	"go.uber.org/zap"
	orm "wiki_global/src/db/mysql"
	"wiki_global/src/utils/log"
)

type MiddleGlobalSearchName struct {
	Code         string `xorm:"Code" json:"Code"`
	LanguageCode string `xorm:"LanguageCode" json:"LanguageCode"`
	Shortname    string `xorm:"ShortName" json:"ShortName"`
	FullName     string `xorm:"FullName" json:"FullName"`
	CreateTime   string `xorm:"CreateTime" json:"CreateTime"`
	UpdateTime   string `xorm:"UpdateTime" json:"UpdateTime"`
	ID           string `json:"ID"`
}

func (*MiddleGlobalSearchName) TableName() string {
	return "middle_global_search_name"
}

// Find 获取多条数据
func (m *MiddleGlobalSearchName) Find() (middleGlobalSearchNames []MiddleGlobalSearchName, err error) {

	if err = orm.Engine.Table(m.TableName()).Cols("Code", "LanguageCode", "ShortName", "FullName", "CreateTime", "UpdateTime").Where("Code = ?", m.Code).Find(&middleGlobalSearchNames); err != nil {
		log.Logger().Error("middleglobalsearchname Find Find err: ", zap.Error(err))
		return nil, err
	}

	return middleGlobalSearchNames, nil
}
