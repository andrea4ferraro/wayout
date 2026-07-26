package calculator

import (
	"time"
	"token-vesting-calculator/internal/exporter"
)

type Schedule struct {
	Total  int
	Start  time.Time
	Months int
}

func DefaultSchedule() Schedule {

	return Schedule{
		Total: 1000000,
		Start: time.Now().AddDate(0, -8, 0),
		Months: 24,
	}
}

func (s Schedule) Unlocked() int {

	months := int(time.Since(s.Start).Hours() / 24 / 30)

	if months > s.Months {
		months = s.Months
	}

	return s.Total * months / s.Months
}

func (s Schedule) Locked() int {

	return s.Total - s.Unlocked()
}

func (s Schedule) Progress() float64 {

	return float64(s.Unlocked()) * 100 / float64(s.Total)
}

func (s Schedule) ExportCalendar() {

	exporter.Export(s)
}
