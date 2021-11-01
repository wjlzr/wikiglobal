package region

import (
	orm "wiki_global/src/db/mysql"
)

type Region struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	CodeFull string `json:"code_full"`
	StateCn  string `json:"state_cn"`
	StateEn  string `json:"state_en"`
	IsHot    int    `json:"is_hot"`
}

type RegionResponse struct {
	Asia         []Region `json:"asia"`          // 亚洲
	Europe       []Region `json:"europe"`        // 欧洲
	Oceania      []Region `json:"oceania"`       // 大洋洲
	NorthAmerica []Region `json:"north_america"` // 北美洲
	SouthAmerica []Region `json:"south_america"` // 南美洲
	Africa       []Region `json:"africa"`        // 非洲
	IsHot        []Region `json:"is_hot"`        // 热门
}

// GetInfo 获取地区信息
func (r Region) GetInfo() (regionResponse RegionResponse, err error) {

	regionResponse.Asia = r.Query(map[string][]string{"state_en = ?": []string{"asia"}})
	regionResponse.Europe = r.Query(map[string][]string{"state_en = ?": []string{"europe"}})
	regionResponse.Oceania = r.Query(map[string][]string{"state_en = ?": []string{"oceania"}})
	regionResponse.NorthAmerica = r.Query(map[string][]string{"state_en = ?": []string{"north_america"}})
	regionResponse.SouthAmerica = r.Query(map[string][]string{"state_en = ?": []string{"south_america"}})
	regionResponse.Africa = r.Query(map[string][]string{"state_en = ?": []string{"africa"}})
	regionResponse.IsHot = r.Query(map[string][]string{"is_hot = ?": []string{"1"}})
	return
}

// 查询
func (r Region) Query(conditions map[string][]string) (regions []Region) {

	query := orm.Engine.Cols("id", "name", "code", "code_full", "state_cn", "state_en", "is_hot")

	for key, val := range conditions {
		for _, v := range val {
			query.Where(key, v)
		}
	}

	_ = query.Find(&regions)
	return
}
