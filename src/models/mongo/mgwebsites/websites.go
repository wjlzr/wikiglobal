package mgwebsites

import (
	"wiki_global/src/db/mongo"

	mongo2 "go.mongodb.org/mongo-driver/mongo"
)

type Websites struct {
	Uuid             string `bson:"_id"`
	CompanyNumber    string `bson:"company_number"`
	JurisdictionCode string `bson:"jurisdiction_code"`
	Url              string `bson:"url"`
	RawUrl           string `bson:"raw_url"`
	StartDate        string `bson:"start_date"`
	EndDate          string `bson:"end_date"`
}

// 返回集合
func W_collection() *mongo2.Collection {
	return mongo.Client("wikiGlobal").Collection("websites")
}
