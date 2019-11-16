package controllers

import (
	"encoding/json"
	"zyra_core/models"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)

type PublishRecordController struct {
	beego.Controller
}

// @Title CreatePublishRecord
// @Description create PublishRecord
// @Param	body		body 	models.PublishRecord	true		"body for PublishRecord content"
// @Success 200 {int} models.PublishRecord
// @Failure 403 body is empty
// @router / [post]
func (b *PublishRecordController) Post() {
	logs.Info("请求参数为: ", string(b.Ctx.Input.RequestBody))
	var record models.PublishRecord
	_ = json.Unmarshal(b.Ctx.Input.RequestBody, &record)
	objId := models.AddPublishRecord(record)
	b.Data["json"] = objId
	b.ServeJSON()
}

type PublishRecordList struct {
	Data     []orm.Params
	Total    int64
	PageNum  int64
	PageSize int64
}

// @Title GetAll
// @Description get all PublishRecord
// @Success 200 {object} models.PublishRecord
// @Param   pageSize   	query   	string  false       "pageSize"
// @Param   pageNum 	query   	string  false       "pageNum"
// @Param   sort    	query		string  false       "sort"
// @router / [get]
func (u *PublishRecordController) GetAll() {
	// 获取参数
	pageSize, _ := u.GetInt64("pageSize")
	pageNum, _ := u.GetInt64("pageNum")
	sort := u.GetString("sort")
	logs.Info("请求参数为 pageNum, pageSize, sort: ", pageNum, pageSize, sort)

	// 分页查询
	publishObjs, count := models.GetAllPublishRecordByFilterCondition(pageNum, pageSize, sort)

	//拼凑返回数据
	publishRecordStruct := &PublishRecordList{publishObjs, count, pageNum, pageSize}
	u.Data["json"] = publishRecordStruct
	u.ServeJSON()
}

// @Title Get
// @Description get PublishRecord by bid
// @Param	bid		path 	int		true		"The key for staticblock"
// @Success 200 {object} models.PublishRecord
// @Failure 403 :bid is empty
// @router /:bid [get]
func (b *PublishRecordController) Get() {
	bid, _ := b.GetInt(":bid")
	obj, _ := models.GetPublishRecordById(bid)
	b.Data["json"] = obj
	b.ServeJSON()
}

// @Title Update
// @Description update the PublishRecord
// @Param	bid		path 	int 	true		"The bid you want to update"
// @Param	body		body 	models.PublishRecord	true		"body for PublishRecord content"
// @Success 200 {object} models.PublishRecord
// @Failure 403 :bid is not int
// @router /:bid [put]
func (u *PublishRecordController) Put() {
	logs.Info("请求参数为: ", string(u.Ctx.Input.RequestBody))
	bid, _ := u.GetInt(":bid")
	var publish models.PublishRecord
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &publish)
	obj := models.UpdatePublishRecord(bid, &publish)
	u.Data["json"] = obj
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	bid		path 	int 	true		"The bid you want to delete"
// @Success 200 {object} models.PublishRecord
// @Failure 403 bid is empty
// @router /:bid [delete]
func (b *PublishRecordController) Delete() {
	bid, _ := b.GetInt(":bid")
	obj := models.DeletePublishRecord(bid)
	b.Data["json"] = obj
	b.ServeJSON()
}
