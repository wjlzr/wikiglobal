package wikifx

// wikifx-特别提示-统计请求参数
type TraderCountReq struct {
	SandBox   int64  `json:"sandbox"`
	Equipld   string `json:"equipld"`
	UserId    string `json:"userid"`
	Platform  int64  `json:"platform"`
	AppType   int64  `json:"apptype"`
	Ip        string `json:"ip"`
	Ver       string `json:"ver"`
	EquipInfo string `json:"equipinfo"`
	Country   string `json:"country"`
	Lang      string `json:"lang"`
	Type      int64  `json:"type"`
	Code      string `json:"code"`
	Spots     int64  `json:"spots"`
	Url       string `json:"url"`
	Modal     int64  `json:"modal"`
}

// wikifx-经纪商历史走势
type BrokerHistoryDataReq struct {
	BrokerId string `json:"evaluation_code" structs:"brokerId"`
	Sort     string `json:"sort" structs:"sort"`
	Period   int64  `json:"period" structs:"period"`
}
