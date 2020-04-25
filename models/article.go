package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	Model
	AuthID           int       `json:"auth_id" orm:"column(auth_id);"`
	VideoNum         int       `json:"video_num" orm:"column(video_num);"`
	Title            string    `json:"title" orm:"column(title);size(100);"`
	SubTitle         string    `json:"sub_title" orm:"column(sub_title);size(100);"`
	Topic            string    `json:"topic" orm:"column(topic);size(100);"`
	Cate             string    `json:"cate" orm:"column(cate);size(100);"`
	Tags             string    `json:"tags" orm:"column(tags);size(200);"`
	Desc             string    `json:"desc" orm:"column(desc);size(500);"`
	Content          string    `json:"content" orm:"column(content);type(text);"`
	CosCoverImageUrl string    `json:"cos_cover_image_url" orm:"column(cos_cover_image_url);size(500);"`
	CosVideoUrl      string    `json:"cos_video_url" orm:"column(cos_video_url);size(500);"`
	BJHArticleUrl    string    `json:"bjh_article_url" orm:"column(bjh_article_url);size(500);"`
	BJHArticleId     string    `json:"bjh_article_id" orm:"column(bjh_article_id);size(500);"`
	State            int       `json:"state" orm:"column(state);"`
	VerifyOn         time.Time `json:"verify_on" orm:"column(verify_on);"`
	Remark           string    `json:"remark" orm:"column(remark);size(500);"`
}

type FindArticle struct {
	Article
	UserName string `json:"username" orm:"column(username);"`
	StateVal string `json:"state_val"`
}

type ArticleTotal struct {
	Count int64 `json:"count" orm:"column(count);"`
}

func (u *Article) TableName() string {
	return "djg_article"
}

func GetArticles(limit, offset int64, auth_id, state int64, verify_on string) ([]FindArticle, int64, error) {
	var err error
	var users []FindArticle
	db := orm.NewOrm()
	field := " a.*,b.username "
	//条件
	var where string = " where 1=1 "
	if auth_id >= 0 {
		where += fmt.Sprintf(" and a.auth_id=%d  ", auth_id)
	}
	if state >= 0 {
		where += fmt.Sprintf(" and a.state=%d  ", state)
	}
	if len(verify_on) > 0 {
		where += fmt.Sprintf(" and a.verify_on = '%s'  ", verify_on)
	}

	//分页
	limitOffset := fmt.Sprintf(" limit %d offset %d ", limit, offset)
	//排序
	orderBy := " order by a.created_on desc "
	sql := fmt.Sprintf("select %s from djg_article a left join djg_auth b on a.auth_id=b.id %s %s %s;", field, where, orderBy, limitOffset)
	_, err = db.Raw(sql).QueryRows(&users)
	if err != nil {
		return nil, 0, err
	}

	var rowCount int64
	sqlCount := fmt.Sprintf("select count(1) from djg_article a left join djg_auth b on a.auth_id=b.id %s ;", where)
	err = db.Raw(sqlCount).QueryRow(&rowCount)
	if err != nil {
		return nil, 0, err
	}
	return users, rowCount, err
}

func EditArticle(model Article) error {
	db := orm.NewOrm()
	_, err := db.Update(&model)
	return err
}

func AddArticle(model Article) error {
	db := orm.NewOrm()
	_, err := db.Insert(&model)
	return err
}

func GetArticle(article_id string) (Article, error) {
	var article Article
	db := orm.NewOrm()
	err := db.Raw("select * from djg_article where bjh_article_id=?;", article_id).QueryRow(&article)
	return article, err
}
