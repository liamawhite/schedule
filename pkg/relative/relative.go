package relative

import (
	"fmt"
	"time"
)

func NextWeek(current time.Time) time.Time {
    return NextNamedDay(current, time.Monday)
}

func NextNamedDay(current time.Time, day time.Weekday) time.Time {
    for {
        current = current.AddDate(0, 0, 1)
        if current.Weekday() == day {
            return current
        }
    }
}

func NextWeekday(current time.Time) time.Time {
    // If it's a weekend or friday, go to the next Monday
    if current.Weekday() == time.Friday || current.Weekday() == time.Saturday || current.Weekday() == time.Sunday {
        return NextNamedDay(current, time.Monday)
    }
    return NextNamedDay(current, current.Weekday()+1)
}

func NextSpecifiedMonthDay(current time.Time, day int) time.Time {
    // If the specified day is greater than the current month day, just go to that day
    if day > current.Day() {
        return time.Date(current.Year(), current.Month(), day, 0, 0, 0, 0, current.Location())
    }
    // If the days are the same, go to the next month
    if day == current.Day() {
        return current.AddDate(0, 1, 0)
    }
    // Otherwise if the specified day is greater than the current month day, go to the next month and set the NextSpecifiedMonthDay
    current = current.AddDate(0, 1, 0)
    return time.Date(current.Year(), current.Month(), day, 0, 0, 0, 0, current.Location())
}


func NextNamedMonth(current time.Time, month time.Month) time.Time {
    // Get to the correct month then go back to the first day of the month
    for {
        current = current.AddDate(0, 1, 0)
        if current.Month() == month {
            break
        }
    }
    return time.Date(current.Year(), month, 1, 0, 0, 0, 0, current.Location())
}

func NextMonth(current time.Time) time.Time {
    if current.Month() == time.December {
        return NextNamedMonth(current, time.January)
    }
    return NextNamedMonth(current, current.Month()+1)
}

func NextSpecifiedYear(current time.Time, year int) (time.Time, error) {
    // If the year is in the past, return an error
    if year < current.Year() {
        return time.Time{}, fmt.Errorf("year %d is before the current time's year %v", year, current.Year())
    }

    // Get to the correct year then go back to the first day of the year
    for {
        current = current.AddDate(1, 0, 0)
        if current.Year() == year {
            break
        }
    }
    return time.Date(year, time.January, 1, 0, 0, 0, 0, current.Location()), nil
}

func NextYear(current time.Time) time.Time {
    // We can ignore the error here because we know the year is in the future
    t, _ := NextSpecifiedYear(current, current.Year()+1)
    return t
}
    
