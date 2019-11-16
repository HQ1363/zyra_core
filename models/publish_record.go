package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego/orm"
)

// PublishRecord 发布表
type PublishRecord struct {
	Id          int       `orm:"column(id);auto;pk"`
	Appid       string    `orm:"column(appid)"`
	OpsType     int       `orm:"column(ops_type)"`
	OpsStatus   int       `orm:"column(ops_status)"`
	Env         string    `orm:"column(env)"`
	UpsList     string    `orm:"column(ups_list)"`
	PackageURL  string    `orm:"column(package_url)"`
	PublishLog  string    `orm:"column(publish_log)"`
	CreateTime  time.Time `orm:"column(create_time)"`
	UpdateTime  time.Time `orm:"column(update_time)"`
	ValidStatus int8      `orm:"column(valid_status)"`
	IsDelete    int8      `orm:"column(is_delete)"`
}

// GetAllByFilterCondition 根据条件分页查询
func GetAllPublishRecordByFilterCondition(pageNum, pageSize int64, sort string) (records []orm.Params, count int64) {
	o := orm.NewOrm()
	qs := o.QueryTable("PublishRecord")
	var offset int64
	if pageNum <= 1 {
		offset = 0
	} else {
		offset = (pageNum - 1) * pageSize
	}
	_, _ = qs.Limit(pageSize, offset).OrderBy(sort).Values(&records, "Id", "Appid", "OpsType", "OpsStatus", "Env", "UpsList", "PackageURL", "CreateTime", "UpdateTime")
	count, _ = qs.Count()
	return records, count
}

// GetById 通过Id查询记录
func GetPublishRecordById(id int) (bObj PublishRecord, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("PublishRecord").Filter("Id", id).One(&bObj)
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
	return PublishRecord{}, nil

}

// Add 增加新纪录
func AddPublishRecord(obj PublishRecord) int64 {
	o := orm.NewOrm()
	logs.Info("增加记录为: ", obj)
	bid, err := o.Insert(&obj)
	if err == nil {
		return bid
	}

	return 0

}

// Update 修改某个字段
func UpdatePublishRecord(bid int, obj *PublishRecord) PublishRecord {
	o := orm.NewOrm()
	bObj := PublishRecord{Id: bid}
	log.Println("修改记录id为: ", bid)
	if o.Read(&bObj) == nil {
		if obj.Appid != "" {
			bObj.Appid = obj.Appid
		}

		if obj.OpsType != 0 {
			bObj.Appid = obj.Appid
		}

		if obj.OpsStatus != 0 {
			bObj.OpsStatus = obj.OpsStatus
		}

		if obj.Env != "" {
			bObj.Env = obj.Env
		}

		if obj.UpsList != "" {
			bObj.UpsList = obj.UpsList
		}

		if obj.PackageURL != "" {
			bObj.PackageURL = obj.PackageURL
		}

		if obj.PublishLog != "" {
			bObj.PublishLog = obj.PublishLog
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

// Delete 软删除1条记录
func DeletePublishRecord(id int) PublishRecord {
	o := orm.NewOrm()
	obj := PublishRecord{Id: id}

	if o.Read(&obj) == nil {
		logs.Info("删除id为: ", obj.Id)
		num, err := o.QueryTable("publishRecord").Filter("Id", id).Update(orm.Params{
			"IsDelete": 1,
		})
		if err != nil {
			logs.Info(num)
		}
	}
	return obj
}

// init 初始化固定
func init() {
	orm.RegisterModel(new(PublishRecord))
}
