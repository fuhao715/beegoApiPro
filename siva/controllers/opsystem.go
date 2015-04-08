package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"siva/models"
	"strconv"
)

// Operations about OpSystem
type OpSystemController struct {
	beego.Controller
}

// @Title createOpSystem
// @Description create opSystem
// @Param	body	body	models.OpSystem	true "body for user content"
// @Success 200 {int} models.OpSystem.Id
// @Failure 403 body is empty
// @router / [post]
func (p *OpSystemController) Post() {
	var opSystem models.OpSystem
	json.Unmarshal(p.Ctx.Input.RequestBody, &opSystem)
	pid := models.AddOpSystem(opSystem)
	p.Data["json"] = map[string]int{"id": pid}
	p.ServeJson()
}

// @Title Get
// @Description get opSystem by pid
// @Param	pid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.OpSystem
// @Failure 403 :pid is empty
// @router /:pid [get]
func (this *OpSystemController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":pid"))
	//  pid := 1 p.GetInt("pid")

	opSystem, err := models.GetOpSystem(id)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = opSystem
	}
	this.Data["json"] = opSystem

	this.ServeJson()
}
