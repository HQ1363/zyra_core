package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	BaseModel
	Username string `orm:"null;size(64);column(username);type(varchar)" description:"用户名"`
	SSOId string `orm:"index;unique;size(64);column(sso_id);type(varchar)" description:"员工工号"`
}

// 表名
func (u *User) TableName() string {
	return "user"
}

func init() {
	// 注册模型
	orm.RegisterModel(new(User))
}

// CRUD
func (u *User) Save() (int64, error) {
	return orm.NewOrm().Insert(u)
}

func InsertMultiUser (users []User) (int64, error) {
	o := orm.NewOrm()
	if successNum, err := o.InsertMulti(len(users), users); err == nil {
		if successNum > 0 {
			fmt.Printf("成功插入 %d 个用户.\n", successNum)
		} else {
			fmt.Println("批量插入数据失败")
		}
		return successNum, err
	}
	return 0, nil
}

func GetUserById(id int) (*User, error) {
	r := new(User)
	err := orm.NewOrm().QueryTable(r.TableName()).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func DeleteUserById(id int, ssoId string) {
	o := orm.NewOrm()
	userPtr := new(User)
	userPtr.Id = int64(id)
	if ssoId != "" {
		userPtr = &User { SSOId: ssoId }
	}
	if num, err := o.Delete(userPtr); err == nil {
		fmt.Printf("删除记录数: %d", num)
	}
}

func (u *User) Update(updateAll bool, fields ... string) {
	o := orm.NewOrm()
	_ = o.Begin()
	var readUser User
	readUser.Id = u.Id
	if o.Read(&readUser) == nil {
		if updateAll {
			if num, err := o.Update(u); err == nil {
				_ = o.Commit()
				fmt.Println("更新记录数: ", num)
			} else {
				_ = o.Rollback()
			}
		} else {
			if num, err := o.Update(u, fields ...); err == nil {
				_ = o.Commit()
				fmt.Println("更新记录数: ", num)
			} else {
				_ = o.Rollback()
			}
		}
	}
}

func GetOrCreateUser(user *User) {
	o := orm.NewOrm()
	if created, id, err := o.ReadOrCreate(user, "SSOId"); err == nil {
		if created {
			fmt.Println("New Insert an user object. Id:", id)
		} else {
			fmt.Println("Get an user object. Id:", id)
		}
	}
}

func (u *User) MgetUserByIds(ids string) ([]*User, error) {
	list := make([]*User, 0)
	sql := "SELECT * FROM " + u.TableName() + " WHERE id in(%s)"
	setter := u.ExecRawSql(sql, ids)
	if _, err := setter.QueryRows(&list); err != nil {
		logger.Println("批量查询DB失败：", err)
	}
	return list, nil
}
