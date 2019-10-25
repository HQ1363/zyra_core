package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["zyra_core/controllers:AppidController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:AppidController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:AppidController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:AppidController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:AppidController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:AppidController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:bid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:AppidController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:AppidController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:bid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:AppidController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:AppidController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:bid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:BuildRecordController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:BuildRecordController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:BuildRecordController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:BuildRecordController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:BuildRecordController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:BuildRecordController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:bid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:BuildRecordController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:BuildRecordController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:bid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:BuildRecordController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:BuildRecordController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:bid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:PublishRecordController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:PublishRecordController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:PublishRecordController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:PublishRecordController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:PublishRecordController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:PublishRecordController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:bid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:PublishRecordController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:PublishRecordController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:bid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:PublishRecordController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:PublishRecordController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:bid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:UserController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:UserController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:UserController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:UserController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:UserController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:UserController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["zyra_core/controllers:UserController"] = append(beego.GlobalControllerRouter["zyra_core/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
