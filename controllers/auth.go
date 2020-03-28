package controllers

type AuthController struct {
	BaseController
}

// @router / [get]
func (c *AuthController) Get() {
	c.Layout = "layout.html"
	c.TplName = "auth.html"
}

func (c *AuthController) GetAuths() {

}

func (c *AuthController) GetAuth() {

}

func (c *AuthController) AddAuth() {

}

func (c *AuthController) UpdAuth() {

}
