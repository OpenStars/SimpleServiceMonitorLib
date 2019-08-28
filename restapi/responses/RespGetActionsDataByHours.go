package responses
import (
	"github.com/openstars/SimpleServiceMonitorLib/models"
)



type ActionsDataByHours struct {
	HourData   [24]models.SecondStat `json:"hourlydata"` // reset hàng ngày 
}
