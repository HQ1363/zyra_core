package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego/orm"
)

// Appid 发布表
type Appid struct {
	Id          int64     `orm:"column(id);auto;pk"`
	Appid       string    `orm:"column(appid)"`
	Details     string    `orm:"column(details)"`
	RepoUrl     string    `orm:"column(repo_url)"`
	BindDomain  string    `orm:"column(bind_domain)"`
	BindPort    int64     `orm:"column(bind_port)"`
	CreateTime  time.Time `orm:"column(create_time)"`
	UpdateTime  time.Time `orm:"column(update_time)"`
	ValidStatus int8      `orm:"column(valid_status)"`
	IsDelete    int8      `orm:"column(is_delete)"`
}

// GetAllAppidByFilterCondition 根据条件分页查询
func GetAllAppidByFilterCondition(pageNum, pageSize int64, sort string, desc bool) ([]Appid, int64) {
	o := orm.NewOrm()
	o.Using("default")

	// test raws
	var offset int64
	if pageNum <= 1 {
		offset = 0
	} else {
		offset = (pageNum - 1) * pageSize
	}

	logs.Info("SELECT id, appid, details, repo_url, bind_domain, bind_port, create_time, update_time, valid_status, is_delete FROM appid limit %d, %d \n", offset, pageSize)

	var appids []Appid
	o.Raw("SELECT id, appid, details, repo_url, bind_domain, bind_port, create_time, update_time, valid_status, is_delete FROM appid limit ?, ?", offset, pageSize).QueryRows(&appids)

	var count int64

	o.Raw("SELECT count(1) FROM appid").QueryRow(&count)
	return appids, count
}

func (sl *Appid) ReadOneById(Id int64) Appid {
	o := orm.NewOrm()
	o.Using("default")

	var appid Appid
	err := o.QueryTable("Appid").Filter("Id", Id).One(&appid)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		logs.Error("应该找到单条数据, 但是找到多条!")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		logs.Error("没有任何记录!")
	}

	if err == nil {
		return appid
	}

	return appid
}

// GetAppidById 通过Id查询记录
func GetAppidById(id int) (bObj Appid, err error) {
	o := orm.NewOrm()
	o.Using("default")
	err = o.QueryTable("Appid").Filter("Id", id).One(&bObj)
	logs.Info(bObj)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		logs.Error("应该找到单条数据, 但是找到多条!")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		logs.Error("没有任何记录!")
	}

	if err == nil {
		return bObj, nil
	}
	return Appid{}, nil

}

// AddAppid 增加新纪录
func AddAppid(obj Appid) int64 {
	o := orm.NewOrm()
	o.Using("default")
	logs.Info("增加记录为: ", obj)
	bid, err := o.Insert(&obj)
	if err == nil {
		return bid
	}

	return 0

}

// UpdateAppid 修改某个字段
func UpdateAppid(bid int64, obj *Appid) Appid {
	o := orm.NewOrm()
	bObj := Appid{Id: bid}
	log.Println("修改记录id为: ", bid)
	if o.Read(&bObj) == nil {
		if obj.Appid != "" {
			bObj.Appid = obj.Appid
		}

		if obj.Details != "" {
			bObj.Details = obj.Details
		}

		if obj.BindDomain != "" {
			bObj.BindDomain = obj.BindDomain
		}

		if obj.BindPort != 0 {
			bObj.BindPort = obj.BindPort
		}

		if obj.IsDelete != 0 {
			bObj.IsDelete = obj.IsDelete
		}

		if obj.ValidStatus != 0 {
			bObj.ValidStatus = obj.ValidStatus
		}

		// *new(time.Time) 时间的默认值
		if obj.CreateTime != *new(time.Time) {
			bObj.CreateTime = obj.CreateTime
		}

		if obj.UpdateTime != *new(time.Time) {
			bObj.UpdateTime = obj.UpdateTime
		}

		if _, err := o.Update(&bObj); err == nil {
			logs.Info("记录id: ", bObj.Id, "已经更新!")
		}
	}
	return bObj
}

// DeleteAppid 软删除1条记录
func DeleteAppid(id int64) Appid {
	o := orm.NewOrm()
	o.Using("default")
	obj := Appid{Id: id}

	if o.Read(&obj) == nil {
		logs.Info("删除id为: ", obj.Id)
		num, err := o.QueryTable("Appid").Filter("Id", id).Update(orm.Params{
			"IsDelete": 1,
		})
		obj.IsDelete = 1
		if err != nil {
			logs.Info(num)
		}
	}
	return obj
}

// init 初始化固定
func init() {
	orm.RegisterModel(new(Appid))
}
