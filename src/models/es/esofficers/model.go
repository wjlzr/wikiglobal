package esofficers

type Officers struct {
	Id                   string  `json:"id" bson:"id"`
	Uuid                 string  `json:"uuid" bson:"uuid"`
	EsId                 string  `json:"es_id"`
	CompanyNumber        string  `json:"company_number" bson:"company_number"`
	RealCompanyNumber    string  `json:"real_company_number" bson:"real_company_number"`
	JurisdictionCode     string  `json:"jurisdiction_code" bson:"jurisdiction_code"`
	JurisdictionCodeFull string  `json:"jurisdiction_code_full" bson:"jurisdiction_code_full"`
	Name                 string  `json:"name" bson:"name"`
	RealName             string  `json:"real_name" bson:"real_name"`
	Title                string  `json:"title" bson:"title"`
	FirstName            string  `json:"first_name" bson:"first_name"`
	LastName             string  `json:"last_name" bson:"last_name"`
	Position             string  `json:"position" bson:"position"`
	StartDate            string  `json:"start_date" bson:"start_date"`
	PersonNumber         string  `json:"person_number" bson:"person_number"`
	PersonUid            string  `json:"person_uid" bson:"person_uid"`
	EndDate              string  `json:"end_date" bson:"end_date"`
	CurrentStatus        string  `json:"current_status" bson:"current_status"`
	Occupation           string  `json:"occupation" bson:"occupation"`
	Nationality          string  `json:"nationality" bson:"nationality"`
	CountryOfResidence   string  `json:"country_of_residence" bson:"country_of_residence"`
	PartialDateOfBirth   string  `json:"partial_date_of_birth" bson:"partial_date_of_birth"`
	Type                 string  `json:"type" bson:"type"`
	IsAttention          bool    `json:"is_attention" bson:"is_attention"`
	Address              address `json:"address" bson:"address"`
	RetrievedAt          string  `json:"retrieved_at" bson:"retrieved_at"`
	SourceUrl            string  `json:"source_url" bson:"source_url"`
	Region               string  `json:"region" bson:"region"`
	HeadPortrait         string  `json:"head_portrait" bson:"head_portrait"`
	Company              company `json:"company" bson:"company"`
}

// 地址
type address struct {
	StreetAddress string `json:"street_address" bson:"street_address"`
	Locality      string `json:"locality" bson:"locality"`
	Region        string `json:"region" bson:"region"`
	PostalCode    string `json:"postal_code" bson:"postal_code"`
	Country       string `json:"country" bson:"country"`
	InFull        string `json:"in_full" bson:"in_full"`
}

// 所属公司
type company struct {
	Uuid                 string `json:"uuid" bson:"uuid"`
	CompanyNumber        string `json:"company_number" bson:"company_number"` // 公司编号
	Name                 string `json:"name" bson:"name"`                     // 名称
	JurisdictionCode     string `json:"jurisdiction_code" bson:"jurisdiction_code"`
	JurisdictionCodeFull string `json:"jurisdiction_code_full" bson:"jurisdiction_code_full"`
	Website              string `json:"url" bson:"url"` //网址
}

type ManualUpdateRequest struct {
	HeadPortrait string `json:"head_portrait" bson:"head_portrait"`
}
