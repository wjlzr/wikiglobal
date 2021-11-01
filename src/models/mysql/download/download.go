package download

import (
	"go.uber.org/zap"
	orm "wiki_global/src/db/mysql"
	"wiki_global/src/utils/log"
)

type Download struct {
	ID   int64  `json:"id"`
	Url  string `json:"url"`
	Type int    `json:"type"`
}

// QueryInfo 查询
func (receiver *Download) QueryInfo() (download []Download, err error) {

	if err = orm.Engine.Cols("id", "url", "type").Find(&download); err != nil {
		log.Logger().Error("download QueryInfo FindAndCount err: ", zap.Error(err))
		return nil, err
	}

	return download, nil
}
