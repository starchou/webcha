package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/starchou/webcha/utils"
)

type ImageController struct {
	beego.Controller
}

func (c *ImageController) Get() {
	/*
		data, err := utils.GetMap("39.983424,116.322987")
		if err != nil {
			c.Ctx.WriteString(err.Error())
		}
		temp := ""
		for _, v := range data.Result.Pois {
			temp += v.Addr + "|"
			temp += v.Cp + "|"
			temp += v.Distance + "|"
			temp += v.Name + "|"
			temp += v.PoiType + "|"
			temp += v.Tel + "|"
			temp += v.Uid + "|"
			temp += "<br />"
		}
	*/
	c.Ctx.WriteString("hello starchou")
}
