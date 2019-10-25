// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"zyra_core/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/publish",
			beego.NSInclude(
				&controllers.PublishRecordController{},
			),
		),
		beego.NSNamespace("/appid",
			beego.NSInclude(
				&controllers.AppidController{},
			),
		),
		beego.NSNamespace("/build",
			beego.NSInclude(
				&controllers.BuildRecordController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
