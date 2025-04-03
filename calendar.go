package chrono

import "github.com/AyakuraYuki/chrono/calendar/lunar"

func (t Time) Lunar() lunar.Lunar {
	l := lunar.Lunar{}
	if t.HasError() {
		l.Error = t.Error
		return l
	}
	return lunar.FromStdTime(t.StdTime())
}

func FromLunar(year, month, day int, isLeapMonth bool) Time {
	l := lunar.FromLunar(year, month, day, isLeapMonth)
	if !l.IsValid() {
		t := New()
		t.Error = l.Error
		return t
	}
	return New(l.ToGregorian(defaultTimezone).Time)
}
