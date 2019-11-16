package controllers

import (
	"zyra_core/models"
	"encoding/json"
	"fmt"
	"strconv"
)

// Operations about Users
type UserController struct {
	BaseController
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	// 显示接收过来的body参数
	fmt.Println(string(u.Ctx.Input.RequestBody))
	var user models.User
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		u.jsonResult(1, "数据解析失败：" + err.Error(), nil)
	}
	_, err = user.Save()
	if err != nil {
		u.jsonResult(1, "插入数据库失败：" + err.Error(), nil)
	}
	u.jsonResult(0, "", user)
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	var user models.User
	users, _ := user.PaginationQuery(user.TableName(),"id", true, 1, 1000)
	u.jsonResult(0, "success", users)
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid, err := strconv.Atoi(u.GetString(":uid"))
	if err == nil {
		user, err := models.GetUserById(uid)
		if err != nil {
			u.jsonResult(1, err.Error(), nil)
		} else {
			u.jsonResult(0, "success", user)
		}
	} else {
		u.jsonResult(1, err.Error(), nil)
	}
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	fmt.Println("uid: ", uid)
	if uid != "" {
		var user models.User
		err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		if err != nil {
			u.jsonResult(1, "参数解析异常: " + err.Error() , nil)
		}
		user.Update(true)
		u.jsonResult(0, "success", user)
	}
	u.jsonResult(1, "用户不存在", nil)
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	var user models.User
	uid, err := strconv.Atoi(u.GetString(":uid"))
	user.Id = int64(uid)
	if err == nil {
		models.DeleteUserById(uid, user.SSOId)
	}
	u.jsonResult(0, "delete success", nil)
}
