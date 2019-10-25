package controllers

import (
	"encoding/json"
	"zyra_core/models"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

type BuildRecordController struct {
	beego.Controller
}

// @Title CreateBuild
// @Description create Build
// @Param	body		body 	models.Build	true		"body for Build content"
// @Success 200 {int} models.Build
// @Failure 403 body is empty
// @router / [post]
func (b *BuildRecordController) Post() {
	logs.Info("请求参数为: ", string(b.Ctx.Input.RequestBody))
	var record models.BuildRecord
	json.Unmarshal(b.Ctx.Input.RequestBody, &record)
	objId := models.AddBuildRecord(record)
	b.Data["json"] = objId
	b.ServeJSON()
}

type ResponseBuildList struct {
	Data     []models.BuildRecord
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
func (u *BuildRecordController) GetAll() {
	// 获取参数
	pageSize, _ := u.GetInt64("pageSize")
	pageNum, _ := u.GetInt64("pageNum")
	sort := u.GetString("sort")
	desc, err := u.GetBool("desc")
	if err != nil {
		logs.Error("请求参数为desc错误,当前值为:", desc)
	}
	logs.Info("请求参数为 pageNum, pageSize, sort, desc: ", pageNum, pageSize, sort, desc)

	// 分页查询
	buildRecordObjs, count := models.GetAllBuildRecordByFilterCondition(pageNum, pageSize, sort, desc)

	//拼凑返回数据
	eesponseBuildList := &ResponseBuildList{buildRecordObjs, count, pageNum, pageSize}
	u.Data["json"] = eesponseBuildList
	u.ServeJSON()
}

// @Title Get
// @Description get Build by bid
// @Param	bid		path 	int		true		"The key for staticblock"
// @Success 200 {object} models.Build
// @Failure 403 :bid is empty
// @router /:bid [get]
func (b *BuildRecordController) Get() {
	bid, _ := b.GetInt(":bid")
	obj, _ := models.GetBuildRecordById(bid)
	b.Data["json"] = obj
	b.ServeJSON()
}

// @Title Update
// @Description update the Build
// @Param	bid		path 	int 	true		"The bid you want to update"
// @Param	body		body 	models.Build	true		"body for Build content"
// @Success 200 {object} models.Build
// @Failure 403 :bid is not int
// @router /:bid [put]
func (u *BuildRecordController) Put() {
	logs.Info("请求参数为: ", string(u.Ctx.Input.RequestBody))
	bid, _ := u.GetInt(":bid")
	var publish models.BuildRecord
	json.Unmarshal(u.Ctx.Input.RequestBody, &publish)
	obj := models.UpdateBuildRecord(bid, &publish)
	u.Data["json"] = obj
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	bid		path 	int 	true		"The bid you want to delete"
// @Success 200 {object} models.Build
// @Failure 403 bid is empty
// @router /:bid [delete]
func (b *BuildRecordController) Delete() {
	bid, _ := b.GetInt(":bid")
	obj := models.DeleteBuildRecord(bid)
	b.Data["json"] = obj
	b.ServeJSON()
}
