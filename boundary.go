package chrono

func (t Time) StartOfCentury() Time {
	if t.IsInvalid() {
		return t
	}
	return t.create(t.Year()/YearsPerCentury*YearsPerCentury, 1, 1, 0, 0, 0, 0)
}

func (t Time) EndOfCentury() Time {
	if t.IsInvalid() {
		return t
	}
	return t.create(t.Year()/YearsPerCentury*YearsPerCentury+99, 12, 31, 23, 59, 59, 999999999)
}

func (t Time) StartOfDecade() Time {
	if t.IsInvalid() {
		return t
	}
	return t.create(t.Year()/YearsPerDecade*YearsPerDecade, 1, 1, 0, 0, 0, 0)
}

func (t Time) EndOfDecade() Time {
	if t.IsInvalid() {
		return t
	}
	return t.create(t.Year()/YearsPerDecade*YearsPerDecade+9, 12, 31, 23, 59, 59, 999999999)
}

func (t Time) StartOfYear() Time {
	if t.IsInvalid() {
		return t
	}
	return t.create(t.Year(), 1, 1, 0, 0, 0, 0)
}

func (t Time) EndOfYear() Time {
	if t.IsInvalid() {
		return t
	}
	return t.create(t.Year(), 12, 31, 23, 59, 59, 999999999)
}

func (t Time) StartOfQuarter() Time {
	if t.IsInvalid() {
		return t
	}
	year, quarter, day := t.Year(), t.Quarter(), 1
	return t.create(year, 3*quarter-2, day, 0, 0, 0, 0)
}

func (t Time) EndOfQuarter() Time {
	if t.IsInvalid() {
		return t
	}
	year, quarter, day := t.Year(), t.Quarter(), 30
	switch quarter {
	case 1, 4:
		day = 31
	case 2, 3:
		day = 30
	}
	return t.create(year, 3*quarter, day, 23, 59, 59, 999999999)
}

func (t Time) StartOfMonth() Time {
	if t.IsInvalid() {
		return t
	}
	year, month, _ := t.Date()
	return t.create(year, month, 1, 0, 0, 0, 0)
}

func (t Time) EndOfMonth() Time {
	if t.IsInvalid() {
		return t
	}
	year, month, _ := t.Date()
	return t.create(year, month+1, 0, 23, 59, 59, 999999999)
}

func (t Time) StartOfWeek() Time {
	if t.IsInvalid() {
		return t
	}
	dayOfWeek, weekStartsAt := t.DayOfWeek(), int(t.weekStartsAt)
	return t.SubDays((DaysPerWeek + dayOfWeek - weekStartsAt) % DaysPerWeek).StartOfDay()
}

func (t Time) EndOfWeek() Time {
	if t.IsInvalid() {
		return t
	}
	dayOfWeek, weekEndsAt := t.DayOfWeek(), int(t.weekStartsAt)+DaysPerWeek-1
	return t.AddDays((DaysPerWeek - dayOfWeek + weekEndsAt) % DaysPerWeek).EndOfDay()
}

func (t Time) StartOfDay() Time {
	if t.IsInvalid() {
		return t
	}
	year, month, day := t.Date()
	return t.create(year, month, day, 0, 0, 0, 0)
}

func (t Time) EndOfDay() Time {
	if t.IsInvalid() {
		return t
	}
	year, month, day := t.Date()
	return t.create(year, month, day, 23, 59, 59, 999999999)
}

func (t Time) StartOfHour() Time {
	if t.IsInvalid() {
		return t
	}
	year, month, day, hour, _, _ := t.DateTime()
	return t.create(year, month, day, hour, 0, 0, 0)
}

func (t Time) EndOfHour() Time {
	if t.IsInvalid() {
		return t
	}
	year, month, day, hour, _, _ := t.DateTime()
	return t.create(year, month, day, hour, 59, 59, 999999999)
}

func (t Time) StartOfMinute() Time {
	if t.IsInvalid() {
		return t
	}
	year, month, day, hour, minute, _ := t.DateTime()
	return t.create(year, month, day, hour, minute, 0, 0)
}

func (t Time) EndOfMinute() Time {
	if t.IsInvalid() {
		return t
	}
	year, month, day, hour, minute, _ := t.DateTime()
	return t.create(year, month, day, hour, minute, 59, 999999999)
}

func (t Time) StartOfSecond() Time {
	if t.IsInvalid() {
		return t
	}
	year, month, day, hour, minute, second := t.DateTime()
	return t.create(year, month, day, hour, minute, second, 0)
}

func (t Time) EndOfSecond() Time {
	if t.IsInvalid() {
		return t
	}
	year, month, day, hour, minute, second := t.DateTime()
	return t.create(year, month, day, hour, minute, second, 999999999)
}
