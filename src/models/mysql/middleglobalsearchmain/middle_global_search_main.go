package middleglobalsearchmain

import (
	"go.uber.org/zap"
	orm "wiki_global/src/db/mysql"
	"wiki_global/src/utils/log"
)

type MiddleGlobalSearchMain struct {
	Code                string `xorm:"Code" json:"Code"`
	Platform            int64  `xorm:"Platform" json:"Platform"`
	NoSearchStatus      int64  `xorm:"BIT(1) NoSearchStatus" json:"NoSearchStatus"`
	NoSearchClosingDate string `xorm:"NoSearchClosingDate" json:"NoSearchClosingDate"`
	Type                int64  `xorm:"Type" json:"Type"`
	RegisterCountry     string `xorm:"RegisterCountry" json:"RegisterCountry"`
	CreateTime          string `xorm:"CreateTime" json:"CreateTime"`
	UpdateTime          string `xorm:"UpdateTime" json:"UpdateTime"`
	ID                  string `json:"ID"`
}

func (*MiddleGlobalSearchMain) TableName() string {
	return "middle_global_search_main"
}

// Find 获取多条数据
func (m *MiddleGlobalSearchMain) Find() (middleGlobalSearchMains []MiddleGlobalSearchMain, err error) { //2449572669 6511612145

	if err = orm.Engine.Table(m.TableName()).Cols("Code", "Platform", "NoSearchStatus", "NoSearchClosingDate", "Type", "CreateTime", "UpdateTime", "RegisterCountry").OrderBy("Code desc").Limit(25000, 0).Find(&middleGlobalSearchMains); err != nil {
		log.Logger().Error("middleglobalsearchmain Find Find err: ", zap.Error(err))
		return middleGlobalSearchMains, err
	}

	return middleGlobalSearchMains, nil
}

// FindOne 根据条件查询单条数据
func (m *MiddleGlobalSearchMain) FindOne() (*MiddleGlobalSearchMain, error) { //2449572669 6511612145

	val, err := orm.Engine.Table(m.TableName()).Cols("Code", "Platform", "NoSearchStatus", "NoSearchClosingDate", "Type", "CreateTime", "UpdateTime", "RegisterCountry").Get(m)
	if !val || err != nil {
		log.Logger().Error("middleglobalsearchmain FindOneByCode Get err: ", zap.Error(err))
		return nil, err
	}

	return m, nil
}

// Count 总数
func (m MiddleGlobalSearchMain) Count() (int64, error) {

	total, err := orm.Engine.Table(m.TableName()).Cols("Code", "Platform", "NoSearchStatus", "NoSearchClosingDate", "Type", "CreateTime", "UpdateTime", "RegisterCountry").Count(&m)
	if err != nil {
		log.Logger().Error("middleglobalsearchmain Count Count err: ", zap.Error(err))
		return 0, err
	}

	return total, nil
}
