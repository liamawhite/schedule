package recurrence

import (
	"time"

    "github.com/liamawhite/schedule/pkg/date"
	"github.com/teambition/rrule-go"
)

func NewRecurrence(opts Option) (*Recurrence, error) {
    rule, err := rrule.NewRRule(rrule.ROption(opts))
    if err != nil {
        return nil, err
    }
    return &Recurrence{*rule}, nil
}

type Recurrence struct {
    rrule.RRule
}

type Option rrule.ROption

func (r *Recurrence) Next() *date.Date {
    if next, ok := r.RRule.Iterator()(); ok {
        return date.FromTime(next)
    }
    return nil
}

func (r *Recurrence) All() []*date.Date {
    instances := r.RRule.All()
    dates := make([]*date.Date, len(instances))
    for i, instance := range instances {
        dates[i] = date.FromTime(instance)
    }
    return dates
}

func (r *Recurrence) Between(after, before time.Time, inc bool) []*date.Date {
    instances := r.RRule.Between(after, before, inc)
    dates := make([]*date.Date, len(instances))
    for i, instance := range instances {
        dates[i] = date.FromTime(instance)
    }
    return dates
}

