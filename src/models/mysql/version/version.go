package version

import (
	"database/sql/driver"
	"encoding/json"
	"go.uber.org/zap"
	orm "wiki_global/src/db/mysql"
	"wiki_global/src/utils/log"
)

type Version struct {
	ID              int64       `json:"id" xorm:"id"`
	Client          string      `json:"client" xorm:"client"`
	VersionNumber   string      `json:"version_number"`
	UpgradeType     int64       `json:"upgrade_type"`
	UpgradeTips     upgradeTips `json:"upgrade_tips"`
	AppStoreUrl     string      `json:"app_store_url"`
	DownloadPackage string      `json:"download_package"`
	GooglePlayUrl   string      `json:"google_play_url"`
	UpdateAt        int64       `json:"update_at" xorm:"created"`
	CreateAt        int64       `json:"create_at" xorm:"updated"`
}

type upgradeTips struct {
	En string `json:"en"`
	Cn string `json:"cn"`
}

func (*Version) TableName() string {
	return "version"
}

func (v upgradeTips) Value() (driver.Value, error) {
	b, err := json.Marshal(v)
	if err != nil {
		log.Logger().Error("Error Value json.Marshal err", zap.Error(err))
	}
	return string(b), err
}

//
func (v *upgradeTips) Scan(input interface{}) (err error) {
	if _, ok := input.(map[string]interface{}); ok {
		jsonStr, err := json.Marshal(input)
		if err != nil {
			log.Logger().Error("Error Scan json.Marshal err", zap.Error(err))
		}
		if err = json.Unmarshal([]byte(jsonStr), v); err != nil {
			log.Logger().Error("Error Scan json.Unmarshal1 err", zap.Error(err))
		}
	}
	if err = json.Unmarshal(input.([]byte), v); err != nil {
		log.Logger().Error("Error Scan json.Unmarshal2 err", zap.Error(err))
	}
	return
}

// QueryInfo 查询
func (receiver *Version) FindOne() (version []Version, err error) {

	err = orm.Engine.Table(receiver.TableName()).Cols("id", "client", "version_number", "upgrade_type", "upgrade_tips", "app_store_url", "download_package", "google_play_url", "update_at", "create_at").Where("client = ?", receiver.Client).OrderBy("id desc").Limit(1, 0).Find(&version)
	if err != nil {
		log.Logger().Error("Version FindOne Get err: ", zap.Error(err))
		return version, err
	}
	return version, nil
}
