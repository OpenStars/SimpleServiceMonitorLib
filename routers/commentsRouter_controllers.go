package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/openstars/SimpleServiceMonitorLib/controllers:StatAPIController"] = append(beego.GlobalControllerRouter["github.com/openstars/SimpleServiceMonitorLib/controllers:StatAPIController"],
        beego.ControllerComments{
            Method: "GetListActions",
            Router: `/getactions`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/openstars/SimpleServiceMonitorLib/controllers:StatAPIController"] = append(beego.GlobalControllerRouter["github.com/openstars/SimpleServiceMonitorLib/controllers:StatAPIController"],
        beego.ControllerComments{
            Method: "GetListRequests",
            Router: `/getrequests`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/openstars/SimpleServiceMonitorLib/controllers:StatAPIController"] = append(beego.GlobalControllerRouter["github.com/openstars/SimpleServiceMonitorLib/controllers:StatAPIController"],
        beego.ControllerComments{
            Method: "GetActionsDataByHours",
            Router: `/totalactionshourly`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/openstars/SimpleServiceMonitorLib/controllers:StatAPIController"] = append(beego.GlobalControllerRouter["github.com/openstars/SimpleServiceMonitorLib/controllers:StatAPIController"],
        beego.ControllerComments{
            Method: "GetRequestDataByHour",
            Router: `/totalrequestshourly`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
