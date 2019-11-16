package models

import (
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego/orm"
)

// BuildRecord 发布表
type BuildRecord struct {
	Id              int       `orm:"column(id);auto;pk"`
	Author          string    `orm:"column(author)"`
	Appid           string    `orm:"column(appid)"`
	EventName       string    `orm:"column(event_name)"`
	CommitUrl       string    `orm:"column(commit_url)"`
	CommitHash      string    `orm:"column(commit_hash)"`
	CommitMessage   string    `orm:"column(commit_message)"`
	CommitBranch    string    `orm:"column(commit_branch)"`
	BuildAddr       string    `orm:"column(build_addr)"`
	BuildStatus     string    `orm:"column(build_status)"`
	BuildFinishTime time.Time `orm:"column(build_finish_time)"`
	PackagePath     string    `orm:"column(package_path)"`
	CreateTime      time.Time `orm:"column(create_time)"`
	UpdateTime      time.Time `orm:"column(update_time)"`
	ValidStatus     int8      `orm:"column(valid_status)"`
	IsDelete        int8      `orm:"column(is_delete)"`
}

// GetAllBuildRecordByFilterCondition 根据条件分页查询
func GetAllBuildRecordByFilterCondition(pageNum, pageSize int64, sort string, desc bool) ([]BuildRecord, int64) {
	o := orm.NewOrm()

	// test raw
	var buildRecord BuildRecord
	_ = o.Raw("SELECT id, appid FROM build_record WHERE id = ?", 1412).QueryRow(&buildRecord)
	logs.Info(buildRecord.Id, buildRecord.Appid, buildRecord.EventName)

	// test raws
	var offset int64
	if pageNum <= 1 {
		offset = 0
	} else {
		offset = (pageNum - 1) * pageSize
	}
	var buildRecords []BuildRecord
	_, _ = o.Raw("SELECT id, appid FROM build_record limit ?, ?", pageSize, offset).QueryRows(&buildRecords)
	fmt.Print(buildRecords)

	var count int64

	_ = o.Raw("SELECT count(1) FROM build_record", pageSize, offset).QueryRow(&count)
	return buildRecords, count
}

// GetBuildRecordById 通过Id查询记录
func GetBuildRecordById(id int) (bObj BuildRecord, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("BuildRecord").Filter("Id", id).One(&bObj)
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
	return BuildRecord{}, nil

}

// AddBuildRecord 增加新纪录
func AddBuildRecord(obj BuildRecord) int64 {
	o := orm.NewOrm()
	logs.Info("增加记录为: ", obj)
	bid, err := o.Insert(&obj)
	if err == nil {
		return bid
	}

	return 0

}

// UpdateBuildRecord 修改某个字段
func UpdateBuildRecord(bid int, obj *BuildRecord) BuildRecord {
	o := orm.NewOrm()
	bObj := BuildRecord{Id: bid}
	log.Println("修改记录id为: ", bid)
	if o.Read(&bObj) == nil {
		if obj.Author != "" {
			bObj.Author = obj.Author
		}

		if obj.Appid != "" {
			bObj.Appid = obj.Appid
		}

		if obj.EventName != "" {
			bObj.EventName = obj.EventName
		}

		if obj.EventName != "" {
			bObj.EventName = obj.EventName
		}

		if obj.CommitUrl != "" {
			bObj.CommitUrl = obj.CommitUrl
		}

		if obj.CommitHash != "" {
			bObj.CommitHash = obj.CommitHash
		}

		if obj.CommitMessage != "" {
			bObj.CommitMessage = obj.CommitMessage
		}

		if obj.CommitBranch != "" {
			bObj.CommitBranch = obj.CommitBranch
		}

		if obj.BuildAddr != "" {
			bObj.BuildAddr = obj.BuildAddr
		}

		if obj.BuildStatus != "" {
			bObj.BuildStatus = obj.BuildStatus
		}

		if obj.BuildFinishTime != *new(time.Time) {
			bObj.BuildFinishTime = obj.BuildFinishTime
		}

		if obj.PackagePath != "" {
			bObj.PackagePath = obj.PackagePath
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

// DeleteBuildRecord 软删除1条记录
func DeleteBuildRecord(id int) BuildRecord {
	o := orm.NewOrm()
	obj := BuildRecord{Id: id}

	if o.Read(&obj) == nil {
		logs.Info("删除id为: ", obj.Id)
		num, err := o.QueryTable("BuildRecord").Filter("Id", id).Update(orm.Params{
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
	orm.RegisterModel(new(BuildRecord))
}
