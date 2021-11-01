package base

import (
	"context"
	"wiki_global/src/models/es"
	"wiki_global/src/models/es/escompany"
	"wiki_global/src/models/es/esofficers"
	"wiki_global/src/utils/log"

	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

// FindCompanyOrOfficersByUuid 根据uuid查询公司或者高管数据
func FindCompanyOrOfficersByUuid(c context.Context, uuid []string, form int, filter *es.Filter, boolQuery *elastic.BoolQuery) (info interface{}, count int, err error) {

	condition := make(map[string][]string)
	condition["uuid"] = uuid
	// 获取组装好的条件

	if form == 1 {
		info, count, err = escompany.ListQueryByBrief(c, filter, boolQuery)
	} else {
		info, count, err = esofficers.ListQueryByBrief(c, filter, boolQuery)
	}

	if err != nil {
		log.Logger().Error("es FindCompanyOrOfficersByUuid 根据uuid查询公司或者高管 err：", zap.Error(err))
		return nil, 0, err
	}
	return
}
