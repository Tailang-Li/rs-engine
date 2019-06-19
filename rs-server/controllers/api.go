package controllers

import "github.com/astaxie/beego"

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Version() {
	c.Ctx.WriteString("dev-0.0.1")
}
