package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Auth struct {
	ID        int       `json:"id" orm:"column(id);size(100);pk"`
	Username  string    `json:"username" orm:"column(username);size(50);"`
	Password  string    `json:"password" orm:"column(password);size(50);"`
	AppID     string    `json:"app_id" orm:"column(app_id);size(200);"`
	AppToken  string    `json:"app_token" orm:"column(app_token);size(200);"`
	SecretID  string    `json:"secret_id" orm:"column(secret_id);size(200);"`
	SecretKey string    `json:"secret_key" orm:"column(secret_key);size(200);"`
	HKIndex   string    `json:"hk_index" orm:"column(hk_index);size(400);"`
	State     int       `json:"state" orm:"column(state);"`
	CreatedOn time.Time `json:"created_on" orm:"column(created_on);type(datetime)"`
}

func (u *Auth) TableName() string {
	return "djg_auth"
}

func GetAuths() ([]Auth, error) {
	var auths []Auth
	db := orm.NewOrm()
	_, err := db.Raw("select * from djg_auth;").QueryRows(&auths)
	if err != nil {
		return nil, err
	}
	return auths, err
}
