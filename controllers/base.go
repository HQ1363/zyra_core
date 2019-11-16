package controllers

import (
	"github.com/astaxie/beego"
	"runtime"
	"strings"
	"zyra_core/enums"
	"zyra_core/models"
)

// JsonResult 用于返回ajax请求的基类
type JsonResult struct {
	Code enums.JsonResultCode `json:"code"`
	Msg  string               `json:"msg"`
	Data  interface{}          `json:"obj"`
}

// BaseQueryParam 用于查询的类
type BaseQueryParam struct {
	Sort   string `json:"sort"`
	Order  string `json:"order"`
	Offset int64  `json:"offset"`
	Limit  int    `json:"limit"`
}

type BaseController struct {
	beego.Controller
	controllerName string             //当前控制名称
	actionName     string             //当前action名称
	curUser        models.User //当前用户信息
}

// prepare方法由框架自动调用
func (c *BaseController) Prepare() {
	//附值
	c.controllerName, c.actionName = c.GetControllerAndAction()
	//从Session里获取数据 设置用户信息
	c.adapterUserInfo()
	//获取panic
	defer func() {
		if panicError := recover(); panicError != nil {
			var buf []byte = make([]byte, 1024)
			runtimec := runtime.Stack(buf, false)
			beego.Error("控制器错误:", panicError, string(buf[0:runtimec]))
		}
	}()
}

//从session里取用户信息
func (c *BaseController) adapterUserInfo() {
	a := c.GetSession("loginUser")
	if a != nil {
		c.curUser = a.(models.User)
		c.Data["loginUser"] = a
	}
}

//SetBackendUser2Session 获取用户信息（包括资源UrlFor）保存至Session
func (c *BaseController) setLoginUser2Session(userId int) error {
	user, err := models.GetUserById(userId)
	if err != nil {
		return err
	}
	//获取这个用户能获取到的所有资源列表
	sessionData := make(map[string]interface{})
	sessionData["username"] = user.Username
	c.SetSession("loginUser", sessionData)
	return nil
}

func (c *BaseController) jsonResult(code enums.JsonResultCode, msg string, data interface{}) {
	r := &JsonResult{code, msg, data}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

// 重定向
func (c *BaseController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

// 重定向 去错误页
func (c *BaseController) pageError(msg string) {
	errorUrl := c.URLFor("HomeController.Error") + "/" + msg
	c.Redirect(errorUrl, 302)
	c.StopRun()
}

// 重定向 去登录页
func (c *BaseController) pageLogin() {
	url := c.URLFor("HomeController.Login")
	c.Redirect(url, 302)
	c.StopRun()
}

//获取用户IP地址
func (c *BaseController) getClientIp() string {
	s := strings.Split(c.Ctx.Request.RemoteAddr, ":")
	return s[0]
}
