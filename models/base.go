package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"reflect"
	"time"
)

var (
	logger = logs.GetLogger("model")
	dbUser = beego.AppConfig.String("mysqlUser")
	dbPass = beego.AppConfig.String("mysqlPass")
	dbHost = beego.AppConfig.String("mysqlHost")
	dbPort = beego.AppConfig.String("mysqlPort")
	dbName = beego.AppConfig.String("mysqlDb")
)

func InitDB() {
	// 创建数据库
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", dbUser, dbPass, dbHost, dbPort)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Println("数据源连接失败：", err)
		os.Exit(2)
	}
	var sqlString = fmt.Sprintf(" CREATE DATABASE if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", dbName)
	r, err := db.Exec(sqlString)
	if err != nil {
		logger.Println(err, r)
		_ = db.Close()
	} else {
		_ = db.Close()
		logger.Println("执行数据库" + dbName + "创建语句成功")
	}
}

func InitConn() {
	// 连接数据库
	maxIdleConn, _ := beego.AppConfig.Int("mysql_max_idle_conn")
	maxOpenConn, _ := beego.AppConfig.Int("mysql_max_open_conn")
	var dbLink = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName) + "&loc=Asia%2FShanghai"
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		logger.Println("数据库连接错误：", err)
		os.Exit(2)
		return
	}
	err = orm.RegisterDataBase("default", "mysql", dbLink, maxIdleConn, maxOpenConn)
	orm.Debug = true
	if err != nil {
		logger.Println("数据库连接错误:", err)
		os.Exit(2)
		return
	} else {
		err := orm.RunSyncdb("default", false, true)
		if err != nil {
			logger.Println("数据库自动建表失败: ", err)
		}
	}
}

/*
  1、自动创建表/model的函数
  2、常用的model方法
*/
type BaseModel struct {
	// model代表具体的对象
	Id        int64       `orm:"pk;auto;column(id)"`
	UpdatedAt time.Time   `orm:"auto_now;type(datetime);column(updated_at)" description:"更新时间"`
	CreatedAt time.Time   `orm:"auto_now_add;type(datetime);column(created_at)" description:"创建时间"`
	IsValid   bool        `orm:"default(1);column(is_valid)" description:"逻辑删除字段"`
}

type BaseModelInterface interface {
	Count(tableName string) int64
	Exists(tableName string, field string, value string) bool
	QueryAll(tableName string, orderBy string, desc bool, withExpr bool, filters map[string]interface{}, records []*interface{}, operators ...string) (int64, error)
	PaginationQuery(tableName string, orderBy string, desc bool, pageNum int, pageSize int, filters ...interface{}) (interface{}, int64)
}

func (model *BaseModel) Count(tableName string) int64 {
	o := orm.NewOrm()
	cnt, _ := o.QueryTable(tableName).Count()
	return cnt
}

func (model *BaseModel) Exists(tableName string, field string, value string) bool {
	o := orm.NewOrm()
	return o.QueryTable(tableName).Filter(field, value).Exist()
}

func (model *BaseModel) QueryAll(tableName string, orderBy string, desc bool, withExpr bool, filters map[string]interface{}, records interface{}, operators ...string) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(tableName)
	index := 0
	for key, value := range filters {
		if !withExpr {
			qs = qs.Filter(key, value)
		} else {
			var expr = key + "__" + operators[index]
			qs = qs.Filter(expr, value)
			index ++
		}
	}
	if orderBy != "" {
		if desc {
			qs = qs.OrderBy("-" + orderBy)
		} else {
			qs = qs.OrderBy(orderBy)
		}
	}
	fmt.Println("type is: ", reflect.TypeOf(records))
	return qs.All(records)
}

func (model *BaseModel) PaginationQuery(tableName string, orderBy string, desc bool, pageNum int, pageSize int, filters ...interface{}) (interface{}, int64) {
	offset := (pageNum - 1) * pageSize
	list := make([]*User, 0)
	query := orm.NewOrm().QueryTable(tableName)
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	if orderBy != "" {
		orderBy = "id"
	}
	if desc {
		orderBy = "-" + orderBy
	} else {
		orderBy = "+" + orderBy
	}
	_, _ = query.OrderBy(orderBy).Limit(pageSize, offset).All(&list)
	return list, total
}

func (model *BaseModel) GetQuerySet(tableName string) orm.QuerySeter {
	o := orm.NewOrm()
	return o.QueryTable(tableName)
}

func (model *BaseModel) ExecRawSql(sqlFmt string, values ...string) orm.RawSeter {
	// 执行裸sql
	o := orm.NewOrm()
	return o.Raw(sqlFmt, values)
}
