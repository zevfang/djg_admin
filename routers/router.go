package routers

import (
	"djg_admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.AddNamespace(NewNamespace())
}

func NewNamespace() *beego.Namespace {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "get:LogOut")

	np := beego.NewNamespace("/v1",
		beego.NSNamespace("/console",
			beego.NSInclude(
				&controllers.ConsoleController{},
			),
		),
		beego.NSNamespace("/article",
			beego.NSInclude(
				&controllers.ArticleController{},
			),
			beego.NSRouter("/get_articles", &controllers.ArticleController{}, "get:GetArticles"),
			beego.NSRouter("/get_article", &controllers.ArticleController{}, "get:GetArticle"),
			beego.NSRouter("/upd_article", &controllers.ArticleController{}, "post:UpdArticle"),
			beego.NSRouter("/add_article", &controllers.ArticleController{}, "post:AddArticle"),
		),
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.AuthController{},
			),
			beego.NSRouter("/get_auths", &controllers.AuthController{}, "get:GetAuths"),
			beego.NSRouter("/get_auths_select", &controllers.AuthController{}, "get:GetAuthsSelect"),
			beego.NSRouter("/get_auth", &controllers.AuthController{}, "get:GetAuth"),
			beego.NSRouter("/add_auth", &controllers.AuthController{}, "get,post:AddAuth"),
			beego.NSRouter("/upd_auth", &controllers.AuthController{}, "post:UpdAuth"),
		),
		beego.NSNamespace("/ctr",
			beego.NSInclude(
				&controllers.CtrController{},
			),
			beego.NSRouter("/get_ctr_top", &controllers.CtrController{}, "get:GetCtrTop"),
		),
	)
	return np
}
