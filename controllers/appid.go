package controllers

import (
	"encoding/json"
	"zyra_core/models"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

type AppidController struct {
	beego.Controller
}

// @Title CreateBuild
// @Description create Build
// @Param	body		body 	models.Build	true		"body for Build content"
// @Success 200 {int} models.Build
// @Failure 403 body is empty
// @router / [post]
func (b *AppidController) Post() {
	logs.Info("请求参数为: ", string(b.Ctx.Input.RequestBody))
	var record models.Appid
	json.Unmarshal(b.Ctx.Input.RequestBody, &record)
	objId := models.AddAppid(record)
	b.Data["json"] = objId
	b.ServeJSON()
}

type ResponseAppidList struct {
	Data     []models.Appid
	Total    int64
	PageNum  int64
	PageSize int64
}

// @Title GetAll
// @Description get all Build
// @Success 200 {object} models.Build
// @Param   pageSize   	query   	string  false       "pageSize"
// @Param   pageNum 	query   	string  false       "pageNum"
// @Param   sort    	query		string  false       "sort"
// @router / [get]
func (u *AppidController) GetAll() {
	// 获取参数
	pageSize, _ := u.GetInt64("pageSize")
	pageNum, _ := u.GetInt64("pageNum")
	sort := u.GetString("sort")
	desc, err := u.GetBool("desc")
	if err != nil {
		logs.Error("请求参数为desc错误, 转为默认值为:", desc)
	}
	logs.Info("请求参数为 pageNum, pageSize, sort, desc: ", pageNum, pageSize, sort, desc)

	// 分页查询
	appidObjs, count := models.GetAllAppidByFilterCondition(pageNum, pageSize, sort, desc)

	//拼凑返回数据
	responseAppList := &ResponseAppidList{appidObjs, count, pageNum, pageSize}
	u.Data["json"] = responseAppList
	u.ServeJSON()
}

// @Title Get
// @Description get Build by bid
// @Param	bid		path 	int		true		"The key for staticblock"
// @Success 200 {object} models.Build
// @Failure 403 :bid is empty
// @router /:bid [get]
func (b *AppidController) Get() {
	bid, _ := b.GetInt(":bid")
	var appid models.Appid
	appid = appid.ReadOneById(int64(bid))
	b.Data["json"] = appid
	b.ServeJSON()
}

// @Title Update
// @Description update the Build
// @Param	bid		path 	int 	true		"The bid you want to update"
// @Param	body		body 	models.Build	true		"body for Build content"
// @Success 200 {object} models.Build
// @Failure 403 :bid is not int
// @router /:bid [put]
func (u *AppidController) Put() {
	logs.Info("请求参数为: ", string(u.Ctx.Input.RequestBody))
	bid, _ := u.GetInt64(":bid")
	var publish models.Appid
	json.Unmarshal(u.Ctx.Input.RequestBody, &publish)
	obj := models.UpdateAppid(bid, &publish)
	u.Data["json"] = obj
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	bid		path 	int 	true		"The bid you want to delete"
// @Success 200 {object} models.Build
// @Failure 403 bid is empty
// @router /:bid [delete]
func (b *AppidController) Delete() {
	bid, _ := b.GetInt64(":bid")
	obj := models.DeleteAppid(bid)
	b.Data["json"] = obj
	b.ServeJSON()
}
