package models

import (
	"github.com/openstars/SimpleServiceMonitorLib/restapi/responses"
	"fmt"
)

/*
Function: GetAllActionNames
Purpose: 
*/
func (md *MonitorModel) GetAllActionNames()[]string{
	// var res []string;
	res := []string{};
	fmt.Println("GetAllActionNames")
	md.ActionDailyLog.Range( func(key, _ interface{} ) bool {
		sKey := key.(string)
		res = append(res, sKey) 
		// dailyStat := value.(*DayStat)
		return true;
	} )
	return res
}


/*
Function: GetAllRequestNames
Purpose: get all request name 
*/
func (md *MonitorModel) GetAllRequestNames()[]string{
	res := []string{};
	md.RequestDailyLog.Range( func(key, _ interface{} ) bool {
		sKey := key.(string)
		res = append(res, sKey) 
		// dailyStat := value.(*DayStat)
		return true;
	} )
	return res
}

func (md* MonitorModel) GetActionDataByHour(out *responses.ActionsDataByHours) {
	md.ActionDailyLog.Range( func(key, value interface{} ) bool {
		// add key - value to 
		sKey := key.(string)
		aDayStat := value.(*DayStat);
		out.HourData[sKey] = aDayStat.HourData //need to copy?
		fmt.Println("action data of ", sKey,  aDayStat.HourData )
		return true;
	} )
}

func (md* MonitorModel) GetRequestDataByHour(out *responses.ActionsDataByHours) {
	md.RequestDailyLog.Range( func(key, value interface{} ) bool {
		// add key - value to 
		sKey := key.(string)
		aDayStat := value.(*DayStat);
		out.HourData[sKey] = aDayStat.HourData //need to copy?
		fmt.Println("request data of ", sKey,  aDayStat.HourData )
		return true;
	} )
}

