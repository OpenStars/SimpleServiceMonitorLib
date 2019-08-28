package main
//This is simple server for testing the lib.
import (
	_ "github.com/openstars/SimpleServiceMonitorLib/routers"

	"github.com/astaxie/beego"
)

func main() {
	// if beego.BConfig.RunMode == "dev" {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/v1/commonservice/swagger"] = "swagger"
	// }
	beego.Run()
}
