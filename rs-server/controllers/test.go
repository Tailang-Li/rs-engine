package controllers

import (
	"github.com/astaxie/beego"
	"rs-engine/rs-server/models"
)

type TestModelController struct {
	beego.Controller
}

func (c *TestModelController) Save() {
	model := new(models.TestModel)
	c.Ctx.WriteString(model.SaveTest())
}

func (c *TestModelController) Update() {
	model := new(models.TestModel)
	c.Ctx.WriteString(model.UpdateTest())
}

func (c *TestModelController) Read() {
	model := new(models.TestModel)
	c.Ctx.WriteString(model.ReadTest())
}

func (c *TestModelController) Delete() {
	model := new(models.TestModel)
	c.Ctx.WriteString(model.DeleteTest())
}
