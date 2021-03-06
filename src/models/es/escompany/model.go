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
	Website                             string                `json:"url" bson:"url"` // ??????
	CompanyLog                          string                `json:"company_log" bson:"company_log"`
	IsAttention                         bool                  `json:"is_attention" bson:"is_attention"`
	Region                              string                `json:"region" bson:"region"`
	NationalFlag                        bool                  `json:"national_flag" bson:"national_flag"`
	Phone                               string                `json:"phone"`    // ??????
	Synopsis                            string                `json:"synopsis"` // ??????
	Twitter                             string                `json:"twitter"`
	Facebook                            string                `json:"facebook"`
	Linkedin                            string                `json:"linkedin"`
	Score                               float64               `json:"score"`                  // ??????
	Scores                              Scores                `json:"scores"`                 // ??????
	Platform                            int64                 `json:"platform"`               // ??????
	NoSearchStatus                      int64                 `json:"no_search_status"`       // ????????????
	NoSearchClosingTime                 int64                 `json:"no_search_closing_time"` // ????????????
	IsVr                                bool                  `json:"isVr"`                   // ?????????vr
	Color                               string                `json:"color"`                  // ????????????
	Annotation                          string                `json:"annotation"`             // ??????
	RegisterCountry                     string                `json:"registerCountry"`        // ?????????
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
	Subsidiary                          []Subsidiary          `json:"subsidiary" bson:"subsidiary"`               // ?????????
	Financing                           []Financing           `json:"financing" bson:"financing"`                 // ????????????
	Ipo                                 Ipo                   `json:"ipo" bson:"ipo"`                             // IPO
	Investor                            []Investor            `json:"investor" bson:"investor"`                   // ?????????
	InvestmentCase                      []Subsidiary          `json:"investment_case" bson:"investment_case"`     // ????????????
	AssetAcquisition                    []AssetAcquisition    `json:"asset_acquisition" bson:"asset_acquisition"` // ????????????
}

// ????????????
type Scores struct {
	ZhCN float64 `json:"zh-CN"`
	En   float64 `json:"en"`
}

// AssetAcquisition ????????????
type AssetAcquisition struct {
	Uuid          string `json:"uuid" bson:"uuid"`
	CompanyNumber string `json:"company_number"` // ????????????
	Name          string `json:"name"`           // ??????
	CompanyLog    string `json:"company_log"`
	OpenDate      int64  `json:"open_date"`
	Affair        string `json:"affair"`
	Price         struct {
		Money string `json:"money"` // ??????
		Unit  string `json:"unit"`  // ??????
	} `json:"price"`
}

// Investor ????????????
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
	ListedFunds      string `json:"listed_funds"`       // ??????????????????
	ListedSharePrice string `json:"listed_share_price"` // ????????????
	ListingDate      int64  `json:"listing_date"`       // ????????????
}

// Financing ????????????
type Financing struct {
	OpenDate          int64  `json:"open_date"`
	Category          string `json:"category"`
	NumberOfInvestors int64  `json:"number_of_investors"`
	Investor          string `json:"investor"`
	FundRaising       struct {
		Capital string `json:"capital"` // ??????
		Unit    string `json:"unit"`    // ??????
	} `json:"fund_raising"`
}

// Subsidiary ?????????
type Subsidiary struct {
	Uuid          string `json:"uuid" bson:"uuid"`
	CompanyNumber string `json:"company_number" bson:"company_number"` // ????????????
	Name          string `json:"name" bson:"name"`                     // ??????
	CompanyLog    string `json:"company_log" bson:"company_log"`
	Describe      string `json:"describe"`
}

// ??????
type registeredAddress struct {
	StreetAddress string `json:"street_address" bson:"street_address"`
	Locality      string `json:"locality" bson:"locality"`
	Region        string `json:"region" bson:"region"`
	PostalCode    string `json:"postal_code" bson:"postal_code"`
	Country       string `json:"country" bson:"country"`
	InFull        string `json:"in_full" bson:"in_full"`
}

// Coordinate ??????
type Coordinate struct {
	Lat float64 `json:"lat" bson:"lat"` // ??????
	Lon float64 `json:"lon" bson:"lon"` // ??????
}

// ????????????
type officers struct {
	Uuid          string `json:"uuid" bson:"uuid"`
	CompanyNumber string `json:"company_number" bson:"company_number"` // ????????????
	Name          string `json:"name" bson:"name"`                     // ??????
	Title         string `json:"title" bson:"title"`                   // ??????
	Position      string `json:"position" bson:"position"`             // ??????
	StartDate     string `json:"start_date" bson:"start_date"`         // ??????????????????
	CurrentStatus string `json:"current_status" bson:"current_status"` // ????????????
	Occupation    string `json:"occupation" bson:"occupation"`         // ??????
}

// ??????
type dynamic struct {
	Date    string `json:"date" bson:"date"`       // ??????
	Content string `json:"content" bson:"content"` // ??????
}

// ????????????
type trademark struct {
	Title        string `json:"title" bson:"title"`                 // ??????
	Img          string `json:"img" bson:"img"`                     // ??????
	Register     string `json:"register" bson:"register"`           // ?????????
	NiceCategory string `json:"nice_category" bson:"nice_category"` // nice??????
	RegisterDate string `json:"register_date" bson:"register_date"` // ????????????
	ExpireDate   string `json:"expire_date" bson:"expire_date"`     // ????????????
}

// ????????????
type shareholder struct {
	Date      string `json:"date" bson:"date"`           // ??????
	Describe  string `json:"describe" bson:"describe"`   // ??????
	Mechanism string `json:"mechanism" bson:"mechanism"` // ??????
}

// ??????
type keepOnRecord struct {
	ApplyDate string `json:"apply_date" bson:"apply_date"` // ????????????
	Title     string `json:"title" bson:"title"`           // ??????
	Describe  string `json:"describe" bson:"describe"`     // ??????
}

// ????????????
type industryStandard struct {
	Code       string `json:"code" bson:"code"`               // ??????
	Describe   string `json:"describe" bson:"describe"`       // ??????
	CodeScheme string `json:"code_scheme" bson:"code_scheme"` // ????????????
}

// ????????????
type gazetteNotice struct {
	Date        string `json:"date" bson:"date"`               // ??????
	Publication string `json:"publication" bson:"publication"` // ?????????
	BeCareful   string `json:"be_careful" bson:"be_careful"`   // ??????
}

// ????????????
type branchs struct {
	Name             string `json:"name" bson:"name"`                           // ????????????
	JurisdictionCode string `json:"jurisdiction_code" bson:"jurisdiction_code"` // ?????????
	CurrentStatus    string `json:"current_status" bson:"current_status"`       // ??????
	StartDate        string `json:"start_date" bson:"start_date"`               // ????????????
	EndDate          string `json:"end_date" bson:"end_date"`                   // ????????????
}

// ??????
type shares struct {
	ShareholderName string  `json:"shareholder_name" bson:"shareholder_name"` // ????????????
	EstablishDate   string  `json:"establish_date" bson:"establish_date"`     // ????????????
	SharesNum       int     `json:"shares_num" bson:"shares_num"`             // ??????
	Turnout         turnout `json:"turnout" bson:"turnout"`                   // ?????????
}

// ?????????
type turnout struct {
	Start string `json:"start" bson:"start"` // ?????????
	End   string `json:"end" bson:"end"`     // ?????????
}

// ????????????
type finance struct {
	CurrentAssets string `json:"current_assets" bson:"current_assets"` // ????????????
	StartDate     string `json:"start_date" bson:"start_date"`         // ??????????????????
	EndDate       string `json:"end_date" bson:"end_date"`             // ??????????????????
}

// ??????????????????
type financialRegulation struct {
	Name             string `json:"name" bson:"name"`                           // ????????????
	JurisdictionCode string `json:"jurisdiction_code" bson:"jurisdiction_code"` // ??????/?????????
	CurrentStatus    string `json:"current_status" bson:"current_status"`       // ??????
}

// ????????????
type businessLicense struct {
	Name             string `json:"name" bson:"name"`                           // ????????????
	JurisdictionCode string `json:"jurisdiction_code" bson:"jurisdiction_code"` // ??????/?????????
	CurrentStatus    string `json:"current_status" bson:"current_status"`       // ??????
}

// ManualUpdateRequest ??????????????????
type ManualUpdateRequest struct {
	CompanyLog        string             `json:"company_log"`
	Phone             string             `json:"phone"`    // ??????
	Synopsis          string             `json:"synopsis"` // ??????
	Twitter           string             `json:"twitter"`
	Facebook          string             `json:"facebook"`
	Linkedin          string             `json:"linkedin"`
	IncorporationDate string             `json:"incorporation_date" bson:"incorporation_date"`
	Subsidiary        []Subsidiary       `json:"subsidiary" bson:"subsidiary"`               // ?????????
	Financing         []Financing        `json:"financing" bson:"financing"`                 // ????????????
	Ipo               Ipo                `json:"ipo" bson:"ipo"`                             // IPO
	Investor          []Investor         `json:"investor" bson:"investor"`                   // ?????????
	InvestmentCase    []Subsidiary       `json:"investment_case" bson:"investment_case"`     // ????????????
	AssetAcquisition  []AssetAcquisition `json:"asset_acquisition" bson:"asset_acquisition"` // ????????????
	RegisteredAddress registeredAddress  `json:"registered_address" bson:"registered_address"`
}

// AggResponse ??????response
type AggResponse struct {
	DocCountErrorUpperBound int64     `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int64     `json:"sum_other_doc_count"`
	Buckets                 []buckets `json:"buckets"`
}

type buckets struct {
	Key      string `json:"key"`
	DocCount int64  `json:"doc_count"`
}
