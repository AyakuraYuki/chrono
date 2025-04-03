package chrono

import "time"

// FromStdTime creates a Time instance from standard time.Time
func FromStdTime(from time.Time, timezone ...string) Time {
	t := New(from)
	t.loc = from.Location()
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	t.time = from
	return t
}

// ----------------------------------------------------------------------------------------------------

// FromTimestamp creates a Time instance from unix timestamp
func (t Time) FromTimestamp(timestamp int64, timezone ...string) Time {
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.HasError() {
		return t
	}
	t.time = time.Unix(timestamp, 0)
	return t
}

// FromTimestamp creates a Time instance from unix timestamp
func FromTimestamp(timestamp int64, timezone ...string) Time {
	return New().FromTimestamp(timestamp, timezone...)
}

// FromTimestampMilli creates a Time instance from timestamp with milliseconds
func (t Time) FromTimestampMilli(timestampMilli int64, timezone ...string) Time {
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.HasError() {
		return t
	}
	t.time = time.UnixMilli(timestampMilli)
	return t
}

// FromTimestampMilli creates a Time instance from timestamp with milliseconds
func FromTimestampMilli(timestampMilli int64, timezone ...string) Time {
	return New().FromTimestampMilli(timestampMilli, timezone...)
}

// FromTimestampMicro creates a Time instance from timestamp with microseconds
func (t Time) FromTimestampMicro(timestampMicro int64, timezone ...string) Time {
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.HasError() {
		return t
	}
	t.time = time.UnixMicro(timestampMicro)
	return t
}

// FromTimestampMicro creates a Time instance from timestamp with microseconds
func FromTimestampMicro(timestampMicro int64, timezone ...string) Time {
	return New().FromTimestampMicro(timestampMicro, timezone...)
}

// FromTimestampNano creates a Time instance from timestamp with nanoseconds
func (t Time) FromTimestampNano(timestampNano int64, timezone ...string) Time {
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.HasError() {
		return t
	}
	t.time = time.Unix(timestampNano/1e9, timestampNano%1e9)
	return t
}

// FromTimestampNano creates a Time instance from timestamp with nanoseconds
func FromTimestampNano(timestampNano int64, timezone ...string) Time {
	return New().FromTimestampNano(timestampNano, timezone...)
}

// ----------------------------------------------------------------------------------------------------

// FromDateTime creates a Time instance from a given date and clock
func (t Time) FromDateTime(year, month, day, hour, minute, second int, timezone ...string) Time {
	return t.create(year, month, day, hour, minute, second, 0, timezone...)
}

// FromDateTime creates a Time instance from a given date and clock
func FromDateTime(year, month, day, hour, minute, second int, timezone ...string) Time {
	return New().FromDateTime(year, month, day, hour, minute, second, timezone...)
}

// FromDateTimeMilli creates a Time instance from a given date, clock and millisecond
func (t Time) FromDateTimeMilli(year, month, day, hour, minute, second, millisecond int, timezone ...string) Time {
	return t.create(year, month, day, hour, minute, second, millisecond*1e6, timezone...)
}

// FromDateTimeMilli creates a Time instance from a given date, clock and millisecond
func FromDateTimeMilli(year, month, day, hour, minute, second, millisecond int, timezone ...string) Time {
	return New().FromDateTimeMilli(year, month, day, hour, minute, second, millisecond, timezone...)
}

// FromDateTimeMicro creates a Time instance from a given date, clock and microsecond
func (t Time) FromDateTimeMicro(year, month, day, hour, minute, second, microsecond int, timezone ...string) Time {
	return t.create(year, month, day, hour, minute, second, microsecond*1e3, timezone...)
}

// FromDateTimeMicro creates a Time instance from a given date, clock and microsecond
func FromDateTimeMicro(year, month, day, hour, minute, second, microsecond int, timezone ...string) Time {
	return New().FromDateTimeMicro(year, month, day, hour, minute, second, microsecond, timezone...)
}

// FromDateTimeNano creates a Time instance from a given date, clock and nanosecond
func (t Time) FromDateTimeNano(year, month, day, hour, minute, second, nanosecond int, timezone ...string) Time {
	return t.create(year, month, day, hour, minute, second, nanosecond, timezone...)
}

// FromDateTimeNano creates a Time instance from a given date, clock and nanosecond
func FromDateTimeNano(year, month, day, hour, minute, second, nanosecond int, timezone ...string) Time {
	return New().FromDateTimeNano(year, month, day, hour, minute, second, nanosecond, timezone...)
}

// ----------------------------------------------------------------------------------------------------

// FromDate creates a Time instance from a given date
func (t Time) FromDate(year, month, day int, timezone ...string) Time {
	return t.create(year, month, day, 0, 0, 0, 0, timezone...)
}

// FromDate creates a Time instance from a given date
func FromDate(year, month, day int, timezone ...string) Time {
	return New().FromDate(year, month, day, timezone...)
}

// FromDateMilli creates a Time instance from a given date and millisecond
func (t Time) FromDateMilli(year, month, day, millisecond int, timezone ...string) Time {
	return t.create(year, month, day, 0, 0, 0, millisecond*1e6, timezone...)
}

// FromDateMilli creates a Time instance from a given date and millisecond
func FromDateMilli(year, month, day, millisecond int, timezone ...string) Time {
	return New().FromDateMilli(year, month, day, millisecond, timezone...)
}

// FromDateMicro creates a Time instance from a given date and microsecond
func (t Time) FromDateMicro(year, month, day, microsecond int, timezone ...string) Time {
	return t.create(year, month, day, 0, 0, 0, microsecond*1e3, timezone...)
}

// FromDateMicro creates a Time instance from a given date and microsecond
func FromDateMicro(year, month, day, microsecond int, timezone ...string) Time {
	return New().FromDateMicro(year, month, day, microsecond, timezone...)
}

// FromDateNano creates a Time instance from a given date and nanosecond
func (t Time) FromDateNano(year, month, day, nanosecond int, timezone ...string) Time {
	return t.create(year, month, day, 0, 0, 0, nanosecond, timezone...)
}

// FromDateNano creates a Time instance from a given date and nanosecond
func FromDateNano(year, month, day, nanosecond int, timezone ...string) Time {
	return New().FromDateNano(year, month, day, nanosecond, timezone...)
}

// ----------------------------------------------------------------------------------------------------

// FromTime creates a Time instance from a given clock
func (t Time) FromTime(hour, minute, second int, timezone ...string) Time {
	year, month, day := t.Now(timezone...).Date()
	return t.create(year, month, day, hour, minute, second, 0, timezone...)
}

// FromTime creates a Time instance from a given clock
func FromTime(hour, minute, second int, timezone ...string) Time {
	return New().FromTime(hour, minute, second, timezone...)
}

// FromTimeMilli creates a Time instance from a given clock and millisecond
func (t Time) FromTimeMilli(hour, minute, second, millisecond int, timezone ...string) Time {
	year, month, day := t.Now(timezone...).Date()
	return t.create(year, month, day, hour, minute, second, millisecond*1e6, timezone...)
}

// FromTimeMilli creates a Time instance from a given clock and millisecond
func FromTimeMilli(hour, minute, second, millisecond int, timezone ...string) Time {
	return New().FromTimeMilli(hour, minute, second, millisecond, timezone...)
}

// FromTimeMicro creates a Time instance from a given clock and microsecond
func (t Time) FromTimeMicro(hour, minute, second, microsecond int, timezone ...string) Time {
	year, month, day := t.Now(timezone...).Date()
	return t.create(year, month, day, hour, minute, second, microsecond*1e3, timezone...)
}

// FromTimeMicro creates a Time instance from a given clock and microsecond
func FromTimeMicro(hour, minute, second, microsecond int, timezone ...string) Time {
	return New().FromTimeMicro(hour, minute, second, microsecond, timezone...)
}

// FromTimeNano creates a Time instance from a given clock and nanosecond
func (t Time) FromTimeNano(hour, minute, second, nanosecond int, timezone ...string) Time {
	year, month, day := t.Now(timezone...).Date()
	return t.create(year, month, day, hour, minute, second, nanosecond, timezone...)
}

// FromTimeNano creates a Time instance from a given clock and nanosecond
func FromTimeNano(hour, minute, second, nanosecond int, timezone ...string) Time {
	return New().FromTimeNano(hour, minute, second, nanosecond, timezone...)
}

// ----------------------------------------------------------------------------------------------------

// create a Time instance from a given date, clock and nanosecond
func (t Time) create(year, month, day, hour, minute, second, nanosecond int, timezone ...string) Time {
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.HasError() {
		return t
	}
	t.time = time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, t.loc)
	return t
}
