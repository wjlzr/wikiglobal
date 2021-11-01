package esofficers

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"sync"
	"wiki_global/src/common/convert"
	"wiki_global/src/config"
	"wiki_global/src/db/es"
	es2 "wiki_global/src/models/es"
	"wiki_global/src/models/mongo/mgofficers"
	"wiki_global/src/utils/log"

	"github.com/gin-gonic/gin"

	"github.com/k0kubun/pp"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

const indices = "officers_v1"

// 高管列表搜索
func ListQuery(ctx context.Context, filter *es2.Filter, boolQuery *elastic.BoolQuery) (officers []*Officers, count int, err error) {

	resp, err := es.Client().Search().Index(indices).Query(boolQuery).Highlight(filter.Highlight).SortBy(filter.Sorters...).From(filter.From).Size(filter.Size).TrackTotalHits(true).Do(ctx)
	if err != nil {
		return nil, 0, err
	}

	if resp.TotalHits() == 0 {
		return nil, 0, nil
	}

	for _, e := range resp.Hits.Hits {
		var officer *Officers
		_ = json.Unmarshal(e.Source, &officer)
		// uuid加密
		officer.Uuid = convert.IntToString(convert.StrToInt(officer.Uuid) + config.Conf().Encryption.SecretKey)
		officer.RealName = officer.Name
		officer.RealCompanyNumber = officer.CompanyNumber
		for k, h := range e.Highlight {
			if k == "name" {
				officer.Name = h[0]
			} else if k == "company_number" {
				officer.CompanyNumber = h[0]
			}
		}
		officers = append(officers, officer)
	}
	return officers, int(resp.TotalHits()), nil
}

// 详情
func DetailQuery(ctx context.Context, uuid string) (officers Officers, err error) {

	pp.Println("uuid:", uuid)
	pp.Println("解密后的uuid:", convert.IntToString(convert.StrToInt(uuid)-config.Conf().Encryption.SecretKey))

	resp, err := es.Client().Search().Index(indices).Query(elastic.NewTermQuery("uuid", convert.IntToString(convert.StrToInt(uuid)-config.Conf().Encryption.SecretKey))).Pretty(true).Do(ctx)
	if err != nil {
		return Officers{}, err
	}

	if resp.TotalHits() == 0 {
		return Officers{}, errors.New("不存在此高管")
	}

	for _, hits := range resp.Hits.Hits {
		_ = json.Unmarshal([]byte(hits.Source), &officers)
	}

	officers.Uuid = convert.IntToString(convert.StrToInt(officers.Uuid) + config.Conf().Encryption.SecretKey)

	if officers.Company.Uuid != "" {
		officers.Company.Uuid = convert.IntToString(convert.StrToInt(officers.Company.Uuid) + config.Conf().Encryption.SecretKey)
	}
	return
}

// 高管筛选简约版
func ListQueryByBrief(ctx context.Context, filter *es2.Filter, boolQuery *elastic.BoolQuery) (officers []*Officers, count int, err error) {

	resp, err := es.Client().Search().Index(indices).Query(boolQuery).From(filter.From).Size(filter.Size).TrackTotalHits(true).Pretty(true).Do(ctx)
	if err != nil {
		return nil, 0, err
	}

	if resp.TotalHits() == 0 {
		return nil, 0, nil
	}

	for _, e := range resp.Each(reflect.TypeOf(&Officers{})) {
		c := e.(*Officers)
		c.Uuid = convert.IntToString(convert.StrToInt(c.Uuid) + config.Conf().Encryption.SecretKey)
		officers = append(officers, c)
	}
	return officers, int(resp.TotalHits()), nil
}

// 更新操作
func Update(ctx context.Context, uuid string, doc map[string]interface{}) (val bool, err error) {

	var officers Officers
	if officers, err = getInfoByUuid(ctx, uuid); err != nil {
		return false, err
	}

	if _, err = es.Client().Update().Index(indices).Id(officers.EsId).Doc(doc).Do(ctx); err != nil {
		return false, err
	}

	return true, nil
}

func getInfoByUuid(ctx context.Context, uuid string) (officers Officers, err error) {

	resp, err := es.Client().Search().Index(indices).Query(elastic.NewTermQuery("uuid", convert.IntToString(convert.StrToInt(uuid)-config.Conf().Encryption.SecretKey))).Pretty(true).Do(ctx)
	if err != nil {
		return Officers{}, err
	}

	if resp.TotalHits() == 0 {
		return Officers{}, errors.New("不存在此高管")
	}

	for _, hits := range resp.Hits.Hits {
		_ = json.Unmarshal([]byte(hits.Source), &officers)
	}
	officers.EsId = resp.Hits.Hits[0].Id
	return
}

func Insert(officers []Officers, num int, wg *sync.WaitGroup, skip int64) error {
	defer wg.Done()
	if len(officers) == 0 {
		pp.Println("mongo数据为空直接返回")
		return nil
	}
	_, _ = pp.Println("---------------第" + convert.IntToString(num*2) + "万条开始导入---------------")
	bulkRequest := es.Client().Bulk()
	for _, officer := range officers {
		doc := elastic.NewBulkIndexRequest().Index(indices).Doc(officer)
		bulkRequest = bulkRequest.Add(doc)
	}

	_, err := bulkRequest.Do(context.Background())
	if err != nil {
		log.Logger().Info("Insert officers err: ", zap.Error(err))
		//log.Logger().Info("company: ", zap.Any("detail", companys))
		var uuids []string
		for _, v := range officers {
			uuids = append(uuids, v.Uuid)
		}
		b, err := json.Marshal(uuids)
		// 失败将uuid保存起来
		name := mgofficers.CollectionName(skip)
		if err = ioutil.WriteFile("o-json/"+name+"-"+convert.IntToString(num*1)+".json", b, os.ModeAppend); err != nil {
			pp.Println(name+"-"+convert.IntToString(num*1)+" 高管生成json文件失败: ", zap.Error(err))
			log.Logger().Info("高管数据生成json文件失败: ", zap.Error(err))
		}
	}
	return nil
}

func Insert2(officers Officers) error {
	_, err := es.Client().Index().Index(indices).BodyJson(officers).Do(context.Background())
	if err != nil {
		_, _ = pp.Println("Insert officers err uuid: ", officers.Uuid, err.Error())
		log.Logger().Info("Insert officers uuid: "+officers.Uuid, zap.String("err: ", err.Error()))
	}
	return nil
}

// @summary 完成推荐
// @param ctx
// @param prefix
// @return officers
// @return err
func CompletionSuggestQuery(ctx *gin.Context, prefix string) (officers []Officers, err error) {

	suggest := elastic.NewCompletionSuggester("officers-suggest").Prefix(prefix).Field("name.relevant")

	resp, err := es.Client().Search().Index(indices).Suggester(suggest).Do(ctx)
	if err != nil {
		return nil, err
	}

	var officer Officers
	for _, v := range resp.Suggest {
		for _, v1 := range v[0].Options {
			_ = json.Unmarshal(v1.Source, &officer)
			officer.Uuid = convert.IntToString(convert.StrToInt(officer.Uuid) + config.Conf().Encryption.SecretKey)
			officers = append(officers, officer)
		}
	}

	return officers, nil
}

// @Description 手动更新
// @param ctx
// @param uuid
// @param update
// @return val
// @return err
func ManualUpdate(ctx *gin.Context, uuid string, update ManualUpdateRequest) (val bool, err error) {

	var officers Officers
	if officers, err = getInfoByUuid(ctx, uuid); err != nil {
		return false, err
	}

	doc, _ := convert.StructToMap(update)

	if _, err = es.Client().Update().Index(indices).Id(officers.EsId).Doc(doc).Do(ctx); err != nil {
		return false, err
	}
	return true, nil
}
