package responses
import (
	"github.com/openstars/SimpleServiceMonitorLib/models/data"
)



type ActionsDataByHours struct {
	HourData   map[string] [24]data.SecondStat `json:"hourlydata"` // reset hàng ngày 
}
