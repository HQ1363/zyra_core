package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego/orm"
)

// Host 发布表
type Host struct {
	Id          int64     `orm:"column(id);auto;pk"`
	Hostname    string    `orm:"column(host_name)"`
	IpAddress   string    `orm:"column(ip_address)"`
	Env         string    `orm:"column(env)"`
	Type        string    `orm:"column(type)"`
	BindPort    int64     `orm:"column(bind_port)"`
	CreateTime  time.Time `orm:"column(create_time)"`
	UpdateTime  time.Time `orm:"column(update_time)"`
	ValidStatus int8      `orm:"column(valid_status)"`
	IsDelete    int8      `orm:"column(is_delete)"`
}

// GetAllHostByFilterCondition 根据条件分页查询
func GetAllHostByFilterCondition(pageNum, pageSize int64, sort string, desc bool) ([]Host, int64) {
	o := orm.NewOrm()

	// test raw
	var host Host
	o.Raw("SELECT id, host_name, ip_address, env, type, bind_port FROM host WHERE id = ?", 1412).QueryRow(&host)

	// test raws
	var offset int64
	if pageNum <= 1 {
		offset = 0
	} else {
		offset = (pageNum - 1) * pageSize
	}
	var hosts []Host
	o.Raw("SELECT id, host_name, ip_address, env, type, bind_port FROM host limit ?, ?", pageSize, offset).QueryRows(&hosts)

	var count int64

	o.Raw("SELECT count(1) FROM hosts", pageSize, offset).QueryRow(&count)
	return hosts, count
}

// GetHostById 通过Id查询记录
func GetHostById(id int) (bObj Host, err error) {
	o := orm.NewOrm()
	o.Using("default")
	err = o.QueryTable("host").Filter("Id", id).One(&bObj)
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
	return Host{}, nil

}

// AddHost 增加新纪录
func AddHost(obj Host) int64 {
	o := orm.NewOrm()
	o.Using("default")
	logs.Info("增加记录为: ", obj)
	bid, err := o.Insert(&obj)
	if err == nil {
		return bid
	}

	return 0

}

// UpdateHost 修改某个字段
func UpdateHost(bid int64, obj *Host) Host {
	o := orm.NewOrm()
	bObj := Host{Id: bid}
	log.Println("修改记录id为: ", bid)
	if o.Read(&bObj) == nil {
		if obj.Hostname != "" {
			bObj.Hostname = obj.Hostname
		}

		if obj.IpAddress != "" {
			bObj.IpAddress = obj.IpAddress
		}

		if obj.Env != "" {
			bObj.Env = obj.Env
		}

		if obj.Type != "" {
			bObj.Type = obj.Type
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

// DeleteHost 软删除1条记录
func DeleteHost(id int64) Host {
	o := orm.NewOrm()
	o.Using("default")
	obj := Host{Id: id}

	if o.Read(&obj) == nil {
		logs.Info("删除id为: ", obj.Id)
		num, err := o.QueryTable("host").Filter("Id", id).Update(orm.Params{
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
	orm.RegisterModel(new(Host))
}
