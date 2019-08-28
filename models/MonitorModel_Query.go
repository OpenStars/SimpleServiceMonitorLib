package models


/*
Function: GetAllActionNames
Purpose: 
*/
func (md *MonitorModel) GetAllActionNames()[]string{
	var res []string;
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
	var res []string;
	md.RequestDailyLog.Range( func(key, _ interface{} ) bool {
		sKey := key.(string)
		res = append(res, sKey) 
		// dailyStat := value.(*DayStat)
		return true;
	} )
	return res
}