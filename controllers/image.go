package controllers

import (
	"github.com/astaxie/beego"
	"github.com/starchou/webcha/utils"
)

type ImageController struct {
	beego.Controller
}

func (c *ImageController) Get() {
	data, err := utils.GetMap("39.983424,116.322987")
	if err != nil {
		c.Ctx.WriteString(err.Error())
	}
	c.Ctx.WriteString(data.Result.Formatted_address)
}
