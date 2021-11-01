package escompany

type Company struct {
	EsId                                string                `json:"es_id"`
	Code                                string                `json:"code"`
	Uuid                                string                `json:"uuid" bson:"uuid"`
	CompanyNumber                       string                `json:"company_number" bson:"company_number"`
	RealCompanyNumber                   string                `json:"real_company_number" bson:"real_company_number"`
	JurisdictionCode                    string                `json:"jurisdiction_code" bson:"jurisdiction_code"`
	JurisdictionCodeFull                string                `json:"jurisdiction_code_full" bson:"jurisdiction_code_full"`
	Name                                string                `json:"name" bson:"name"`
	RealName                            string                `json:"real_name" bson:"real_name"`
	NormalisedName                      string                `json:"normalised_name" bson:"normalised_name"`
	CompanyType                         string                `json:"company_type" bson:"company_type"`
	Nonprofit                           string                `json:"nonprofit" bson:"nonprofit"`
	CurrentStatus                       string                `json:"current_status" bson:"current_status"`
	IncorporationDate                   string                `json:"incorporation_date" bson:"incorporation_date"`
	DissolutionDate                     string                `json:"dissolution_date" bson:"dissolution_date"`
	Branch                              string                `json:"branch" bson:"branch"`
	BusinessNumber                      string                `json:"business_number" bson:"business_number"`
	CurrentAlternativeLegalName         string                `json:"current_alternative_legal_name" bson:"current_alternative_legal_name"`
	CurrentAlternativeLegalNameLanguage string                `json:"current_alternative_legal_name_language" bson:"current_alternative_legal_name_language"`
	HomeJurisdictionText                string                `json:"home_jurisdiction_text" bson:"home_jurisdiction_text"`
	NativeCompanyNumber                 string                `json:"native_company_number" bson:"native_company_number"`
	RealPreviousNames                   string                `json:"real_previous_names" bson:"real_previous_names"`
	PreviousNames                       string                `json:"previous_names" bson:"previous_names"`
	AlternativeNames                    string                `json:"alternative_names" bson:"alternative_names"`
	RetrievedAt                         string                `json:"retrieved_at" bson:"retrieved_at"`
	RegistryUrl                         string                `json:"registry_url" bson:"registry_url"`
	RestrictedForMarketing              string                `json:"restricted_for_marketing" bson:"restricted_for_marketing"`
	Inactive                            string                `json:"inactive" bson:"inactive"`
	AccountsNextDue                     string                `json:"accounts_next_due" bson:"accounts_next_due"`
	AccountsReferenceDate               string                `json:"accounts_reference_date" bson:"accounts_reference_date"`
	AccountsLastMadeUpDate              string                `json:"accounts_last_made_up_date" bson:"accounts_last_made_up_date"`
	AnnualReturnNextDue                 string                `json:"annual_return_next_due" bson:"annual_return_next_due"`
	AnnualReturnLastMadeUpDate          string                `json:"annual_return_last_made_up_date" bson:"annual_return_last_made_up_date"`
	HasBeenLiquidated                   string                `json:"has_been_liquidated" bson:"has_been_liquidated"`
	HasInsolvencyHistory                string                `json:"has_insolvency_history" bson:"has_insolvency_history"`
	HasCharges                          string                `json:"has_charges" bson:"has_charges"`
	HomeJurisdictionCode                string                `json:"home_jurisdiction_code" bson:"home_jurisdiction_code"`
	HomeJurisdictionCompanyNumber       string                `json:"home_jurisdiction_company_number" bson:"home_jurisdiction_company_number"`
	IndustryCodeUids                    string                `json:"industry_code_uids" bson:"industry_code_uids"`
	LatestAccountsDate                  string                `json:"latest_accounts_date" bson:"latest_accounts_date"`
	LatestAccountsCash                  string                `json:"latest_accounts_cash" bson:"latest_accounts_cash"`
	LatestAccountsAssets                string                `json:"latest_accounts_assets" bson:"latest_accounts_assets"`
	LatestAccountsLiabilities           string                `json:"latest_accounts_liabilities" bson:"latest_accounts_liabilities"`
	Website                             string                `json:"url" bson:"url"` // 官网
	CompanyLog                          string                `json:"company_log" bson:"company_log"`
	IsAttention                         bool                  `json:"is_attention" bson:"is_attention"`
	Region                              string                `json:"region" bson:"region"`
	NationalFlag                        bool                  `json:"national_flag" bson:"national_flag"`
	Phone                               string                `json:"phone"`    // 电话
	Synopsis                            string                `json:"synopsis"` // 简介
	Twitter                             string                `json:"twitter"`
	Facebook                            string                `json:"facebook"`
	Linkedin                            string                `json:"linkedin"`
	Score                               float64               `json:"score"`                  // 评分
	Scores                              Scores                `json:"scores"`                 // 评分
	Platform                            int64                 `json:"platform"`               // 类型
	NoSearchStatus                      int64                 `json:"no_search_status"`       // 是否禁搜
	NoSearchClosingTime                 int64                 `json:"no_search_closing_time"` // 禁搜时间
	IsVr                                bool                  `json:"isVr"`                   // 是否有vr
	Color                               string                `json:"color"`                  // 角标颜色
	Annotation                          string                `json:"annotation"`             // 角标
	RegisterCountry                     string                `json:"registerCountry"`        // 注册地
	RegisteredAddress                   registeredAddress     `json:"registered_address" bson:"registered_address"`
	Coordinate                          Coordinate            `json:"coordinate" bson:"coordinate"`
	Officers                            []officers            `json:"officers" bson:"officers"`
	Dynamic                             []dynamic             `json:"dynamic" bson:"dynamic"`
	Trademark                           []trademark           `json:"trademark" bson:"trademark"`
	Shareholder                         []shareholder         `json:"shareholder" bson:"shareholder"`
	KeepOnRecord                        []keepOnRecord        `json:"keep_on_record" bson:"keep_on_record"`
	IndustryStandard                    []industryStandard    `json:"industry_standard" bson:"industry_standard"`
	GazetteNotice                       []gazetteNotice       `json:"gazette_notice" bson:"gazette_notice"`
	Branchs                             []branchs             `json:"branchs" bson:"branchs"`
	Finance                             []finance             `json:"finance" bson:"finance"`
	FinancialRegulation                 []financialRegulation `json:"financial_regulation" bson:"financial_regulation"`
	BusinessLicense                     []businessLicense     `json:"business_license" bson:"business_license"`
	IssueOfShares                       []shares              `json:"issue_of_shares" bson:"issue_of_shares"`
	EquityInOtherCompanies              []shares              `json:"equity_in_other_companies" bson:"equity_in_other_companies"`
	Subsidiary                          []Subsidiary          `json:"subsidiary" bson:"subsidiary"`               // 子公司
	Financing                           []Financing           `json:"financing" bson:"financing"`                 // 融资轮数
	Ipo                                 Ipo                   `json:"ipo" bson:"ipo"`                             // IPO
	Investor                            []Investor            `json:"investor" bson:"investor"`                   // 投资者
	InvestmentCase                      []Subsidiary          `json:"investment_case" bson:"investment_case"`     // 投资案例
	AssetAcquisition                    []AssetAcquisition    `json:"asset_acquisition" bson:"asset_acquisition"` // 资产收购
}

// 地区分数
type Scores struct {
	ZhCN float64 `json:"zh-CN"`
	En   float64 `json:"en"`
}

// AssetAcquisition 资产收购
type AssetAcquisition struct {
	Uuid          string `json:"uuid" bson:"uuid"`
	CompanyNumber string `json:"company_number"` // 公司编号
	Name          string `json:"name"`           // 名称
	CompanyLog    string `json:"company_log"`
	OpenDate      int64  `json:"open_date"`
	Affair        string `json:"affair"`
	Price         struct {
		Money string `json:"money"` // 资金
		Unit  string `json:"unit"`  // 单位
	} `json:"price"`
}

// Investor 投资大佬
type Investor struct {
	Name             string `json:"name"`
	IsMajorInvestors bool   `json:"is_major_investors"`
	Financing        string `json:"financing"`
	Partner          string `json:"partner"`
}

// Ipo IPO
type Ipo struct {
	Describe         string `json:"describe"`
	StockCode        string `json:"stock_code"`
	ListingValuation string `json:"listing_valuation"`
	ListedFunds      string `json:"listed_funds"`       // 上市募集资金
	ListedSharePrice string `json:"listed_share_price"` // 上市股价
	ListingDate      int64  `json:"listing_date"`       // 上市日期
}

// Financing 融资轮数
type Financing struct {
	OpenDate          int64  `json:"open_date"`
	Category          string `json:"category"`
	NumberOfInvestors int64  `json:"number_of_investors"`
	Investor          string `json:"investor"`
	FundRaising       struct {
		Capital string `json:"capital"` // 资金
		Unit    string `json:"unit"`    // 单位
	} `json:"fund_raising"`
}

// Subsidiary 子公司
type Subsidiary struct {
	Uuid          string `json:"uuid" bson:"uuid"`
	CompanyNumber string `json:"company_number" bson:"company_number"` // 公司编号
	Name          string `json:"name" bson:"name"`                     // 名称
	CompanyLog    string `json:"company_log" bson:"company_log"`
	Describe      string `json:"describe"`
}

// 地址
type registeredAddress struct {
	StreetAddress string `json:"street_address" bson:"street_address"`
	Locality      string `json:"locality" bson:"locality"`
	Region        string `json:"region" bson:"region"`
	PostalCode    string `json:"postal_code" bson:"postal_code"`
	Country       string `json:"country" bson:"country"`
	InFull        string `json:"in_full" bson:"in_full"`
}

// Coordinate 坐标
type Coordinate struct {
	Lat float64 `json:"lat" bson:"lat"` // 维度
	Lon float64 `json:"lon" bson:"lon"` // 经度
}

// 公司高管
type officers struct {
	Uuid          string `json:"uuid" bson:"uuid"`
	CompanyNumber string `json:"company_number" bson:"company_number"` // 公司编号
	Name          string `json:"name" bson:"name"`                     // 名称
	Title         string `json:"title" bson:"title"`                   // 称呼
	Position      string `json:"position" bson:"position"`             // 职位
	StartDate     string `json:"start_date" bson:"start_date"`         // 开始任职日期
	CurrentStatus string `json:"current_status" bson:"current_status"` // 任职状态
	Occupation    string `json:"occupation" bson:"occupation"`         // 职务
}

// 动态
type dynamic struct {
	Date    string `json:"date" bson:"date"`       // 日期
	Content string `json:"content" bson:"content"` // 内容
}

// 商标信息
type trademark struct {
	Title        string `json:"title" bson:"title"`                 // 标题
	Img          string `json:"img" bson:"img"`                     // 图片
	Register     string `json:"register" bson:"register"`           // 寄存器
	NiceCategory string `json:"nice_category" bson:"nice_category"` // nice分类
	RegisterDate string `json:"register_date" bson:"register_date"` // 注册日期
	ExpireDate   string `json:"expire_date" bson:"expire_date"`     // 到期日期
}

// 股东信息
type shareholder struct {
	Date      string `json:"date" bson:"date"`           // 日期
	Describe  string `json:"describe" bson:"describe"`   // 描述
	Mechanism string `json:"mechanism" bson:"mechanism"` // 机制
}

// 备案
type keepOnRecord struct {
	ApplyDate string `json:"apply_date" bson:"apply_date"` // 申请日期
	Title     string `json:"title" bson:"title"`           // 标题
	Describe  string `json:"describe" bson:"describe"`     // 描述
}

// 行业规范
type industryStandard struct {
	Code       string `json:"code" bson:"code"`               // 代码
	Describe   string `json:"describe" bson:"describe"`       // 描述
	CodeScheme string `json:"code_scheme" bson:"code_scheme"` // 代码方案
}

// 宪报公告
type gazetteNotice struct {
	Date        string `json:"date" bson:"date"`               // 日期
	Publication string `json:"publication" bson:"publication"` // 出版物
	BeCareful   string `json:"be_careful" bson:"be_careful"`   // 注意
}

// 分支机构
type branchs struct {
	Name             string `json:"name" bson:"name"`                           // 公司名称
	JurisdictionCode string `json:"jurisdiction_code" bson:"jurisdiction_code"` // 区域码
	CurrentStatus    string `json:"current_status" bson:"current_status"`       // 状态
	StartDate        string `json:"start_date" bson:"start_date"`               // 开始时间
	EndDate          string `json:"end_date" bson:"end_date"`                   // 结束时间
}

// 股份
type shares struct {
	ShareholderName string  `json:"shareholder_name" bson:"shareholder_name"` // 股东名称
	EstablishDate   string  `json:"establish_date" bson:"establish_date"`     // 成立日期
	SharesNum       int     `json:"shares_num" bson:"shares_num"`             // 股数
	Turnout         turnout `json:"turnout" bson:"turnout"`                   // 投票率
}

// 投票率
type turnout struct {
	Start string `json:"start" bson:"start"` // 起始率
	End   string `json:"end" bson:"end"`     // 截止率
}

// 财务摘要
type finance struct {
	CurrentAssets string `json:"current_assets" bson:"current_assets"` // 当前资产
	StartDate     string `json:"start_date" bson:"start_date"`         // 开始任职日期
	EndDate       string `json:"end_date" bson:"end_date"`             // 开始任职日期
}

// 金融监管信息
type financialRegulation struct {
	Name             string `json:"name" bson:"name"`                           // 公司名称
	JurisdictionCode string `json:"jurisdiction_code" bson:"jurisdiction_code"` // 国家/区域码
	CurrentStatus    string `json:"current_status" bson:"current_status"`       // 状态
}

// 营业执照
type businessLicense struct {
	Name             string `json:"name" bson:"name"`                           // 公司名称
	JurisdictionCode string `json:"jurisdiction_code" bson:"jurisdiction_code"` // 国家/区域码
	CurrentStatus    string `json:"current_status" bson:"current_status"`       // 状态
}

// ManualUpdateRequest 手动更新参数
type ManualUpdateRequest struct {
	CompanyLog        string             `json:"company_log"`
	Phone             string             `json:"phone"`    // 电话
	Synopsis          string             `json:"synopsis"` // 简介
	Twitter           string             `json:"twitter"`
	Facebook          string             `json:"facebook"`
	Linkedin          string             `json:"linkedin"`
	IncorporationDate string             `json:"incorporation_date" bson:"incorporation_date"`
	Subsidiary        []Subsidiary       `json:"subsidiary" bson:"subsidiary"`               // 子公司
	Financing         []Financing        `json:"financing" bson:"financing"`                 // 融资轮数
	Ipo               Ipo                `json:"ipo" bson:"ipo"`                             // IPO
	Investor          []Investor         `json:"investor" bson:"investor"`                   // 投资者
	InvestmentCase    []Subsidiary       `json:"investment_case" bson:"investment_case"`     // 投资案例
	AssetAcquisition  []AssetAcquisition `json:"asset_acquisition" bson:"asset_acquisition"` // 资产收购
	RegisteredAddress registeredAddress  `json:"registered_address" bson:"registered_address"`
}

// AggResponse 聚合response
type AggResponse struct {
	DocCountErrorUpperBound int64     `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int64     `json:"sum_other_doc_count"`
	Buckets                 []buckets `json:"buckets"`
}

type buckets struct {
	Key      string `json:"key"`
	DocCount int64  `json:"doc_count"`
}
