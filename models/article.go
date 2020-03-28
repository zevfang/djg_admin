package models

import (
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Model
	AuthID           int    `json:"auth_id" orm:"column(auth_id);"`
	UserName         string `json:"username" orm:"column(username);"`
	VideoNum         int    `json:"video_num" orm:"column(video_num);"`
	Title            string `json:"title" orm:"column(title);size(100);"`
	SubTitle         string `json:"sub_title" orm:"column(sub_title);size(100);"`
	Topic            string `json:"topic" orm:"column(topic);size(100);"`
	Cate             string `json:"cate" orm:"column(cate);size(100);"`
	Tags             string `json:"tags" orm:"column(tags);size(200);"`
	Desc             string `json:"desc" orm:"column(desc);size(500);"`
	Content          string `json:"content" orm:"column(content);type(text);"`
	CosCoverImageUrl string `json:"cos_cover_image_url" orm:"column(cos_cover_image_url);size(500);"`
	CosVideoUrl      string `json:"cos_video_url" orm:"column(cos_video_url);size(500);"`
	BJHArticleId     string `json:"bjh_article_id" orm:"column(bjh_article_id);size(500);"`
	State            int    `json:"state" orm:"column(state);"`
}

func (u *Article) TableName() string {
	return "djg_article"
}

func GetArticles(limit, offset int64) ([]Article, int64, error) {
	var users []Article
	db := orm.NewOrm()
	var err error
	_, err = db.Raw("select a.*,b.username from djg_article a left join djg_auth b on a.auth_id=b.id limit ? offset ?;", limit, offset).QueryRows(&users)
	if err != nil {
		return nil, 0, err
	}
	num, err := db.QueryTable(Article{}).Count()
	if err != nil {
		return nil, 0, err
	}
	return users, num, err
}
