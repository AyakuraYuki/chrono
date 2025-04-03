package chrono

import (
	"time"
)

// StdTime returns held standard time.Time
func (t Time) StdTime() time.Time {
	if t.IsZero() {
		return t.time
	}
	if t.loc == nil {
		return t.time
	}
	return t.time.In(t.loc)
}

func (t Time) DaysInYear() int {
	if t.HasError() {
		return 0
	}
	if t.IsLeapYear() {
		return DaysPerLeapYear
	}
	return DaysPerNormalYear
}

func (t Time) DaysInMonth() int {
	if t.HasError() {
		return 0
	}
	return t.EndOfMonth().StdTime().Day()
}

func (t Time) MonthOfYear() int {
	if t.HasError() {
		return 0
	}
	return int(t.StdTime().Month())
}

func (t Time) DayOfYear() int {
	if t.HasError() {
		return 0
	}
	return t.StdTime().YearDay()
}

func (t Time) DayOfMonth() int {
	if t.HasError() {
		return 0
	}
	return t.StdTime().Day()
}

func (t Time) DayOfWeek() int {
	if t.HasError() {
		return 0
	}
	day := t.StdTime().Weekday()
	if day == time.Sunday {
		return DaysPerWeek
	}
	return int(day)
}

func (t Time) WeekOfYear() int {
	if t.HasError() {
		return 0
	}
	_, week := t.StdTime().ISOWeek()
	return week
}

func (t Time) WeekOfMonth() int {
	if t.HasError() {
		return 0
	}
	days := t.Day() + t.StartOfMonth().DayOfWeek() - 1
	if days%DaysPerWeek == 0 {
		return days / DaysPerWeek
	}
	return days/DaysPerWeek + 1
}

func (t Time) DateTime() (year, month, day, hour, minute, second int) {
	if t.HasError() {
		return
	}
	year, month, day = t.Date()
	hour, minute, second = t.Time()
	return
}

func (t Time) DateTimeMilli() (year, month, day, hour, minute, second, millisecond int) {
	if t.HasError() {
		return
	}
	year, month, day, hour, minute, second = t.DateTime()
	millisecond = t.Millisecond()
	return
}

func (t Time) DateTimeMicro() (year, month, day, hour, minute, second, microsecond int) {
	if t.HasError() {
		return
	}
	year, month, day, hour, minute, second = t.DateTime()
	microsecond = t.Microsecond()
	return
}

func (t Time) DateTimeNano() (year, month, day, hour, minute, second, nanosecond int) {
	if t.HasError() {
		return
	}
	year, month, day, hour, minute, second = t.DateTime()
	nanosecond = t.Nanosecond()
	return
}

func (t Time) Date() (year, month, day int) {
	if t.HasError() {
		return
	}
	var tm time.Month
	year, tm, day = t.StdTime().Date()
	return year, int(tm), day
}

func (t Time) DateMilli() (year, month, day, millisecond int) {
	if t.HasError() {
		return
	}
	year, month, day = t.Date()
	millisecond = t.Millisecond()
	return
}

func (t Time) DateMicro() (year, month, day, microsecond int) {
	if t.HasError() {
		return
	}
	year, month, day = t.Date()
	microsecond = t.Microsecond()
	return
}

func (t Time) DateNano() (year, month, day, nanosecond int) {
	if t.HasError() {
		return
	}
	year, month, day = t.Date()
	nanosecond = t.Nanosecond()
	return
}

func (t Time) Time() (hour, minute, second int) {
	if t.HasError() {
		return
	}
	return t.StdTime().Clock()
}

func (t Time) TimeMilli() (hour, minute, second, millisecond int) {
	if t.HasError() {
		return
	}
	hour, minute, second = t.Time()
	millisecond = t.Millisecond()
	return
}

func (t Time) TimeMicro() (hour, minute, second, microsecond int) {
	if t.HasError() {
		return
	}
	hour, minute, second = t.Time()
	microsecond = t.Microsecond()
	return
}

func (t Time) TimeNano() (hour, minute, second, nanosecond int) {
	if t.HasError() {
		return
	}
	hour, minute, second = t.Time()
	nanosecond = t.Nanosecond()
	return
}

func (t Time) Century() int {
	if t.HasError() {
		return 0
	}
	return t.Year()/YearsPerCentury + 1
}

func (t Time) Decade() int {
	if t.HasError() {
		return 0
	}
	return t.Year() % YearsPerCentury / YearsPerDecade * YearsPerDecade
}

func (t Time) Year() int {
	if t.HasError() {
		return 0
	}
	return t.StdTime().Year()
}

func (t Time) Quarter() int {
	if t.HasError() {
		return 0
	}
	month := t.Month()
	switch {
	case month >= 10:
		return 4
	case month >= 7:
		return 3
	case month >= 4:
		return 2
	case month >= 1:
		return 1
	}
	return 0
}

func (t Time) Month() int {
	return t.MonthOfYear()
}

func (t Time) Week() int {
	if t.HasError() {
		return -1
	}
	return (t.DayOfWeek() + DaysPerWeek - int(t.weekStartsAt)) % DaysPerWeek
}

func (t Time) Day() int {
	return t.DayOfMonth()
}

func (t Time) Hour() int {
	if t.HasError() {
		return 0
	}
	return t.StdTime().Hour()
}

func (t Time) Minute() int {
	if t.HasError() {
		return 0
	}
	return t.StdTime().Minute()
}

func (t Time) Second() int {
	if t.HasError() {
		return 0
	}
	return t.StdTime().Second()
}

func (t Time) Millisecond() int {
	if t.HasError() {
		return 0
	}
	return t.StdTime().Nanosecond() / 1e6
}

func (t Time) Microsecond() int {
	if t.HasError() {
		return 0
	}
	return t.StdTime().Nanosecond() / 1e3
}

func (t Time) Nanosecond() int {
	if t.HasError() {
		return 0
	}
	return t.StdTime().Nanosecond()
}

func (t Time) Timestamp() int64 {
	if t.HasError() {
		return 0
	}
	return t.StdTime().Unix()
}

func (t Time) TimestampMilli() int64 {
	if t.HasError() {
		return 0
	}
	return t.StdTime().UnixMilli()
}

func (t Time) TimestampMicro() int64 {
	if t.HasError() {
		return 0
	}
	return t.StdTime().UnixMicro()
}

func (t Time) TimestampNano() int64 {
	if t.HasError() {
		return 0
	}
	return t.StdTime().UnixNano()
}

func (t Time) Timezone() string {
	if t.HasError() {
		return ""
	}
	return t.loc.String()
}

func (t Time) ZoneName() string {
	if t.HasError() {
		return ""
	}
	name, _ := t.StdTime().Zone()
	return name
}

func (t Time) ZoneOffset() int {
	if t.HasError() {
		return 0
	}
	_, offset := t.StdTime().Zone()
	return offset
}

func (t Time) Locale() string {
	if t.HasError() {
		return ""
	}
	return t.lang.locale
}

func (t Time) WeekStartsAt() string {
	if t.HasError() {
		return ""
	}
	return t.weekStartsAt.String()
}

func (t Time) CurrentLayout() string {
	if t.HasError() {
		return ""
	}
	return t.layout
}

func (t Time) Age() int {
	if t.HasError() {
		return 0
	}
	birth := t
	now := t.Now()
	if !now.IsLeapYear() && birth.Month() == 2 && birth.Day() == 29 {
		birth = birth.AddDay()
	}
	if now.Lt(birth) {
		return 0
	}
	years := now.Year() - birth.Year()
	if now.Month() < birth.Month() || (now.IsSameMonth(birth) && now.Day() < birth.Day()) {
		years--
	}
	return years
}

func (t Time) NominalAge() int {
	birth := t.SetTimezone(Shanghai)
	bl := FromLunar(birth.Year(), 1, 1, false).SetTimezone(Shanghai)
	now := birth.Now()
	nl := FromLunar(now.Year(), 1, 1, false).SetTimezone(Shanghai)
	years := now.Year() - birth.Year()
	if birth.Lt(bl) {
		years++
	}
	if now.Gte(nl) {
		years++
	}
	return years
}
