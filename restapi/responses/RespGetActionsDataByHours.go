package responses
import (
	"gitlab.123xe.vn/OpenPlatform/SimpleServiceMonitorLib/models"
)



type ActionsDataByHours struct {
	HourData   [24]models.SecondStat `json:"hourlydata"` // reset hàng ngày 
}
