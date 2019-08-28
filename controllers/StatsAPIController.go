package controllers

import (
	"github.com/astaxie/beego"
	"github.com/openstars/SimpleServiceMonitorLib/models"
	"github.com/openstars/SimpleServiceMonitorLib/models/data"
	"github.com/openstars/SimpleServiceMonitorLib/restapi/responses"
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


	cp2 := models.DefaultMonitorModel().StartAction("Act_GetListActions")
	u.Data["json"] = models.DefaultMonitorModel().GetAllActionNames();

	models.DefaultMonitorModel().FinishAction(cp2);

	u.ServeJSON()
	
}

// @Title GetListRequests
// @Description Get total request names 
// @Success 200 {json} 	[]string
// @router /getrequests [get]
func (u *StatAPIController) GetListRequests() {
	cp := models.DefaultMonitorModel().StartRequest("Req_GetListRequests")
	defer models.DefaultMonitorModel().FinishRequest(cp); 

	cp2 := models.DefaultMonitorModel().StartAction("Act_GetListRequests")

	u.Data["json"] = models.DefaultMonitorModel().GetAllRequestNames();
	
	models.DefaultMonitorModel().FinishAction(cp2);

	u.ServeJSON()
	
}


// @Title GetActionsDataByHours
// @Description Get total hourly action data 
// @Success 200 {object}	responses.ActionsDataByHours
// @router /totalactionshourly [get]
func (u *StatAPIController) GetActionsDataByHours() {
	cp := models.DefaultMonitorModel().StartRequest("Req_GetActionsDataByHours")
	defer models.DefaultMonitorModel().FinishRequest(cp); 
	
	aRes := responses.ActionsDataByHours{
		HourData: make(map[string] [24]data.SecondStat ),
	}  // = make(responses.ActionsDataByHours);
	models.DefaultMonitorModel().GetActionDataByHour(&aRes);
	u.Data["json"] = aRes;
	u.ServeJSON();
}



// @Title GetRequestDataByHour
// @Description Get total hourly action data 
// @Success 200 {object}	responses.ActionsDataByHours
// @router /totalrequestshourly [get]
func (u *StatAPIController) GetRequestDataByHour() {
	cp := models.DefaultMonitorModel().StartRequest("Req_GetRequestDataByHours")
	defer models.DefaultMonitorModel().FinishRequest(cp); 
	
	// var aRes responses.ActionsDataByHours // = make(responses.ActionsDataByHours);

	aRes := responses.ActionsDataByHours{
		HourData: make(map[string] [24]data.SecondStat ),
	} 
	models.DefaultMonitorModel().GetRequestDataByHour(&aRes);
	u.Data["json"] = aRes;
	u.ServeJSON();
}

