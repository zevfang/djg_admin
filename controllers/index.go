package controllers

type IndexController struct {
	BaseController
}

func (c *IndexController) Prepare() {
	isLogin := c.GetUser()
	if !isLogin {
		c.Redirect("/login", 302)
		return
	}
}

// @router / [get]
func (c *IndexController) Get() {
	c.Layout = "layout.html"
	c.TplName = "index.html"
}
