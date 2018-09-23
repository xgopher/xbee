package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/xgopher/xbee/modules/common/models"
)

type ObjectController struct {
	beego.Controller
}

// 创建
func (o *ObjectController) Store() {
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"ObjectId": objectid}
	o.ServeJSON()
}

func (o *ObjectController) Show() {
	objectId := o.Ctx.Input.Param(":id")
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// 列表
func (o *ObjectController) Index() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// 更新
func (o *ObjectController) Update() {
	objectId := o.Ctx.Input.Param(":id")
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectId, ob.Score)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// 删除
func (o *ObjectController) Destroy() {
	objectId := o.Ctx.Input.Param(":id")
	models.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}
