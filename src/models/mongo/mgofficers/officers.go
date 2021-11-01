package mgofficers

import (
	"github.com/k0kubun/pp"
	"wiki_global/src/db/mongo"

	mongo2 "go.mongodb.org/mongo-driver/mongo"
)

type Officers struct {
	Uuid               string  `json:"uuid" bson:"uuid"`
	ID                 string  `json:"id" bson:"id"`
	CompanyNumber      string  `json:"company_number" bson:"company_number"`
	JurisdictionCode   string  `json:"jurisdiction_code" bson:"jurisdiction_code"`
	Name               string  `json:"name" bson:"name"`
	Title              string  `json:"title" bson:"title"`
	FirstName          string  `json:"first_name" bson:"first_name"`
	LastName           string  `json:"last_name" bson:"last_name"`
	Position           string  `json:"position" bson:"position"`
	StartDate          string  `json:"start_date" bson:"start_date"`
	PersonNumber       string  `json:"person_number" bson:"person_number"`
	PersonUid          string  `json:"person_uid" bson:"person_uid"`
	EndDate            string  `json:"end_date" bson:"end_date"`
	CurrentStatus      string  `json:"current_status" bson:"current_status"`
	Occupation         string  `json:"occupation" bson:"occupation"`
	Nationality        string  `json:"nationality" bson:"nationality"`
	CountryOfResidence string  `json:"country_of_residence" bson:"country_of_residence"`
	PartialDateOfBirth string  `json:"partial_date_of_birth" bson:"partial_date_of_birth"`
	Type               string  `json:"type" bson:"type"`
	RetrievedAt        string  `json:"retrieved_at" bson:"retrieved_at"`
	SourceUrl          string  `json:"source_url" bson:"source_url"`
	Address            address `json:"address" bson:"address"`
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

// 返回集合
func O_collection(skip int64) *mongo2.Collection {
	return mongo.Client("wikiGlobal").Collection("officers")
}

func CollectionName(skip int64) string {
	company := ""
	if skip == 20000000000 {
		company = "officers"
	} else {
		company = "officers1"
		if skip > 19999999 && skip <= 39999998 {
			company = "officers2"
		} else if skip > 39999998 && skip <= 59999997 {
			company = "officers3"
		} else if skip > 59999997 && skip <= 79999996 {
			company = "officers4"
		} else if skip > 79999996 && skip <= 99999995 {
			company = "officers5"
		} else if skip > 99999995 && skip <= 119999994 {
			company = "officers6"
		} else if skip > 119999994 && skip <= 139999993 {
			company = "officers7"
		} else if skip > 139999993 && skip <= 159999992 {
			company = "officers8"
		} else if skip > 159999992 && skip <= 179999991 {
			company = "officers9"
		} else if skip > 179999991 && skip <= 199999990 {
			company = "officers10"
		} else if skip > 199999990 && skip <= 219999989 {
			company = "officers11"
		} else if skip > 219999989 && skip <= 239999988 {
			company = "officers12"
		}
	}
	_, _ = pp.Println("==============集合：" + company + "============")
	return company
}
