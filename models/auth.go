package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Auth struct {
	ID         int       `json:"id" orm:"column(id);size(100);pk"`
	Username   string    `json:"username" orm:"column(username);size(50);"`
	Password   string    `json:"password" orm:"column(password);size(50);"`
	AppID      string    `json:"app_id" orm:"column(app_id);size(200);"`
	AppToken   string    `json:"app_token" orm:"column(app_token);size(200);"`
	SecretID   string    `json:"secret_id" orm:"column(secret_id);size(200);"`
	SecretKey  string    `json:"secret_key" orm:"column(secret_key);size(200);"`
	CosAddress string    `json:"cos_address" orm:"column(cos_address);size(200);"`
	HKIndex    string    `json:"hk_index" orm:"column(hk_index);size(400);"`
	State      int       `json:"state" orm:"column(state);"`
	CreatedOn  time.Time `json:"created_on" orm:"column(created_on);type(datetime)"`
}

type AuthSelect struct {
	Name     string `json:"name"`
	Value    int    `json:"value"`
	Velected string `json:"velected"` //selected
	Disabled string `json:"disabled"` //disabled
}

func (u *Auth) TableName() string {
	return "djg_auth"
}

func GetAuthsAll() ([]Auth, error) {
	var auths []Auth
	db := orm.NewOrm()
	_, err := db.Raw("select * from djg_auth").QueryRows(&auths)
	if err != nil {
		return nil, err
	}
	return auths, err
}

func GetAuths(username string) ([]Auth, error) {
	var where string = " where 1=1 "
	if len(username) > 0 {
		where += " and username like '%" + username + "%' "
	}

	var auths []Auth
	db := orm.NewOrm()
	sqlStr := fmt.Sprintf("select * from djg_auth %s;", where)
	_, err := db.Raw(sqlStr).QueryRows(&auths)
	if err != nil {
		return nil, err
	}
	return auths, err
}

func GetAuthByIds(ids string) ([]Auth, error) {
	var where string = " where 1=1 "
	if len(ids) > 0 {
		where += " and id in (" + ids + ")"
	}

	var auths []Auth
	db := orm.NewOrm()
	sqlStr := fmt.Sprintf("select * from djg_auth %s;", where)
	_, err := db.Raw(sqlStr).QueryRows(&auths)
	if err != nil {
		return nil, err
	}
	return auths, err
}

func EditAuth(model Auth) error {
	db := orm.NewOrm()
	_, err := db.Update(&model)
	return err
}
