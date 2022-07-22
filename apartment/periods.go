package apartment

import (
	"time"
)

type Period struct {
	Id string
}

type Month struct {
	Period
}

func (month *Month) NextMonth() *Month {
	nextMonth := new(Month)

	monthTime, _ := time.Parse("2006-01", month.Id)
	monthTime = monthTime.AddDate(0, 1, 0)

	nextMonth.Id = monthTime.Format("2006-01")

	return nextMonth
}

func CurrentMonth() *Month {
	month := new(Month)
	currentTime := time.Now()
	month.Id = currentTime.Format("2006-01")
	return month
}
