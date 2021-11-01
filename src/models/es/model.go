package es

import (
	"github.com/k0kubun/pp"
	"github.com/olivere/elastic/v7"
	"time"
)

// 关键词字段
var (
	CompanyFields  = []string{"name^4", "company_number^3", "previous_names^2"}
	OfficersFields = []string{"name"}
)

// 查询参数
type SearchListRequest struct {
	Keyword   string              `json:"keyword"` // 关键词搜索
	Screen    map[string][]string `json:"screen"`  // 条件筛选
	Sort      int                 `json:"sort"`    // 排序 1.最早成立 2.最新成立
	Agg       string              `json:"agg"`
	Language  string              `json:"language"`
	Client    string              `json:"client"`
	PageIndex int                 `json:"page_index"`
	PageSize  int                 `json:"page_size"`
}

type Filter struct {
	MustQuery        []elastic.Query
	MustNotQuery     []elastic.Query
	ShouldQuery      []elastic.Query
	Filters          []elastic.Query
	Sorters          []elastic.Sorter
	MultiMatchQuery  *elastic.MultiMatchQuery
	FieldValueFactor *elastic.FieldValueFactorFunction
	Highlight        *elastic.Highlight
	TermsQuery       []*elastic.TermsQuery
	Aggregation      *elastic.TermsAggregation
	From             int //分页
	Size             int
}

// FilterCondition 筛选条件统一组装
func (c *SearchListRequest) FilterCondition(fields []string) (*Filter, *elastic.BoolQuery, *elastic.FunctionScoreQuery) {

	var filter Filter
	scoreQuery := elastic.NewFunctionScoreQuery()
	boolQuery := elastic.NewBoolQuery()

	// 多字段模糊匹配
	if c.Keyword != "" {
		filter.MultiMatchQuery = elastic.NewMultiMatchQuery(c.Keyword, fields...)
		boolQuery.Must(filter.MultiMatchQuery)
	}

	// 关键词高亮
	filter.Highlight = elastic.NewHighlight().PreTags("<font color='red'>").PostTags("</font>")
	for _, v := range fields {
		filter.Highlight.Fields(elastic.NewHighlighterField(v))
	}

	// 多条件筛选
	if len(c.Screen) != 0 {
		for k, v := range c.Screen {
			filter.TermsQuery = append(filter.TermsQuery, elastic.NewTermsQueryFromStrings(k, v...))
		}
	}

	// 聚合
	if c.Agg != "" {
		filter.Aggregation = elastic.NewTermsAggregation().Field(c.Agg).Size(1000).OrderByCountDesc()
	}

	if c.Sort != 0 {
		if c.Sort == 1 {
			// 按照成立时间最早排序 因为有部分公司没有成立时间 则暂时将没有成立时间的数据排除掉
			boolQuery.MustNot(elastic.NewTermQuery("incorporation_date", ""))
			filter.Sorters = append(filter.Sorters, elastic.NewFieldSort("incorporation_date").Asc())
		} else if c.Sort == 2 {
			filter.Sorters = append(filter.Sorters, elastic.NewFieldSort("incorporation_date").Desc())
		}
	}
	filter.Sorters = append(filter.Sorters, elastic.NewFieldSort("_score").Desc())
	// 组装筛选条件
	for _, term := range filter.TermsQuery {
		boolQuery.Must(term)
	}

	// pc端先过滤掉区块和交易商数据
	//boolQuery.MustNot(elastic.NewTermsQuery("platform", 1, 2))
	scoreQuery.Query(boolQuery)
	pp.Println(scoreQuery.Source())
	filter.From = (c.PageIndex - 1) * c.PageSize
	filter.Size = c.PageSize
	return &filter, boolQuery, scoreQuery
}

// FilterCondition1 筛选条件统一组装
func (c *SearchListRequest) FilterConditionV2(fields []string, language string) (*Filter, *elastic.BoolQuery, *elastic.FunctionScoreQuery) {

	var filter Filter
	scoreQuery := elastic.NewFunctionScoreQuery()
	boolQuery := elastic.NewBoolQuery()
	boolQuery1 := elastic.NewBoolQuery()
	boolQuery2 := elastic.NewBoolQuery()
	boolQuery3 := elastic.NewBoolQuery()

	if language != "en" && language != "zh-CN" {
		language = "en"
	}

	// 禁搜逻辑
	boolQuery.Must(boolQuery3.Should(boolQuery1.Must(elastic.NewTermQuery("no_search_status", 1), elastic.NewRangeQuery("no_search_closing_time").Lt(time.Now().Unix())), boolQuery2.Must(elastic.NewTermQuery("no_search_status", 0))))

	// 多字段模糊匹配
	if c.Keyword != "" {
		filter.MultiMatchQuery = elastic.NewMultiMatchQuery(c.Keyword, fields...)
		boolQuery.Must(filter.MultiMatchQuery)
	}

	// 关键词高亮
	filter.Highlight = elastic.NewHighlight().PreTags("<font color='red'>").PostTags("</font>")
	for _, v := range fields {
		filter.Highlight.Fields(elastic.NewHighlighterField(v))
	}

	// 多条件筛选
	if len(c.Screen) != 0 {
		for k, v := range c.Screen {
			filter.TermsQuery = append(filter.TermsQuery, elastic.NewTermsQueryFromStrings(k, v...))
		}
	}

	// 聚合
	if c.Agg != "" {
		filter.Aggregation = elastic.NewTermsAggregation().Field(c.Agg).Size(1000).OrderByCountDesc()
	}

	if c.Sort != 0 {
		if c.Sort == 1 {
			// 按照成立时间最早排序 因为有部分公司没有成立时间 则暂时将没有成立时间的数据排除掉
			boolQuery.MustNot(elastic.NewTermQuery("incorporation_date", ""))
			filter.Sorters = append(filter.Sorters, elastic.NewFieldSort("incorporation_date").Asc())
		} else if c.Sort == 2 {
			filter.Sorters = append(filter.Sorters, elastic.NewFieldSort("incorporation_date").Desc())
		}
	}
	filter.Sorters = append(filter.Sorters, elastic.NewFieldSort("_score").Desc())
	filter.Sorters = append(filter.Sorters, elastic.NewFieldSort("scores."+language).Desc())
	// 组装筛选条件
	for _, term := range filter.TermsQuery {
		boolQuery.Must(term)
	}
	scoreFun := elastic.NewFieldValueFactorFunction().Field("scores." + language)
	scoreQuery.Query(boolQuery)
	scoreQuery.AddScoreFunc(scoreFun)
	pp.Println(scoreQuery.Source())
	//pp.Println(bb)
	filter.From = (c.PageIndex - 1) * c.PageSize
	filter.Size = c.PageSize
	return &filter, boolQuery, scoreQuery
}

// FilterConditionByBrief 简约筛选条件统一组装
func (c SearchListRequest) FilterConditionByBrief() (*Filter, *elastic.BoolQuery) {

	var filter Filter
	boolQuery := elastic.NewBoolQuery()

	// 组装筛选条件
	for k, v := range c.Screen {
		boolQuery.Must(elastic.NewTermsQueryFromStrings(k, v...))
	}

	filter.From = (c.PageIndex - 1) * c.PageSize
	filter.Size = c.PageSize
	pp.Println(boolQuery.Source())
	return &filter, boolQuery
}
