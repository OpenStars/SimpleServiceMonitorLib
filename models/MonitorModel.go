package models
import (
	"time"
	"fmt"
	"sync"
	"github.com/openstars/SimpleServiceMonitorLib/models/data"
)

var (
		defaultMonitorModel = NewMonitorModel()
)

func DefaultMonitorModel() *MonitorModel{
	return defaultMonitorModel;
}

/* MonitorModel
*/
type MonitorModel struct {
	ActionDailyLog *sync.Map //map[string]*DayStat
	RequestDailyLog *sync.Map //map[string]*DayStat
	
	LastHour int
	LastTime time.Time;

}

type DayStat = data.DayStat

func NewMonitorModel() *MonitorModel{
	return &MonitorModel{
		ActionDailyLog: &sync.Map{} , //make( map[string]*DayStat),
		RequestDailyLog: &sync.Map{}, //make ( map[string]*DayStat) ,
		LastHour: time.Now().Hour(),
		LastTime: time.Now(),
	}
}

func (md*MonitorModel) GetActionStat(actionName string) *DayStat{
	res, _ := md.ActionDailyLog. LoadOrStore( actionName, &DayStat{} ) 
	return res.(*DayStat)
}

func (md*MonitorModel) GetRequestStat(requestName string) *DayStat{
	res, _ := md.RequestDailyLog.LoadOrStore( requestName, &DayStat{} ) 
	return res.(*DayStat)
}

func (md *MonitorModel) CheckReset(){
	currentHour := time.Now().Hour();
	if (md.LastHour > currentHour) {
		md.LastHour = currentHour;
		md.ActionDailyLog = &sync.Map{} //make( map[string]*DayStat),
		md.RequestDailyLog = &sync.Map{} 
		md.LastTime = time.Now();
	}
}

func (md *MonitorModel) LogAction(actionName string, runTime time.Duration ){
	// fmt.Println("LogAction ", actionName, runTime)
	currentTime := time.Now() ; 
	currentHour := currentTime.Hour()
	currentSecondsInDay := currentTime.Hour() * 3600 + currentTime.Minute()*60 + currentTime.Second();

	actionStat := md.GetActionStat(actionName)

	actionStat.SecondData[currentSecondsInDay].TotalCount++;
	actionStat.SecondData[currentSecondsInDay].TotalTime += runTime;

	currentMinsInday :=  currentTime.Hour() * 60 + currentTime.Minute();
	actionStat.MinuteData[currentMinsInday].TotalCount ++;
	actionStat.MinuteData[currentMinsInday].TotalTime += runTime;

	actionStat.HourData[currentHour].TotalCount ++;
	actionStat.HourData[currentHour].TotalTime += runTime;
	actionStat.WholeDay.TotalCount ++;
	actionStat.WholeDay.TotalTime += runTime;
}

func (md *MonitorModel) LogRequest(requestName string, runTime time.Duration){

	// fmt.Println("LogRequest ", requestName, runTime)
	currentTime := time.Now() ; 
	currentHour := currentTime.Hour()
	currentSecondsInDay := currentTime.Hour() * 3600 + currentTime.Minute()*60 + currentTime.Second();

	actionStat := md.GetRequestStat(requestName)

	actionStat.SecondData[currentSecondsInDay].TotalCount++;
	actionStat.SecondData[currentSecondsInDay].TotalTime += runTime;

	currentMinsInday :=  currentTime.Hour() * 60 + currentTime.Minute();
	actionStat.MinuteData[currentMinsInday].TotalCount ++;
	actionStat.MinuteData[currentMinsInday].TotalTime += runTime;

	actionStat.HourData[currentHour].TotalCount ++;
	actionStat.HourData[currentHour].TotalTime += runTime;
	actionStat.WholeDay.TotalCount ++;
	actionStat.WholeDay.TotalTime += runTime;

}

type  MonitorCheckpoint struct{
	StartTime time.Time
	Name  string
}

func (md *MonitorModel) StartAction(actionName string) MonitorCheckpoint{
	return MonitorCheckpoint{Name: actionName , StartTime: time.Now() }
}

func (md *MonitorModel) FinishAction(cp MonitorCheckpoint){
	md.LogAction(cp.Name, time.Since(cp.StartTime) )
}

func (md *MonitorModel) StartRequest(reqName string) MonitorCheckpoint{
	return MonitorCheckpoint{Name: reqName , StartTime: time.Now() }
}


func (md *MonitorModel) FinishRequest(cp MonitorCheckpoint){
	fmt.Println("Finish Request ", cp);
	md.LogRequest(cp.Name, time.Since(cp.StartTime) )
}

