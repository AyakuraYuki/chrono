package chrono

import (
	"strings"
	"time"
)

var (
	minimumDuration time.Duration = -1 << 63
	maximumDuration time.Duration = 1<<63 - 1
)

// DiffInYears gets the difference in years
func (t Time) DiffInYears(rhs ...Time) int64 {
	start, end := t, t.Now()
	if len(rhs) > 0 {
		end = rhs[0]
	}
	diffY, diffM, diffD := end.Year()-start.Year(), end.Month()-start.Month(), end.Day()-start.Day()
	if diffM < 0 || (diffM == 0 && diffD < 0) {
		diffY--
	}
	if diffY < 0 && (diffD != 0 || diffM != 0) {
		diffY++
	}
	return int64(diffY)
}

// DiffAbsInYears gets the difference in years with absolute value.
func (t Time) DiffAbsInYears(rhs ...Time) int64 {
	return getAbsValue(t.DiffInYears(rhs...))
}

// DiffInMonths gets the difference in months
func (t Time) DiffInMonths(rhs ...Time) int64 {
	start, end := t, t.Now()
	if len(rhs) > 0 {
		end = rhs[0]
	}
	if start.Year() == end.Year() && start.Month() == end.Month() {
		return 0
	}
	diffD := s
}

func (t Time) DiffForHumans(rhs ...Time) string {
	end := t.Now()
	if len(rhs) > 0 {
		end = rhs[0]
	}
	if t.HasError() || end.HasError() {
		return ""
	}
	unit, value := t.diff(end)
	translation := t.lang.translate(unit, getAbsValue(value))
	if unit == "now" {
		return translation
	}
	if t.Lt(end) && len(rhs) == 0 {
		return strings.Replace(t.lang.resources["ago"], "%s", translation, 1)
	}
	if t.Lt(end) && len(rhs) > 0 {
		return strings.Replace(t.lang.resources["before"], "%s", translation, 1)
	}
	if t.Gt(end) && len(rhs) == 0 {
		return strings.Replace(t.lang.resources["from_now"], "%s", translation, 1)
	}
	return strings.Replace(t.lang.resources["after"], "%s", translation, 1)
}

func (t Time) diff(end Time) (unit string, value int64) {
	switch {
	case t.DiffAbsInYears(end) > 0:
		unit, value = "year", t.DiffInYears(end)
	case t.DiffAbsInMonths(end) > 0:
		unit, value = "month", t.DiffInMonths(end)
	case t.DiffAbsInWeeks(end) > 0:
		unit, value = "week", t.DiffInWeeks(end)
	case t.DiffAbsInDays(end) > 0:
		unit, value = "day", t.DiffInDays(end)
	case t.DiffAbsInHours(end) > 0:
		unit, value = "hour", t.DiffInHours(end)
	case t.DiffAbsInMinutes(end) > 0:
		unit, value = "minute", t.DiffInMinutes(end)
	case t.DiffAbsInSeconds(end) > 0:
		unit, value = "second", t.DiffInSeconds(end)
	case t.DiffAbsInSeconds(end) == 0:
		unit, value = "now", 0
	}
	return
}

func diffInMonths(start, end Time) int64 {
	y, m, d, h, i, s, ns := start.DateTimeNano()
	endY, endM, _ := end.Date()

	diffY := endY - y
	diffM := endM - m
	totalM := diffY*12 + diffM

	if time.Date(y, time.Month(m+totalM), d, h, i, s, ns, start.StdTime().Location()).After(end.StdTime()) {
		return int64(totalM - 1)
	}
	return int64(totalM)
}
