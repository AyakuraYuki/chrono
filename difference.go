package chrono

import (
	"math"
	"strings"
	"time"
)

// DiffInYears gets the difference in years
func (t Time) DiffInYears(rhs ...Time) int64 {
	start, end := t, t.Now()
	if len(rhs) > 0 {
		end = rhs[0]
	}
	diffY := end.Year() - start.Year()
	diffM := end.Month() - start.Month()
	diffD := end.Day() - start.Day()
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
	start, end := t, Now(t.Timezone())
	if len(rhs) > 0 {
		end = rhs[0]
	}
	if start.Year() == end.Year() && start.Month() == end.Month() {
		return 0
	}
	diffD := start.DiffInDays(end)
	if diffD <= 0 {
		return -1 * diffInMonths(end, start)
	}
	return diffInMonths(start, end)
}

func (t Time) DiffAbsInMonths(rhs ...Time) int64 {
	return getAbsValue(t.DiffInMonths(rhs...))
}

func (t Time) DiffInWeeks(rhs ...Time) int64 {
	start, end := t, Now(t.Timezone())
	if len(rhs) > 0 {
		end = rhs[0]
	}
	if start.IsInvalid() || end.IsInvalid() {
		return 0
	}
	return int64(math.Floor(float64((end.Timestamp() - start.Timestamp()) / (7 * 24 * 3600))))
}

func (t Time) DiffAbsInWeeks(rhs ...Time) int64 {
	return getAbsValue(t.DiffInWeeks(rhs...))
}

func (t Time) DiffInDays(rhs ...Time) int64 {
	start, end := t, Now(t.Timezone())
	if len(rhs) > 0 {
		end = rhs[0]
	}
	if start.IsInvalid() || end.IsInvalid() {
		return 0
	}
	return int64(math.Floor(float64((end.Timestamp() - start.Timestamp()) / (24 * 3600))))
}

func (t Time) DiffAbsInDays(rhs ...Time) int64 {
	return getAbsValue(t.DiffInDays(rhs...))
}

func (t Time) DiffInHours(rhs ...Time) int64 {
	start, end := t, Now(t.Timezone())
	if len(rhs) > 0 {
		end = rhs[0]
	}
	if start.IsInvalid() || end.IsInvalid() {
		return 0
	}
	return start.DiffInSeconds(end) / SecondsPerHour
}

func (t Time) DiffAbsInHours(rhs ...Time) int64 {
	return getAbsValue(t.DiffInHours(rhs...))
}

func (t Time) DiffInMinutes(rhs ...Time) int64 {
	start, end := t, Now(t.Timezone())
	if len(rhs) > 0 {
		end = rhs[0]
	}
	if start.IsInvalid() || end.IsInvalid() {
		return 0
	}
	return start.DiffInSeconds(end) / SecondsPerMinute
}

func (t Time) DiffAbsInMinutes(rhs ...Time) int64 {
	return getAbsValue(t.DiffInMinutes(rhs...))
}

func (t Time) DiffInSeconds(rhs ...Time) int64 {
	start, end := t, Now(t.Timezone())
	if len(rhs) > 0 {
		end = rhs[0]
	}
	if start.IsInvalid() || end.IsInvalid() {
		return 0
	}
	return end.Timestamp() - start.Timestamp()
}

func (t Time) DiffAbsInSeconds(rhs ...Time) int64 {
	return getAbsValue(t.DiffInSeconds(rhs...))
}

func (t Time) DiffInString(rhs ...Time) string {
	start, end := t, Now(t.Timezone())
	if len(rhs) > 0 {
		end = rhs[0]
	}
	if start.IsInvalid() || end.IsInvalid() {
		return ""
	}
	unit, value := start.diff(end)
	return t.lang.translate(unit, value)
}

func (t Time) DiffAbsInString(rhs ...Time) string {
	start, end := t, Now(t.Timezone())
	if len(rhs) > 0 {
		end = rhs[0]
	}
	if start.IsInvalid() || end.IsInvalid() {
		return ""
	}
	unit, value := start.diff(end)
	return t.lang.translate(unit, getAbsValue(value))
}

func (t Time) DiffInDuration(rhs ...Time) time.Duration {
	start, end := t, Now(t.Timezone())
	if len(rhs) > 0 {
		end = rhs[0]
	}
	if start.IsInvalid() || end.IsInvalid() {
		return 0
	}
	return end.StdTime().Sub(start.StdTime())
}

func (t Time) DiffAbsInDuration(rhs ...Time) time.Duration {
	d := t.DiffInDuration(rhs...)
	if d >= 0 {
		return d
	}
	return -d
}

func (t Time) DiffForHumans(rhs ...Time) string {
	start, end := t, Now(t.Timezone())
	if len(rhs) > 0 {
		end = rhs[0]
	}
	if start.IsInvalid() || end.IsInvalid() {
		return ""
	}
	unit, value := start.diff(end)
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
	if start.IsInvalid() || end.IsInvalid() {
		return 0
	}

	y, m, d, h, i, s, ns := start.DateTimeNano()
	endYear, endMonth, _ := end.Date()

	diffY := endYear - y
	diffM := endMonth - m
	totalM := diffY*12 + diffM

	if time.Date(y, time.Month(m+totalM), d, h, i, s, ns, start.StdTime().Location()).After(end.StdTime()) {
		return int64(totalM - 1)
	}
	return int64(totalM)
}
