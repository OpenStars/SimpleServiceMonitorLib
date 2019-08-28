package models

import (
	"time"
)

type SecondStat struct{
	TotalCount uint64		 `json:"totalcount"`
	TotalTime time.Duration   `json:"totaltime"`
}

type HourStat struct{ 
	SecondData [3600]SecondStat `json:"dataInsecond`
}

// func (hs *HourStat) 

type DayStat struct {
	// HourData [24]HourStat
	SecondData [24*3600]SecondStat `json:"dataInsecond`
	MinuteData [24*60]SecondStat `json:"datainminute"`// 
	HourData   [24]SecondStat `json:"hourlydata"` // reset hàng ngày 
	WholeDay SecondStat
}
