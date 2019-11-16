package models

import (
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego/orm"
)

// Hostgroup 发布表
type Hostgroup struct {
	Id          int64     `orm:"column(id);auto;pk"`
	Name        string    `orm:"column(name)"`
	Env         string    `orm:"column(env)"`
	Useage      string    `orm:"column(useage)"`
	CreateTime  time.Time `orm:"column(create_time)"`
	UpdateTime  time.Time `orm:"column(update_time)"`
	ValidStatus int8      `orm:"column(valid_status)"`
	IsDelete    int8      `orm:"column(is_delete)"`
}

// GetAllHostgroupByFilterCondition 根据条件分页查询
func GetAllHostgroupByFilterCondition(pageNum, pageSize int64, sort string, desc bool) ([]Hostgroup, int64) {
	o := orm.NewOrm()

	// test raw
	var hostgroup Hostgroup
	_ = o.Raw("SELECT id, appid FROM hostgroup WHERE id = ?", 1412).QueryRow(&hostgroup)

	// test raws
	var offset int64
	if pageNum <= 1 {
		offset = 0
	} else {
		offset = (pageNum - 1) * pageSize
	}
	var hostgroups []Hostgroup
	_, _ = o.Raw("SELECT id, host_id, hostgroup_id FROM hostgroup limit ?, ?", pageSize, offset).QueryRows(&hostgroups)

	fmt.Print(hostgroups)

	var count int64

	_ = o.Raw("SELECT count(1) FROM hostgroup", pageSize, offset).QueryRow(&count)
	return hostgroups, count
}

// GetHostgroupById 通过Id查询记录
func GetHostgroupById(id int) (bObj Hostgroup, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("hostgroup").Filter("Id", id).One(&bObj)
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
	return Hostgroup{}, nil

}

// AddHostgroup 增加新纪录
func AddHostgroup(obj Hostgroup) int64 {
	o := orm.NewOrm()
	logs.Info("增加记录为: ", obj)
	bid, err := o.Insert(&obj)
	if err == nil {
		return bid
	}

	return 0

}

// UpdateHostgroup 修改某个字段
func UpdateHostgroup(bid int64, obj *Hostgroup) Hostgroup {
	o := orm.NewOrm()
	bObj := Hostgroup{Id: bid}
	log.Println("修改记录id为: ", bid)
	if o.Read(&bObj) == nil {
		if obj.Name != "" {
			bObj.Name = obj.Name
		}

		if obj.Env != "" {
			bObj.Env = obj.Env
		}

		if obj.Useage != "" {
			bObj.Useage = obj.Useage
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

// DeleteHostgroup 软删除1条记录
func DeleteHostgroup(id int64) Hostgroup {
	o := orm.NewOrm()
	obj := Hostgroup{Id: id}

	if o.Read(&obj) == nil {
		logs.Info("删除id为: ", obj.Id)
		num, err := o.QueryTable("hostgroup").Filter("Id", id).Update(orm.Params{
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
	orm.RegisterModel(new(Hostgroup))
}
