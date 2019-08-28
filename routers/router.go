// @APIVersion 1.0.0
// @Title API for tracking web request and internal action performance
// @Description This is sample for fast tracking api performance. Opensource by TrustKeys Network., 2019
// @Contact thanhnt@sonek.vn
// @TermsOfServiceUrl https://trustkeys.network
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/openstars/SimpleServiceMonitorLib/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/commonservice/stats",
			beego.NSInclude(
				&controllers.StatAPIController{},
			),
		),
		// beego.NSNamespace("/user",
		// 	beego.NSInclude(
		// 		&controllers.UserController{},
		// 	),
		// ),
	)
	beego.AddNamespace(ns)
}
