package models

import "time"

type Auth struct {
	ID        int       `json:"id" orm:"size(100);pk"`
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
