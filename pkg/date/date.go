package date

import "time"

type Date struct {
	Day   int
	Month time.Month
	Year  int
}

func FromTime(t time.Time) *Date {
	return &Date{
		Day:   t.Day(),
		Month: t.Month(),
		Year:  t.Year(),
	}
}
