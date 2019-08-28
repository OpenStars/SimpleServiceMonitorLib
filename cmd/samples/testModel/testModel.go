package main

import (
	"github.com/openstars/SimpleServiceMonitorLib/models"
	"fmt"
	"time"
)

func main(){
	m := models.NewMonitorModel()
	m.LogAction("Thrift_GetProfile", 1000000);	
	X := m.GetActionStat("Thrift_GetProfile");
	currentTime := time.Now() ; 
	dailySecond := currentTime.Hour() * 3600 + currentTime.Minute()*60 + currentTime.Second();
	actionStat := X//md.GetActionStat(actionName)
	// actionStat.SecondData[dailySecond].TotalCount++;
	// actionStat.SecondData[dailySecond].TotalTime += runTime;
	fmt.Println(actionStat.SecondData[dailySecond])

	fmt.Println(m.GetAllActionNames() )
	
	// m.ActionDailyLog["CallThrift_Profile"] = &models.DayStat{}
	// X := (&m.ActionDailyLog["CallThrift_Profile"].SecondData[0])
	// X.TotalTime = 100000000
	// X.TotalCount = 500;

	// m.ActionDailyLog["CallThrift_Profile"].HourData[0].SecondData[0].TotalTime = 100000000;
	// fmt.Println(m);
	// fmt.Println(X)
}