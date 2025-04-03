package chrono

import "time"

func (t Time) HasError() bool { return t.Error != nil }

// IsDST reports whether is daylight saving time
func (t Time) IsDST() bool { return t.time.IsDST() }

func (t Time) IsZero() bool { return t.time.IsZero() }

func (t Time) IsValid() bool {
	if t.HasError() {
		return false
	}
	return t.Year() >= MinTime().Year() && t.Year() <= MaxTime().Year() && t.Month() > 0 && t.Day() > 0
}
func (t Time) IsInvalid() bool { return !t.IsValid() }

func (t Time) IsAM() bool { return t.Format("a") == "am" }
func (t Time) IsPM() bool { return t.Format("a") == "pm" }

func (t Time) IsLeapYear() bool {
	if t.IsInvalid() {
		return false
	}
	year := t.Year()
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

func (t Time) IsLongYear() bool {
	if t.IsInvalid() {
		return false
	}
	_, week := time.Date(t.Year(), 12, 31, 0, 0, 0, 0, t.loc).ISOWeek()
	return week == weeksPerLongYear
}

func (t Time) IsMonth(month time.Month) bool {
	if t.IsInvalid() {
		return false
	}
	return t.Month() == int(month)
}

func (t Time) IsJanuary() bool   { return t.IsMonth(time.January) }
func (t Time) IsFebruary() bool  { return t.IsMonth(time.February) }
func (t Time) IsMarch() bool     { return t.IsMonth(time.March) }
func (t Time) IsApril() bool     { return t.IsMonth(time.April) }
func (t Time) IsMay() bool       { return t.IsMonth(time.May) }
func (t Time) IsJune() bool      { return t.IsMonth(time.June) }
func (t Time) IsJuly() bool      { return t.IsMonth(time.July) }
func (t Time) IsAugust() bool    { return t.IsMonth(time.August) }
func (t Time) IsSeptember() bool { return t.IsMonth(time.September) }
func (t Time) IsOctober() bool   { return t.IsMonth(time.October) }
func (t Time) IsNovember() bool  { return t.IsMonth(time.November) }
func (t Time) IsDecember() bool  { return t.IsMonth(time.December) }

func (t Time) isSpecificWeekday(weekday time.Weekday) bool {
	if t.IsInvalid() {
		return false
	}
	return t.StdTime().Weekday() == weekday
}

func (t Time) IsSunday() bool    { return t.isSpecificWeekday(time.Sunday) }
func (t Time) IsMonday() bool    { return t.isSpecificWeekday(time.Monday) }
func (t Time) IsTuesday() bool   { return t.isSpecificWeekday(time.Tuesday) }
func (t Time) IsWednesday() bool { return t.isSpecificWeekday(time.Wednesday) }
func (t Time) IsThursday() bool  { return t.isSpecificWeekday(time.Thursday) }
func (t Time) IsFriday() bool    { return t.isSpecificWeekday(time.Friday) }
func (t Time) IsSaturday() bool  { return t.isSpecificWeekday(time.Saturday) }
func (t Time) IsWeekend() bool   { return t.IsSaturday() || t.IsSunday() }
func (t Time) IsWeekday() bool   { return !t.IsWeekend() }

func (t Time) IsNow() bool {
	if t.IsInvalid() {
		return false
	}
	return t.Timestamp() == t.Now().Timestamp()
}

func (t Time) IsFuture() bool {
	if t.IsInvalid() {
		return false
	}
	return t.Timestamp() > t.Now().Timestamp()
}

func (t Time) IsPast() bool {
	if t.IsInvalid() {
		return false
	}
	return t.Timestamp() < t.Now().Timestamp()
}

func (t Time) IsToday() bool {
	if t.IsInvalid() {
		return false
	}
	return t.ToDateString() == Now(t.ZoneName()).ToDateString()
}

func (t Time) IsTomorrow() bool {
	if t.IsInvalid() {
		return false
	}
	return t.ToDateString() == Tomorrow(t.ZoneName()).ToDateString()
}

func (t Time) IsYesterday() bool {
	if t.IsInvalid() {
		return false
	}
	return t.ToDateString() == Yesterday(t.ZoneName()).ToDateString()
}

func (t Time) IsSameCentury(rhs Time) bool {
	if t.IsInvalid() || rhs.IsInvalid() {
		return false
	}
	return t.Century() == rhs.Century()
}

func (t Time) IsSameDecade(rhs Time) bool {
	if t.IsInvalid() || rhs.IsInvalid() {
		return false
	}
	return t.Decade() == rhs.Decade()
}

func (t Time) IsSameYear(rhs Time) bool {
	if t.IsInvalid() || rhs.IsInvalid() {
		return false
	}
	return t.Year() == rhs.Year()
}

func (t Time) IsSameQuarter(rhs Time) bool {
	if t.IsInvalid() || rhs.IsInvalid() {
		return false
	}
	return t.Year() == rhs.Year() && t.Quarter() == rhs.Quarter()
}

func (t Time) IsSameMonth(rhs Time) bool {
	if t.IsInvalid() || rhs.IsInvalid() {
		return false
	}
	return t.Year() == rhs.Year() && t.Month() == rhs.Month()
}

func (t Time) IsSameDay(rhs Time) bool {
	if t.IsInvalid() || rhs.IsInvalid() {
		return false
	}
	return t.Year() == rhs.Year() && t.Month() == rhs.Month() && t.Day() == rhs.Day()
}

func (t Time) IsSameHour(rhs Time) bool {
	if t.IsInvalid() || rhs.IsInvalid() {
		return false
	}
	return t.IsSameDay(rhs) && t.Hour() == rhs.Hour()
}

func (t Time) IsSameMinute(rhs Time) bool {
	if t.IsInvalid() || rhs.IsInvalid() {
		return false
	}
	return t.IsSameDay(rhs) && t.Hour() == rhs.Hour() && t.Minute() == rhs.Minute()
}

func (t Time) IsSameSecond(rhs Time) bool {
	if t.IsInvalid() || rhs.IsInvalid() {
		return false
	}
	return t.IsSameDay(rhs) && t.Hour() == rhs.Hour() && t.Minute() == rhs.Minute() && t.Second() == rhs.Second()
}

func (t Time) Compare(operator string, rhs Time) bool {
	if t.IsInvalid() || rhs.IsInvalid() {
		return false
	}
	switch operator {
	case "==":
		return t.Eq(rhs)
	case "!=", "<>":
		return t.Ne(rhs)
	case ">":
		return t.Gt(rhs)
	case ">=":
		return t.Gte(rhs)
	case "<":
		return t.Lt(rhs)
	case "<=":
		return t.Lte(rhs)
	}
	return false
}

func (t Time) Eq(rhs Time) bool {
	if t.IsInvalid() || rhs.IsInvalid() {
		return false
	}
	return t.time.Equal(rhs.time)
}

func (t Time) Ne(rhs Time) bool {
	if t.IsInvalid() || rhs.IsInvalid() {
		return false
	}
	return !t.Eq(rhs)
}

func (t Time) Gt(rhs Time) bool {
	if t.IsInvalid() || rhs.IsInvalid() {
		return false
	}
	return t.time.After(rhs.time)
}

func (t Time) Gte(rhs Time) bool {
	if t.IsInvalid() || rhs.IsInvalid() {
		return false
	}
	return t.Gt(rhs) || t.Eq(rhs)
}

func (t Time) Lt(rhs Time) bool {
	if t.IsInvalid() || rhs.IsInvalid() {
		return false
	}
	return t.time.Before(rhs.time)
}

func (t Time) Lte(rhs Time) bool {
	if t.IsInvalid() || rhs.IsInvalid() {
		return false
	}
	return t.Lt(rhs) || t.Eq(rhs)
}

func (t Time) Between(start, end Time) bool {
	if t.IsInvalid() || start.IsInvalid() || end.IsInvalid() {
		return false
	}
	if start.Gt(end) {
		return false
	}
	return t.Gt(start) && t.Lt(end)
}

func (t Time) BetweenIncludedStart(start, end Time) bool {
	if t.IsInvalid() || start.IsInvalid() || end.IsInvalid() {
		return false
	}
	if start.Gt(end) {
		return false
	}
	return t.Gte(start) && t.Lt(end)
}

func (t Time) BetweenIncludedEnd(start, end Time) bool {
	if t.IsInvalid() || start.IsInvalid() || end.IsInvalid() {
		return false
	}
	if start.Gt(end) {
		return false
	}
	return t.Gt(start) && t.Lte(end)
}

func (t Time) BetweenIncludedBoth(start, end Time) bool {
	if t.IsInvalid() || start.IsInvalid() || end.IsInvalid() {
		return false
	}
	if start.Gt(end) {
		return false
	}
	return t.Gte(start) && t.Lte(end)
}
