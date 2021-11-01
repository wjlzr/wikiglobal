package mgcompany

import (
	"wiki_global/src/db/mongo"

	mongo2 "go.mongodb.org/mongo-driver/mongo"
)

type Company struct {
	Uuid                                string            `json:"uuid" bson:"uuid"`
	CompanyNumber                       string            `json:"company_number" bson:"company_number"`
	JurisdictionCode                    string            `json:"jurisdiction_code" bson:"jurisdiction_code"`
	Name                                string            `json:"name" bson:"name"`
	NormalisedName                      string            `json:"normalised_name" bson:"normalised_name"`
	CompanyType                         string            `json:"company_type" bson:"company_type"`
	Nonprofit                           string            `json:"nonprofit" bson:"nonprofit"`
	CurrentStatus                       string            `json:"current_status" bson:"current_status"`
	IncorporationDate                   string            `json:"incorporation_date" bson:"incorporation_date"`
	DissolutionDate                     string            `json:"dissolution_date" bson:"dissolution_date"`
	Branch                              string            `json:"branch" bson:"branch"`
	BusinessNumber                      string            `json:"business_number" bson:"business_number"`
	CurrentAlternativeLegalName         string            `json:"current_alternative_legal_name" bson:"current_alternative_legal_name"`
	CurrentAlternativeLegalNameLanguage string            `json:"current_alternative_legal_name_language" bson:"current_alternative_legal_name_language"`
	HomeJurisdictionText                string            `json:"home_jurisdiction_text" bson:"home_jurisdiction_text"`
	NativeCompanyNumber                 string            `json:"native_company_number" bson:"native_company_number"`
	PreviousNames                       string            `json:"previous_names" bson:"previous_names"`
	AlternativeNames                    string            `json:"alternative_names" bson:"alternative_names"`
	RetrievedAt                         string            `json:"retrieved_at" bson:"retrieved_at"`
	RegistryUrl                         string            `json:"registry_url" bson:"registry_url"`
	RestrictedForMarketing              string            `json:"restricted_for_marketing" bson:"restricted_for_marketing"`
	Inactive                            string            `json:"inactive" bson:"inactive"`
	AccountsNextDue                     string            `json:"accounts_next_due" bson:"accounts_next_due"`
	AccountsReferenceDate               string            `json:"accounts_reference_date" bson:"accounts_reference_date"`
	AccountsLastMadeUpDate              string            `json:"accounts_last_made_up_date" bson:"accounts_last_made_up_date"`
	AnnualReturnNextDue                 string            `json:"annual_return_next_due" bson:"annual_return_next_due"`
	AnnualReturnLastMadeUpDate          string            `json:"annual_return_last_made_up_date" bson:"annual_return_last_made_up_date"`
	HasBeenLiquidated                   string            `json:"has_been_liquidated" bson:"has_been_liquidated"`
	HasInsolvencyHistory                string            `json:"has_insolvency_history" bson:"has_insolvency_history"`
	HasCharges                          string            `json:"has_charges" bson:"has_charges"`
	HomeJurisdictionCode                string            `json:"home_jurisdiction_code" bson:"home_jurisdiction_code"`
	HomeJurisdictionCompanyNumber       string            `json:"home_jurisdiction_company_number" bson:"home_jurisdiction_company_number"`
	IndustryCodeUids                    string            `json:"industry_code_uids" bson:"industry_code_uids"`
	LatestAccountsDate                  string            `json:"latest_accounts_date" bson:"latest_accounts_date"`
	LatestAccountsCash                  string            `json:"latest_accounts_cash" bson:"latest_accounts_cash"`
	LatestAccountsAssets                string            `json:"latest_accounts_assets" bson:"latest_accounts_assets"`
	LatestAccountsLiabilities           string            `json:"latest_accounts_liabilities" bson:"latest_accounts_liabilities"`
	RegisteredAddress                   registeredAddress `json:"registered_address" bson:"registered_address"`
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

// 返回集合
func C_collection(skip int64) *mongo2.Collection {

	return mongo.Client("wikiGlobal").Collection(CollectionName(skip))
}

func CollectionName(skip int64) string {

	company := ""
	if skip == 20000000000 {
		company = "company"
	} else {
		company = "company1"
		if skip > 19999999 && skip <= 39999998 {
			company = "company2"
		} else if skip > 39999998 && skip <= 59999997 {
			company = "company3"
		} else if skip > 59999997 && skip <= 79999996 {
			company = "company4"
		} else if skip > 79999996 && skip <= 99999995 {
			company = "company5"
		} else if skip > 99999995 && skip <= 119999994 {
			company = "company6"
		} else if skip > 119999994 && skip <= 139999993 {
			company = "company7"
		} else if skip > 139999993 && skip <= 159999992 {
			company = "company8"
		} else if skip > 159999992 && skip <= 179999991 {
			company = "company9"
		} else if skip > 179999991 && skip <= 200000000 {
			company = "company10"
		}
	}
	return company
}
