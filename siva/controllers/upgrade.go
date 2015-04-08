package controllers

import (
	// "encoding/json"
	"github.com/astaxie/beego"
	"siva/models"
	// "strconv"
)

// Operations about Upgrade
type UpgradeController struct {
	beego.Controller
}

// @Title Get
// @Description get Upgrade by pid
// @Param	pid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Upgrade
// @Failure 403 :pid is empty
// @router /:pid [get]
func (this *UpgradeController) Get() {
	id := this.Ctx.Input.Param(":pid") // strconv.Atoi()
	//  pid := 1 p.GetInt("pid")

	opSystem, err := models.GetUpgrade(id)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = opSystem
	}
	this.Data["json"] = opSystem

	this.ServeJson()
}
