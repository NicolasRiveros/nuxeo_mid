package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/nuxeo_mid/controllers:ValidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/nuxeo_mid/controllers:ValidacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/nuxeo_mid/controllers:ValidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/nuxeo_mid/controllers:ValidacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/nuxeo_mid/controllers:ValidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/nuxeo_mid/controllers:ValidacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/nuxeo_mid/controllers:WorkflowController"] = append(beego.GlobalControllerRouter["github.com/udistrital/nuxeo_mid/controllers:WorkflowController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/nuxeo_mid/controllers:WorkflowController"] = append(beego.GlobalControllerRouter["github.com/udistrital/nuxeo_mid/controllers:WorkflowController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/nuxeo_mid/controllers:WorkflowController"] = append(beego.GlobalControllerRouter["github.com/udistrital/nuxeo_mid/controllers:WorkflowController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
