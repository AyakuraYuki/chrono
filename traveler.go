package chrono

import "time"

// Now returns a Time instance for now
func (t Time) Now(timezone ...string) Time {
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.HasError() {
		return t
	}
	t.time = time.Now().In(t.loc)
	return t
}

// Now returns a Time instance for now
func Now(timezone ...string) Time {
	return New().Now(timezone...)
}

// Tomorrow returns a Time instance for tomorrow
func (t Time) Tomorrow(timezone ...string) Time {
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.HasError() {
		return t
	}
	if t.IsZero() {
		return t.Now().AddDay()
	}
	return t.AddDay()
}

// Tomorrow returns a Time instance for tomorrow
func Tomorrow(timezone ...string) Time {
	return New().Tomorrow(timezone...)
}

// Yesterday returns a Time instance for yesterday
func (t Time) Yesterday(timezone ...string) Time {
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.HasError() {
		return t
	}
	if t.IsZero() {
		return t.Now().SubDay()
	}
	return t.SubDay()
}

// Yesterday returns a Time instance for yesterday
func Yesterday(timezone ...string) Time {
	return New().Yesterday(timezone...)
}

// region duration

func (t Time) AddDuration(duration string) Time {
	if t.IsInvalid() {
		return t
	}
	dur, err := parseByDuration(duration)
	t.time, t.Error = t.StdTime().Add(dur), err
	return t
}

func (t Time) SubDuration(duration string) Time {
	return t.AddDuration("-" + duration)
}

// endregion

// region centuries

func (t Time) AddCenturies(centuries int) Time { return t.AddYears(centuries * YearsPerCentury) }
func (t Time) AddCentury() Time                { return t.AddCenturies(1) }
func (t Time) SubCenturies(centuries int) Time { return t.SubYears(centuries * YearsPerCentury) }
func (t Time) SubCentury() Time                { return t.SubCenturies(1) }

func (t Time) AddCenturiesNoOverflow(centuries int) Time {
	return t.AddYearsNoOverflow(centuries * YearsPerCentury)
}
func (t Time) AddCenturyNoOverflow() Time { return t.AddCenturiesNoOverflow(1) }
func (t Time) SubCenturiesNoOverflow(centuries int) Time {
	return t.SubYearsNoOverflow(centuries * YearsPerCentury)
}
func (t Time) SubCenturyNoOverflow() Time { return t.SubCenturiesNoOverflow(1) }

// endregion

// region decades

func (t Time) AddDecades(decades int) Time { return t.AddYears(decades * YearsPerDecade) }
func (t Time) AddDecade() Time             { return t.AddDecades(1) }
func (t Time) SubDecades(decades int) Time { return t.SubYears(decades * YearsPerDecade) }
func (t Time) SubDecade() Time             { return t.SubDecades(1) }

func (t Time) AddDecadesNoOverflow(decades int) Time {
	return t.AddYearsNoOverflow(decades * YearsPerDecade)
}
func (t Time) AddDecadeNoOverflow() Time { return t.AddDecadesNoOverflow(1) }
func (t Time) SubDecadesNoOverflow(decades int) Time {
	return t.SubYearsNoOverflow(decades * YearsPerDecade)
}
func (t Time) SubDecadeNoOverflow() Time { return t.SubDecadesNoOverflow(1) }

// endregion

// region years

func (t Time) AddYears(years int) Time {
	if t.IsInvalid() {
		return t
	}
	t.time = t.StdTime().AddDate(years, 0, 0)
	return t
}
func (t Time) AddYear() Time           { return t.AddYears(1) }
func (t Time) SubYears(years int) Time { return t.AddYears(-years) }
func (t Time) SubYear() Time           { return t.SubYears(1) }

func (t Time) AddYearsNoOverflow(years int) Time {
	if t.IsInvalid() {
		return t
	}
	nanosecond := t.Nanosecond()
	year, month, day, hour, minute, second := t.DateTime()
	lastYear, lastMonth, lastDay := t.create(year+years, month+1, 0, hour, minute, second, nanosecond).Date()
	if day > lastDay {
		day = lastDay
	}
	return t.create(lastYear, lastMonth, day, hour, minute, second, nanosecond)
}
func (t Time) AddYearNoOverflow() Time           { return t.AddYearsNoOverflow(1) }
func (t Time) SubYearsNoOverflow(years int) Time { return t.AddYearsNoOverflow(-years) }
func (t Time) SubYearNoOverflow() Time           { return t.SubYearsNoOverflow(1) }

// endregion

// region quarters

func (t Time) AddQuarters(quarters int) Time { return t.AddMonths(quarters * MonthsPerQuarter) }
func (t Time) AddQuarter() Time              { return t.AddQuarters(1) }
func (t Time) SubQuarters(quarters int) Time { return t.AddQuarters(-quarters) }
func (t Time) SubQuarter() Time              { return t.SubQuarters(1) }

func (t Time) AddQuartersNoOverflow(quarters int) Time {
	return t.AddMonthsNoOverflow(quarters * MonthsPerQuarter)
}
func (t Time) AddQuarterNoOverflow() Time { return t.AddQuartersNoOverflow(1) }
func (t Time) SubQuartersNoOverflow(quarters int) Time {
	return t.SubMonthsNoOverflow(quarters * MonthsPerQuarter)
}
func (t Time) SubQuarterNoOverflow() Time { return t.SubQuartersNoOverflow(1) }

// endregion

// region months

func (t Time) AddMonths(months int) Time {
	if t.IsInvalid() {
		return t
	}
	t.time = t.StdTime().AddDate(0, months, 0)
	return t
}
func (t Time) AddMonth() Time            { return t.AddMonths(1) }
func (t Time) SubMonths(months int) Time { return t.AddMonths(-months) }
func (t Time) SubMonth() Time            { return t.SubMonths(1) }

func (t Time) AddMonthsNoOverflow(months int) Time {
	if t.IsInvalid() {
		return t
	}
	nanosecond := t.Nanosecond()
	year, month, day, hour, minute, second := t.DateTime()
	lastYear, lastMonth, lastDay := t.create(year, month+months+1, 0, hour, minute, second, nanosecond).Date()
	if day > lastDay {
		day = lastDay
	}
	return t.create(lastYear, lastMonth, day, hour, minute, second, nanosecond)
}
func (t Time) AddMonthNoOverflow() Time            { return t.AddMonthsNoOverflow(1) }
func (t Time) SubMonthsNoOverflow(months int) Time { return t.AddMonthsNoOverflow(-months) }
func (t Time) SubMonthNoOverflow() Time            { return t.SubMonthsNoOverflow(1) }

// endregion

// region weeks

func (t Time) AddWeeks(weeks int) Time { return t.AddDays(weeks * DaysPerWeek) }
func (t Time) AddWeek() Time           { return t.AddWeeks(1) }
func (t Time) SubWeeks(weeks int) Time { return t.SubDays(weeks * DaysPerWeek) }
func (t Time) SubWeek() Time           { return t.SubWeeks(1) }

// endregion

// region days

func (t Time) AddDays(days int) Time {
	if t.IsInvalid() {
		return t
	}
	t.time = t.StdTime().AddDate(0, 0, days)
	return t
}
func (t Time) AddDay() Time          { return t.AddDays(1) }
func (t Time) SubDays(days int) Time { return t.AddDays(-days) }
func (t Time) SubDay() Time          { return t.SubDays(1) }

// endregion

// region hours

func (t Time) AddHours(hours int) Time {
	if t.IsInvalid() {
		return t
	}
	dur := time.Duration(hours) * time.Hour
	t.time = t.StdTime().Add(dur)
	return t
}
func (t Time) AddHour() Time           { return t.AddHours(1) }
func (t Time) SubHours(hours int) Time { return t.AddHours(-hours) }
func (t Time) SubHour() Time           { return t.SubHours(1) }

// endregion

// region minutes

func (t Time) AddMinutes(minutes int) Time {
	if t.IsInvalid() {
		return t
	}
	dur := time.Duration(minutes) * time.Minute
	t.time = t.StdTime().Add(dur)
	return t
}
func (t Time) AddMinute() Time             { return t.AddMinutes(1) }
func (t Time) SubMinutes(minutes int) Time { return t.AddMinutes(-minutes) }
func (t Time) SubMinute() Time             { return t.SubMinutes(1) }

// endregion

// region seconds

func (t Time) AddSeconds(seconds int) Time {
	if t.IsInvalid() {
		return t
	}
	dur := time.Duration(seconds) * time.Second
	t.time = t.StdTime().Add(dur)
	return t
}
func (t Time) AddSecond() Time             { return t.AddSeconds(1) }
func (t Time) SubSeconds(seconds int) Time { return t.AddSeconds(-seconds) }
func (t Time) SubSecond() Time             { return t.SubSeconds(1) }

// endregion

// region milliseconds

func (t Time) AddMilliseconds(milliseconds int) Time {
	if t.IsInvalid() {
		return t
	}
	dur := time.Duration(milliseconds) * time.Millisecond
	t.time = t.StdTime().Add(dur)
	return t
}
func (t Time) AddMillisecond() Time                  { return t.AddMilliseconds(1) }
func (t Time) SubMilliseconds(milliseconds int) Time { return t.AddMilliseconds(-milliseconds) }
func (t Time) SubMillisecond() Time                  { return t.SubMilliseconds(1) }

// endregion

// region microseconds

func (t Time) AddMicroseconds(microseconds int) Time {
	if t.IsInvalid() {
		return t
	}
	dur := time.Duration(microseconds) * time.Microsecond
	t.time = t.StdTime().Add(dur)
	return t
}
func (t Time) AddMicrosecond() Time                  { return t.AddMicroseconds(1) }
func (t Time) SubMicroseconds(microseconds int) Time { return t.AddMicroseconds(-microseconds) }
func (t Time) SubMicrosecond() Time                  { return t.SubMicroseconds(1) }

// endregion

// region nanoseconds

func (t Time) AddNanoseconds(nanoseconds int) Time {
	if t.IsInvalid() {
		return t
	}
	dur := time.Duration(nanoseconds) * time.Nanosecond
	t.time = t.StdTime().Add(dur)
	return t
}
func (t Time) AddNanosecond() Time                 { return t.AddNanoseconds(1) }
func (t Time) SubNanoseconds(nanoseconds int) Time { return t.AddNanoseconds(-nanoseconds) }
func (t Time) SubNanosecond() Time                 { return t.SubNanoseconds(1) }

// endregion
