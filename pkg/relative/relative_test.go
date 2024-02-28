package relative_test

import (
	"testing"
	"time"

	"github.com/liamawhite/schedule/pkg/relative"
)

func TestNextNamedDay(t *testing.T) {
    tests := []struct {
        name string
        current time.Time
        day time.Weekday
        want time.Time
    } {
        {
            name: "Monday, next Monday",
            current: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
            day: time.Monday,
            want: time.Date(2018, 1, 8, 0, 0, 0, 0, time.UTC),
        },
        {
            name: "Monday, next Tuesday",
            current: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
            day: time.Tuesday,
            want: time.Date(2018, 1, 2, 0, 0, 0, 0, time.UTC),
        },
        {
            name: "Wednesday, next Monday",
            current: time.Date(2018, 1, 3, 0, 0, 0, 0, time.UTC),
            day: time.Monday,
            want: time.Date(2018, 1, 8, 0, 0, 0, 0, time.UTC),
        },
    }
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            got := relative.NextNamedDay(test.current, test.day)
            if got != test.want {
                t.Errorf("expected %v, got %v", test.want.Format(time.RFC3339), got.Format(time.RFC3339))
            }
        })
    }
}

func TestNextWeekday(t *testing.T) {
    tests := []struct {
        name string
        current time.Time
        want time.Time
    } {
        {
            name: "Monday",
            current: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
            want: time.Date(2018, 1, 2, 0, 0, 0, 0, time.UTC),
        },
        {
            name: "Friday",
            current: time.Date(2018, 1, 5, 0, 0, 0, 0, time.UTC),
            want: time.Date(2018, 1, 8, 0, 0, 0, 0, time.UTC),
        },
        {
            name: "Saturday",
            current: time.Date(2018, 1, 6, 0, 0, 0, 0, time.UTC),
            want: time.Date(2018, 1, 8, 0, 0, 0, 0, time.UTC),
        },
        {
            name: "Sunday",
            current: time.Date(2018, 1, 7, 0, 0, 0, 0, time.UTC),
            want: time.Date(2018, 1, 8, 0, 0, 0, 0, time.UTC),
        },
    }
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            got := relative.NextWeekday(test.current)
            if got != test.want {
                t.Errorf("expected %v, got %v", test.want.Format(time.RFC3339), got.Format(time.RFC3339))
            }
        })
    }
}

func TestNextSpecifiedMonthDay(t *testing.T) {
    tests := []struct {
        name string
        current time.Time
        day int
        want time.Time
    } {
        {
            name: "January 12th, next 1st",
            current: time.Date(2018, 1, 12, 0, 0, 0, 0, time.UTC),
            day: 1,
            want: time.Date(2018, 2, 1, 0, 0, 0, 0, time.UTC),
        },
        {
            name: "January 15th, next 15th",
            current: time.Date(2018, 1, 15, 0, 0, 0, 0, time.UTC),
            day: 15,
            want: time.Date(2018, 2, 15, 0, 0, 0, 0, time.UTC),
        },
        {
            name: "January 21st, next 25th",
            current: time.Date(2018, 1, 21, 0, 0, 0, 0, time.UTC),
            day: 25,
            want: time.Date(2018, 1, 25, 0, 0, 0, 0, time.UTC),
        },
        {
            name: "Dec 20th, next 1st",
            current: time.Date(2018, 12, 20, 0, 0, 0, 0, time.UTC),
            day: 1,
            want: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
        },
        {
            name: "Dec 20th, next 25th",
            current: time.Date(2018, 12, 20, 0, 0, 0, 0, time.UTC),
            day: 25,
            want: time.Date(2018, 12, 25, 0, 0, 0, 0, time.UTC),
        },        
    }
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            got := relative.NextSpecifiedMonthDay(test.current, test.day)
            if got != test.want {
                t.Errorf("expected %v, got %v", test.want.Format(time.RFC3339), got.Format(time.RFC3339))
            }
        })
    }
}

func TestNextNamedMonth(t *testing.T) {
    tests := []struct {
        name string
        current time.Time
        month time.Month
        want time.Time
    } {
        {
            name: "January 12th, next February",
            current: time.Date(2018, 1, 12, 0, 0, 0, 0, time.UTC),
            month: time.February,
            want: time.Date(2018, 2, 1, 0, 0, 0, 0, time.UTC),
        },
        {
            name: "January 15th, next January",
            current: time.Date(2018, 1, 15, 0, 0, 0, 0, time.UTC),
            month: time.January,
            want: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
        },
        {
            name: "January 21st, next December",
            current: time.Date(2018, 1, 21, 0, 0, 0, 0, time.UTC),
            month: time.December,
            want: time.Date(2018, 12, 1, 0, 0, 0, 0, time.UTC),
        },
    }
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            got := relative.NextNamedMonth(test.current, test.month)
            if got != test.want {
                t.Errorf("expected %v, got %v", test.want.Format(time.RFC3339), got.Format(time.RFC3339))
            }
        })
    }
}

func TestNextMonth(t *testing.T) {
    tests := []struct {
        name string
        current time.Time
        want time.Time
    } {
        {
            name: "January 21st",
            current: time.Date(2018, 1, 21, 0, 0, 0, 0, time.UTC),
            want: time.Date(2018, 2, 1, 0, 0, 0, 0, time.UTC),
        },
        {
            name: "December 17th",
            current: time.Date(2018, 12, 17, 0, 0, 0, 0, time.UTC),
            want: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
        },
    }
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            got := relative.NextMonth(test.current)
            if got != test.want {
                t.Errorf("expected %v, got %v", test.want.Format(time.RFC3339), got.Format(time.RFC3339))
            }
        })
    }
}

func TestNextSpecifiedYear(t *testing.T) {
    tests := []struct {
        name string
        current time.Time
        year int
        want time.Time
        wantErr bool
    } {
        {
            name: "January 21st 2018, 2020",
            current: time.Date(2018, 1, 21, 0, 0, 0, 0, time.UTC),
            year: 2020,
            want: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
        },
        {
            name: "December 17th 2018, 2100",
            current: time.Date(2018, 12, 17, 0, 0, 0, 0, time.UTC),
            year: 2100,
            want: time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC),
        },
        {
            name: "year is in the past",
            current: time.Date(2018, 1, 21, 0, 0, 0, 0, time.UTC),
            year: 2017,
            wantErr: true,
        },
    }
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            got, err := relative.NextSpecifiedYear(test.current, test.year)
            
            if got != test.want {
                t.Errorf("expected %v, got %v", test.want.Format(time.RFC3339), got.Format(time.RFC3339))
            }
            if test.wantErr && err == nil {
                t.Errorf("expected an error, got nil")
            }
            if !test.wantErr && err != nil {
                t.Errorf("expected no error, got %v", err)
            }
            
        })
    }
}

func TestNextYear(t *testing.T) {
    tests := []struct {
        name string
        current time.Time
        want time.Time
    } {
        {
            name: "January 21st",
            current: time.Date(2018, 1, 21, 0, 0, 0, 0, time.UTC),
            want: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
        },
        {
            name: "December 17th",
            current: time.Date(2018, 12, 17, 0, 0, 0, 0, time.UTC),
            want: time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
        },
    }
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            got := relative.NextYear(test.current)
            if got != test.want {
                t.Errorf("expected %v, got %v", test.want.Format(time.RFC3339), got.Format(time.RFC3339))
            }
        })
    }
}
