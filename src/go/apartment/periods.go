package apartment

import (
	"strings"
	"time"
)

type Month struct {
	Id string `json:"id"`
}

func (month *Month) String() string {
	return month.Id
}

func (month *Month) Y() string {
	return strings.Split(month.Id, "-")[0]
}

func (month *Month) M() string {
	return strings.Split(month.Id, "-")[1]
}

func (month *Month) NextMonth() *Month {
	nextMonth := new(Month)

	monthTime, _ := time.Parse("2006-01", month.Id)
	monthTime = monthTime.AddDate(0, 1, 0)

	nextMonth.Id = monthTime.Format("2006-01")

	return nextMonth
}
