package chrono

import "time"

func SetLayout(layout string) Time        { return New().SetLayout(layout) }
func SetFormat(format string) Time        { return New().SetFormat(format) }
func SetWeekStartsAt(day string) Time     { return New().SetWeekStartsAt(day) }
func SetTimezone(timezone string) Time    { return New().SetTimezone(timezone) }
func SetLocation(loc *time.Location) Time { return New().SetLocation(loc) }
func SetLocale(locale string) Time        { return New().SetLocale(locale) }

func (t Time) SetLayout(layout string) Time {
	if layout == "" {
		t.Error = emptyLayoutError()
	}
	if t.HasError() {
		return t
	}
	t.layout = layout
	return t
}

func (t Time) SetFormat(format string) Time {
	if format == "" {
		t.Error = emptyFormatError()
	}
	if t.HasError() {
		return t
	}
	t.layout = format2layout(format)
	return t
}

func (t Time) SetWeekStartsAt(day string) Time {
	if day == "" {
		t.Error = emptyWeekStartsAtError()
	}
	if t.HasError() {
		return t
	}
	if weekday, ok := weekdays[day]; ok {
		t.weekStartsAt = weekday
	} else {
		t.Error = invalidWeekStartsAtError(day)
	}
	return t
}

func (t Time) SetLocation(loc *time.Location) Time {
	if loc == nil {
		t.Error = invalidLocationError()
	}
	if t.HasError() {
		return t
	}
	t.loc = loc
	return t
}

func (t Time) SetTimezone(timezone string) Time {
	if timezone == "" {
		t.Error = emptyTimezoneError()
	}
	if t.HasError() {
		return t
	}
	t.loc, t.Error = getLocationByTimezone(timezone)
	return t
}

func (t Time) SetLocale(locale string) Time {
	if locale == "" {
		t.Error = emptyLocaleError()
	}
	if t.HasError() {
		return t
	}
	t.lang.SetLocale(locale)
	t.Error = t.lang.Error
	return t
}

func (t Time) SetLanguage(lang *Language) Time {
	if lang == nil {
		t.Error = nilLanguageError()
		return t
	}
	if t.HasError() {
		return t
	}
	t.lang.dir = lang.dir
	t.lang.locale = lang.locale
	t.lang.resources = lang.resources
	t.lang.Error = lang.Error
	return t
}

func (t Time) SetDateTime(year, month, day, hour, minute, second int) Time {
	if t.HasError() {
		return t
	}
	return t.create(year, month, day, hour, minute, second, t.Nanosecond())
}

func (t Time) SetDateTimeMilli(year, month, day, hour, minute, second, millisecond int) Time {
	if t.HasError() {
		return t
	}
	return t.create(year, month, day, hour, minute, second, millisecond*1e6)
}

func (t Time) SetDateTimeMicro(year, month, day, hour, minute, second, microsecond int) Time {
	if t.HasError() {
		return t
	}
	return t.create(year, month, day, hour, minute, second, microsecond*1e3)
}

func (t Time) SetDateTimeNano(year, month, day, hour, minute, second, nanosecond int) Time {
	if t.HasError() {
		return t
	}
	return t.create(year, month, day, hour, minute, second, nanosecond)
}

func (t Time) SetDate(year, month, day int) Time {
	if t.HasError() {
		return t
	}
	hour, minute, second := t.Time()
	return t.create(year, month, day, hour, minute, second, t.Nanosecond())
}

func (t Time) SetDateMilli(year, month, day, millisecond int) Time {
	if t.HasError() {
		return t
	}
	hour, minute, second := t.Time()
	return t.create(year, month, day, hour, minute, second, millisecond*1e6)
}

func (t Time) SetDateMicro(year, month, day, microsecond int) Time {
	if t.HasError() {
		return t
	}
	hour, minute, second := t.Time()
	return t.create(year, month, day, hour, minute, second, microsecond*1e3)
}

func (t Time) SetDateNano(year, month, day, nanosecond int) Time {
	if t.HasError() {
		return t
	}
	hour, minute, second := t.Time()
	return t.create(year, month, day, hour, minute, second, nanosecond)
}

func (t Time) SetTime(hour, minute, second int) Time {
	if t.HasError() {
		return t
	}
	year, month, day := t.Date()
	return t.create(year, month, day, hour, minute, second, t.Nanosecond())
}

func (t Time) SetTimeMilli(hour, minute, second, millisecond int) Time {
	if t.HasError() {
		return t
	}
	year, month, day := t.Date()
	return t.create(year, month, day, hour, minute, second, millisecond*1e6)
}

func (t Time) SetTimeMicro(hour, minute, second, microsecond int) Time {
	if t.HasError() {
		return t
	}
	year, month, day := t.Date()
	return t.create(year, month, day, hour, minute, second, microsecond*1e3)
}

func (t Time) SetTimeNano(hour, minute, second, nanosecond int) Time {
	if t.HasError() {
		return t
	}
	year, month, day := t.Date()
	return t.create(year, month, day, hour, minute, second, nanosecond)
}

func (t Time) SetYear(year int) Time {
	if t.HasError() {
		return t
	}
	_, month, day, hour, minute, second := t.DateTime()
	return t.create(year, month, day, hour, minute, second, t.Nanosecond())
}

func (t Time) SetYearNoOverflow(year int) Time {
	if t.HasError() {
		return t
	}
	return t.AddYearsNoOverflow(year - t.Year())
}

func (t Time) SetMonth(month int) Time {
	if t.HasError() {
		return t
	}
	year, _, day, hour, minute, second := t.DateTime()
	return t.create(year, month, day, hour, minute, second, t.Nanosecond())
}

func (t Time) SetMonthNoOverflow(month int) Time {
	if t.HasError() {
		return t
	}
	return t.AddMonthsNoOverflow(month - t.Month())
}

func (t Time) SetDay(day int) Time {
	if t.HasError() {
		return t
	}
	year, month, _, hour, minute, second := t.DateTime()
	return t.create(year, month, day, hour, minute, second, t.Nanosecond())
}

func (t Time) SetHour(hour int) Time {
	if t.HasError() {
		return t
	}
	year, month, day, _, minute, second := t.DateTime()
	return t.create(year, month, day, hour, minute, second, t.Nanosecond())
}

func (t Time) SetMinute(minute int) Time {
	if t.HasError() {
		return t
	}
	year, month, day, hour, _, second := t.DateTime()
	return t.create(year, month, day, hour, minute, second, t.Nanosecond())
}

func (t Time) SetSecond(second int) Time {
	if t.HasError() {
		return t
	}
	year, month, day, hour, minute, _ := t.DateTime()
	return t.create(year, month, day, hour, minute, second, t.Nanosecond())
}

func (t Time) SetMillisecond(millisecond int) Time {
	if t.HasError() {
		return t
	}
	year, month, day, hour, minute, second := t.DateTime()
	return t.create(year, month, day, hour, minute, second, millisecond*1e6)
}

func (t Time) SetMicrosecond(microsecond int) Time {
	if t.HasError() {
		return t
	}
	year, month, day, hour, minute, second := t.DateTime()
	return t.create(year, month, day, hour, minute, second, microsecond*1e3)
}

func (t Time) SetNanosecond(nanosecond int) Time {
	if t.HasError() {
		return t
	}
	year, month, day, hour, minute, second := t.DateTime()
	return t.create(year, month, day, hour, minute, second, nanosecond)
}
