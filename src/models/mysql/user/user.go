package user

import (
	"go.uber.org/zap"
	orm "wiki_global/src/db/mysql"
	"wiki_global/src/utils/log"
)

type User struct {
	ID          int64  `json:"id"`
	UserId      string `json:"user_id"`
	OpenId      string `json:"open_id"`
	AccountType string `json:"account_type"`
	AreaCode    string `json:"area_code"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	CountryCode string `json:"country_code"`
	CreateAt    int64  `json:"create_at" xorm:"created"`
}

// Create 创建
func (u *User) Create() (user *User, err error) {

	if _, err = orm.Engine.Cols("id", "user_id", "open_id", "account_type", "area_code", "phone", "email", "password", "country_code", "create_at").InsertOne(u); err != nil {
		log.Logger().Error("model user Create InsertOne err：", zap.Error(err))
		return nil, err
	}
	return u, nil
}
