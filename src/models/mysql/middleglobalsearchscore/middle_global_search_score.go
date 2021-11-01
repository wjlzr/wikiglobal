package middleglobalsearchscore

import (
	"go.uber.org/zap"
	"time"
	orm "wiki_global/src/db/mysql"
	"wiki_global/src/utils/log"
)

type MiddleGlobalSearchScore struct {
	Code       string    `json:"code"`
	Region     string    `json:"region"`
	Score      float64   `json:"score"`
	CreateTime time.Time `json:"createtime"`
	UpdateTime time.Time `json:"updatetime"`
	ID         string    `json:"id"`
}

func (*MiddleGlobalSearchScore) TableName() string {
	return "middle_global_search_score"
}

// Find 获取多条数据
func (m *MiddleGlobalSearchScore) Find() (middleGlobalSearchScores []MiddleGlobalSearchScore, err error) {

	if err = orm.Engine.Table(m.TableName()).Cols("Code", "region", "score", "createtime", "updatetime").Where("Code = ?", m.Code).Find(&middleGlobalSearchScores); err != nil {
		log.Logger().Error("middleglobalsearchscore Find Find err: ", zap.Error(err))
		return nil, err
	}

	return middleGlobalSearchScores, nil
}
