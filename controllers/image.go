package controllers

import (
	"github.com/astaxie/beego"
)

type ImageController struct {
	beego.Controller
}

func (c *ImageController) Get() {
	c.Ctx.WriteString("hello starchou")
}
