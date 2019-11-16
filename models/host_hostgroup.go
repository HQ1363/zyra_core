package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego/orm"
)

// HostHostgroup 发布表
type HostHostgroup struct {
	Id           int64     `orm:"column(id);auto;pk"`
	HostId       string    `orm:"column(host_id)"`
	Hostgroup_Id string    `orm:"column(hostgroup_id)"`
	CreateTime   time.Time `orm:"column(create_time)"`
	UpdateTime   time.Time `orm:"column(update_time)"`
	ValidStatus  int8      `orm:"column(valid_status)"`
	IsDelete     int8      `orm:"column(is_delete)"`
}

// GetAllHostHostgroupByFilterCondition 根据条件分页查询
func GetAllHostHostgroupByFilterCondition(pageNum, pageSize int64, sort string, desc bool) ([]HostHostgroup, int64) {
	o := orm.NewOrm()

	// test raw
	var hostHostgroup HostHostgroup
	_ = o.Raw("SELECT id, host_id,hostgroup_id FROM host_hostgroup WHERE id = ?", 1412).QueryRow(&hostHostgroup)

	// test raws
	var offset int64
	if pageNum <= 1 {
		offset = 0
	} else {
		offset = (pageNum - 1) * pageSize
	}
	var hostHostgroups []HostHostgroup
	_, _ = o.Raw("SELECT id, appid FROM host_hostgroup limit ?, ?", pageSize, offset).QueryRows(&hostHostgroups)

	var count int64

	_ = o.Raw("SELECT count(1) FROM host_hostgroup", pageSize, offset).QueryRow(&count)
	return hostHostgroups, count
}

// GetHostHostgroupById 通过Id查询记录
func GetHostHostgroupById(id int) (bObj HostHostgroup, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("host_hostgroup").Filter("Id", id).One(&bObj)
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
	return HostHostgroup{}, nil

}

// AddHostHostgroup 增加新纪录
func AddHostHostgroup(obj Appid) int64 {
	o := orm.NewOrm()
	logs.Info("增加记录为: ", obj)
	bid, err := o.Insert(&obj)
	if err == nil {
		return bid
	}

	return 0

}

// UpdateHostHostgroup 修改某个字段
func UpdateHostHostgroup(bid int64, obj *HostHostgroup) HostHostgroup {
	o := orm.NewOrm()
	bObj := HostHostgroup{Id: bid}
	log.Println("修改记录id为: ", bid)
	if o.Read(&bObj) == nil {
		if obj.HostId != "" {
			bObj.HostId = obj.HostId
		}

		if obj.Hostgroup_Id != "" {
			bObj.Hostgroup_Id = obj.Hostgroup_Id
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

// DeleteHostHostgroup 软删除1条记录
func DeleteHostHostgroup(id int64) HostHostgroup {
	o := orm.NewOrm()
	obj := HostHostgroup{Id: id}

	if o.Read(&obj) == nil {
		logs.Info("删除id为: ", obj.Id)
		num, err := o.QueryTable("host_hostgroup").Filter("Id", id).Update(orm.Params{
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
	orm.RegisterModel(new(HostHostgroup))
}
