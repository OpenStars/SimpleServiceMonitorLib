package controllers

import (
	"github.com/astaxie/beego"
	"github.com/openstars/SimpleServiceMonitorLib/models"
)

// Operations about Users
type StatAPIController struct {
	beego.Controller
}

// Listing danh sách request, danh sách actions 

// Xem biểu đồ thể hiện lượng request trong ngày. Độ chi tiết theo : Ngày, Giờ, Phút 

// Cụ thể có thể lấy chi tiết theo giây 


// Store data to Bigset (làm sau)


// @Title GetListActions
// @Description Get total action names 
// @Success 200 {json} 	[]string
// @router /getactions [get]
func (u *StatAPIController) GetListActions() {
	cp := models.DefaultMonitorModel().StartRequest("Req_GetListActions")
	defer models.DefaultMonitorModel().FinishRequest(cp); 

	u.Data["json"] = models.DefaultMonitorModel().GetAllActionNames();
	u.ServeJSON()
	
}

// @Title GetListRequests
// @Description Get total request names 
// @Success 200 {json} 	[]string
// @router /getrequests [get]
func (u *StatAPIController) GetListRequests() {
	cp := models.DefaultMonitorModel().StartRequest("Req_GetListRequests")
	defer models.DefaultMonitorModel().FinishRequest(cp); 
	
	u.Data["json"] = models.DefaultMonitorModel().GetAllRequestNames();
	u.ServeJSON()
	
}


// @Title GetActionsDataByHours
// @Description Get total request names 
// @Success 200 {object} 	[]string
// @router /getrequests [get]
func (u *StatAPIController) GetActionsDataByHours() {

}