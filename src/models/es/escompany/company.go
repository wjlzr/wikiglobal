package escompany

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"sync"
	"time"
	"wiki_global/src/common/convert"
	"wiki_global/src/config"
	"wiki_global/src/db/es"
	"wiki_global/src/models/mongo/mgcompany"
	"wiki_global/src/services/data"
	"wiki_global/src/services/googlemap"
	"wiki_global/src/utils/log"

	"github.com/gin-gonic/gin"

	es2 "wiki_global/src/models/es"

	"github.com/k0kubun/pp"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

const indices = "company_v1"
const indices2 = "company_v2"

func CreateIndex(mapping string) {
	if _, err := es.Client().CreateIndex(indices).BodyJson(mapping).Do(context.Background()); err != nil {
		log.Logger().Error("Company CreateIndex Error", zap.Error(err))
	}
}

func Insert(companys []Company, num int, group *sync.WaitGroup, skip int64) error {
	defer group.Done()
	if len(companys) == 0 {
		pp.Println("mongo数据为空直接返回")
		return nil
	}
	_, _ = pp.Println("----------第" + convert.IntToString(num*1) + "万条开始导入-----------")
	//_, err := es.Client().Index().Index(indices).BodyJson(company).Do(context.Background())
	//if err != nil {
	//	//_, _ = pp.Println("Insert company err: ", err)
	//	_, _ = pp.Println("Insert company err id: ", company.Uuid, err.Error())
	//	log.Logger().Info("Insert company id: "+company.Uuid, zap.String("err: ", err.Error()))
	//}
	bulkRequest := es.Client().Bulk()
	for _, company := range companys {
		doc := elastic.NewBulkIndexRequest().Index(indices).Doc(company)
		bulkRequest = bulkRequest.Add(doc)
	}

	_, err := bulkRequest.Do(context.Background())
	//pp.Println(result)
	if err != nil {
		log.Logger().Info("Insert company err: ", zap.Error(err))
		//log.Logger().Info("company: ", zap.Any("detail", companys))
		var uuids []string
		for _, v := range companys {
			uuids = append(uuids, v.Uuid)
		}
		b, err := json.Marshal(uuids)
		// 失败将uuid保存起来
		name := mgcompany.CollectionName(skip)
		if err = ioutil.WriteFile("c-json/"+name+"-"+convert.IntToString(num*1)+".json", b, os.ModeAppend); err != nil {
			pp.Println(name+"-"+convert.IntToString(num*1)+" 公司生成json文件失败: ", zap.Error(err))
			log.Logger().Info(name+"-"+convert.IntToString(num*1)+" 公司生成json文件失败: ", zap.Error(err))
		}
	}
	return nil
}

func Insert2(company Company) error {
	_, err := es.Client().Index().Index(indices).BodyJson(company).Do(context.Background())
	if err != nil {
		//_, _ = pp.Println("Insert company err: ", err)
		_, _ = pp.Println("Insert company err id: ", company.Uuid, err.Error())
		//log.Logger().Info("Insert company id: "+company.Uuid, zap.String("err: ", err.Error()))
	}
	return nil
}

// ListQuery 筛选
func ListQuery(ctx context.Context, language string, filter *es2.Filter, boolQuery *elastic.BoolQuery, score *elastic.FunctionScoreQuery) (companies []*Company, count int, err error) {

	resp, err := es.Client().Search().Index(indices).Highlight(filter.Highlight).Query(score).SortBy(filter.Sorters...).From(filter.From).Size(filter.Size).TrackTotalHits(true).Pretty(true).Do(ctx)
	if err != nil {
		return nil, 0, err
	}

	if resp.TotalHits() == 0 {
		return nil, 0, nil
	}

	var codes []string
	for _, e := range resp.Hits.Hits {
		var company *Company
		_ = json.Unmarshal(e.Source, &company)
		// uuid加密
		company.Uuid = convert.IntToString(convert.StrToInt(company.Uuid) + config.Conf().Encryption.SecretKey)
		company.RealCompanyNumber = company.CompanyNumber
		company.RealName = company.Name
		company.RealPreviousNames = company.PreviousNames

		// 处理地区code为空
		if company.JurisdictionCode == "" {
			company.JurisdictionCode = "--"
		}

		// 组装code
		if (company.Platform == 1 || company.Platform == 2) && language != "" {
			codes = append(codes, company.Code)
		}
		for k, h := range e.Highlight {
			if k == "name" {
				company.Name = h[0]
			} else if k == "company_number" {
				company.CompanyNumber = h[0]
			} else if k == "previous_names" {
				company.PreviousNames = h[0]
			}
		}
		companies = append(companies, company)
	}

	// 获取动态数据
	start1 := time.Now()
	if len(codes) != 0 {
		result, err := data.GetSearchData(codes, language)
		if err != nil {
			return nil, 0, errors.New("500")
		}
		pp.Println(result.Result)
		if len(result.Result) != 0 {
			for k, v := range companies {
				for _, v1 := range result.Result {
					if v.Code == v1.Code {
						companies[k].IsVr = v1.IsVr
						companies[k].Color = v1.Color
						companies[k].Annotation = v1.Annotation
						companies[k].CompanyLog = v1.Logo
						companies[k].RegisterCountry = v1.RegisterCountry
						companies[k].NormalisedName = v1.LocalFullName
					}
				}
			}
		}
	}
	_, _ = pp.Println("=============请求动态接口数据耗时: =============", time.Since(start1)/1000000)

	return companies, int(resp.TotalHits()), nil
}

// GetCondition 获取筛选条件聚合
func GetCondition(ctx context.Context, filter *es2.Filter, boolQuery *elastic.BoolQuery) (buckets []buckets, err error) {

	resp, err := es.Client().Search().Index(indices).Query(boolQuery).Aggregation("aggregation", filter.Aggregation).Pretty(true).Do(ctx)
	if err != nil {
		return nil, err
	}

	if resp.TotalHits() == 0 {
		return nil, nil
	}

	var aggResponse AggResponse
	var _ = json.Unmarshal(resp.Aggregations["aggregation"], &aggResponse)

	return aggResponse.Buckets, nil
}

// DetailQuery 详情
func DetailQuery(ctx *gin.Context, uuid string) (company Company, err error) {

	if company, err = getInfoByUuid(ctx, uuid); err != nil {
		return Company{}, err
	}

	// 查询经纬度
	coordinate, err := googlemap.FindCoordinateByAddress(company.RegisteredAddress.InFull)
	if err == nil {
		company.Coordinate.Lat = coordinate.Lat
		company.Coordinate.Lon = coordinate.Lng
	}

	return company, nil
}

// CompletionSuggestQuery 完成推荐
func CompletionSuggestQuery(ctx *gin.Context, prefix string) (companies []Company, err error) {

	suggest := elastic.NewCompletionSuggester("company-suggest").Prefix(prefix).Field("name.relevant")

	resp, err := es.Client().Search().Index(indices).Suggester(suggest).Do(ctx)
	if err != nil {
		return nil, err
	}

	var company Company
	for _, v := range resp.Suggest {
		for _, v1 := range v[0].Options {
			_ = json.Unmarshal(v1.Source, &company)
			company.Uuid = convert.IntToString(convert.StrToInt(company.Uuid) + config.Conf().Encryption.SecretKey)
			companies = append(companies, company)
		}
	}

	return companies, nil
}

// ManualUpdate 手动更新
func ManualUpdate(ctx *gin.Context, uuid string, update ManualUpdateRequest) (val bool, err error) {

	var company Company
	if company, err = getInfoByUuid(ctx, uuid); err != nil {
		return false, err
	}

	doc, _ := convert.StructToMap(update)

	if _, err = es.Client().Update().Index(indices).Id(company.EsId).Doc(doc).Do(ctx); err != nil {
		return false, err
	}
	return true, nil
}

// ListQueryByBrief 公司筛选简约版
func ListQueryByBrief(ctx context.Context, filter *es2.Filter, boolQuery *elastic.BoolQuery) (companies []*Company, count int, err error) {
	resp, err := es.Client().Search().Index(indices).Query(boolQuery).From(filter.From).Size(filter.Size).TrackTotalHits(true).Pretty(true).Do(ctx)
	if err != nil {
		log.Logger().Error("公司筛选简约版 Err：", zap.Error(err))
		return nil, 0, err
	}

	if resp.TotalHits() == 0 {
		return nil, 0, nil
	}

	for _, e := range resp.Each(reflect.TypeOf(&Company{})) {
		c := e.(*Company)
		c.Uuid = convert.IntToString(convert.StrToInt(c.Uuid) + config.Conf().Encryption.SecretKey)
		companies = append(companies, c)
	}
	return companies, int(resp.TotalHits()), nil
}

// UpdateByCode 更新操作
func UpdateByCode(code string, doc map[string]interface{}) (val bool, err error) {

	var company Company
	if company, err = getInfoByCode(code); err != nil {
		return false, err
	}

	if _, err = es.Client().Update().Index(indices).Id(company.EsId).Doc(doc).Do(context.Background()); err != nil {
		return false, err
	}

	return true, nil
}

// DeleteByCode 删除操作
func DeleteByCode(code string) (val bool, err error) {

	var company Company
	if company, err = getInfoByCode(code); err != nil {
		return false, err
	}

	if _, err = es.Client().Delete().Index(indices).Id(company.EsId).Do(context.Background()); err != nil {
		return false, err
	}

	return true, nil
}

func getInfoByUuid(ctx context.Context, uuid string) (company Company, err error) {

	pp.Println("uuid:", uuid)
	pp.Println("解密后的uuid:", convert.IntToString(convert.StrToInt(uuid)-config.Conf().Encryption.SecretKey))

	resp, err := es.Client().Search().Index(indices).Query(elastic.NewTermQuery("uuid", convert.IntToString(convert.StrToInt(uuid)-config.Conf().Encryption.SecretKey))).Pretty(true).Do(ctx)
	if err != nil {
		return Company{}, err
	}

	if resp.TotalHits() == 0 {
		return Company{}, errors.New("不存在此公司")
	}

	for _, hits := range resp.Hits.Hits {
		_ = json.Unmarshal([]byte(hits.Source), &company)
	}
	company.EsId = resp.Hits.Hits[0].Id
	company.Uuid = convert.IntToString(convert.StrToInt(company.Uuid) + config.Conf().Encryption.SecretKey)
	if len(company.Officers) != 0 {
		for k, v := range company.Officers {
			company.Officers[k].Uuid = convert.IntToString(convert.StrToInt(v.Uuid) + config.Conf().Encryption.SecretKey)
		}
	}
	return
}

// getInfoByCode
func getInfoByCode(code string) (company Company, err error) {

	resp, err := es.Client().Search().Index(indices).Query(elastic.NewTermQuery("code", code)).Pretty(true).Do(context.Background())
	if err != nil {
		return Company{}, err
	}

	if resp.TotalHits() == 0 {
		return Company{}, errors.New("不存在此公司")
	}

	for _, hits := range resp.Hits.Hits {
		_ = json.Unmarshal([]byte(hits.Source), &company)
	}
	company.EsId = resp.Hits.Hits[0].Id
	company.Uuid = convert.IntToString(convert.StrToInt(company.Uuid) + config.Conf().Encryption.SecretKey)
	if len(company.Officers) != 0 {
		for k, v := range company.Officers {
			company.Officers[k].Uuid = convert.IntToString(convert.StrToInt(v.Uuid) + config.Conf().Encryption.SecretKey)
		}
	}
	return
}

// GetResultByCode 根据code查询数据是否存在
func GetResultByCode(code string) bool {

	resp, err := es.Client().Search().Index(indices).Query(elastic.NewTermQuery("code", code)).Pretty(true).Do(context.Background())
	if err != nil {
		return false
	}

	if resp.TotalHits() == 0 {
		return false
	}

	return true
}
