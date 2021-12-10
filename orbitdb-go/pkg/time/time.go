package time

import (
	"sort"
	"time"

	"github.com/jinzhu/now"
)

type (
	Time     = time.Time
	Duration = time.Duration
)

type Date struct {
	Year  int
	Month int
	Day   int
}

const (
	Nanosecond  = time.Nanosecond
	Microsecond = time.Microsecond
	Millisecond = time.Millisecond
	Second      = time.Second
	Minute      = time.Minute
	Hour        = time.Hour
	Day         = 24 * Hour

	DaySunday    = time.Sunday
	DayMonday    = time.Monday
	DayTuesday   = time.Tuesday
	DayWednesday = time.Wednesday
	DayThursday  = time.Thursday
	DayFriday    = time.Friday
	DaySaturday  = time.Saturday

	MonthJanuary   = time.January
	MonthFebruary  = time.February
	MonthMarch     = time.March
	MonthApril     = time.April
	MonthMay       = time.May
	MonthJune      = time.June
	MonthJuly      = time.July
	MonthAugust    = time.August
	MonthSeptember = time.September
	MonthOctober   = time.October
	MonthNovember  = time.November
	MonthDecember  = time.December

	LayoutANSIC       = time.ANSIC
	LayoutUnixDate    = time.UnixDate
	LayoutRubyDate    = time.RubyDate
	LayoutRFC822      = time.RFC822
	LayoutRFC822Z     = time.RFC822Z
	LayoutRFC850      = time.RFC850
	LayoutRFC1123     = time.RFC1123
	LayoutRFC1123Z    = time.RFC1123Z
	LayoutRFC3339     = time.RFC3339
	LayoutRFC3339Nano = time.RFC3339Nano
	LayoutKitchen     = time.Kitchen
)

var (
	New   = time.Date
	Now   = time.Now
	Parse = time.Parse
	Since = time.Since
	Until = time.Until
	Sleep = time.Sleep
	Unix  = time.Unix
)

func Nowptr() *Time { // because fuck go, you can not do &(Now())
	t := Now()
	return &t
}

//

func BeginningOfMinute(t Time) Time  { return now.New(t).BeginningOfMinute() }
func BeginningOfHour(t Time) Time    { return now.New(t).BeginningOfHour() }
func BeginningOfDay(t Time) Time     { return now.New(t).BeginningOfDay() }
func BeginningOfWeek(t Time) Time    { return now.New(t).BeginningOfWeek() }
func BeginningOfMonth(t Time) Time   { return now.New(t).BeginningOfMonth() }
func BeginningOfQuarter(t Time) Time { return now.New(t).BeginningOfQuarter() }
func BeginningOfHalf(t Time) Time    { return now.New(t).BeginningOfHalf() }
func BeginningOfYear(t Time) Time    { return now.New(t).BeginningOfYear() }
func EndOfMinute(t Time) Time        { return now.New(t).EndOfMinute() }
func EndOfHour(t Time) Time          { return now.New(t).EndOfHour() }
func EndOfDay(t Time) Time           { return now.New(t).EndOfDay() }
func EndOfWeek(t Time) Time          { return now.New(t).EndOfWeek() }
func EndOfMonth(t Time) Time         { return now.New(t).EndOfMonth() }
func EndOfQuarter(t Time) Time       { return now.New(t).EndOfQuarter() }
func EndOfHalf(t Time) Time          { return now.New(t).EndOfHalf() }
func EndOfYear(t Time) Time          { return now.New(t).EndOfYear() }

func AddDate(t Time, d Date) Time { return t.AddDate(d.Year, d.Month, d.Day) }

//

func BusinessDay(t Time) bool {
	// TODO: register holidays for location here
	switch t.Weekday() {
	case DaySaturday:
	case DaySunday:
	default:
		return false
	}
	return true
}

func ClosestBusinessDay(t Time) Time {
	for {
		if BusinessDay(t) {
			break
		}
		t = t.Add(24 * Hour)
	}
	return t
}

//

func Samples(start Time, buf []Time, next func(current Time) Time) {
	t := start

	for n := 0; n < cap(buf); n++ {
		t = next(t)
		buf[n] = t
	}
}

func Range(start Time, end Time, next func(current Time) Time) []Time {
	var (
		t   = start
		buf = []Time{}
	)

	for t.Before(end) {
		t = next(t)
		buf = append(buf, t)
	}

	return buf
}

func Uniq(ts []Time) []Time {
	var (
		seen = make(map[Time]struct{}, len(ts))
		buf  = make([]Time, len(ts))
		k    int
		ok   bool
	)

	for _, v := range ts {
		if _, ok = seen[v]; !ok {
			seen[v] = struct{}{}
			buf[k] = v
			k++
		}
	}

	return buf[:k]
}

func Sort(ts []Time) {
	sort.Slice(
		ts,
		func(i int, j int) bool {
			return ts[i].Before(ts[j])
		},
	)
}
