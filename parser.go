package chrono

import (
	"strconv"
	"time"
)

// Parse time string as a Time instance
func (t Time) Parse(value string, timezone ...string) Time {
	if value == "" {
		t.Error = emptyValueError()
		return t
	}
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.HasError() {
		return t
	}
	switch value {
	case "now", "Now", "NOW", "today", "Today", "TODAY":
		return t.Now(timezone...)
	case "yesterday", "Yesterday", "YESTERDAY":
		return t.Yesterday(timezone...)
	case "tomorrow", "Tomorrow", "TOMORROW":
		return t.Tomorrow(timezone...)
	}
	for _, layout := range defaultLayouts {
		if tt, err := time.ParseInLocation(layout, value, t.loc); err == nil {
			t.time = tt
			t.layout = layout
			return t
		}
	}
	t.Error = failedParseError(value)
	return t
}

// Parse time string as a Time instance
func Parse(value string, timezone ...string) Time {
	return New().Parse(value, timezone...)
}

// ParseByFormat parses time string as a Time instance by chrono format
func (t Time) ParseByFormat(value, format string, timezone ...string) Time {
	if format == "" {
		t.Error = emptyFormatError()
		return t
	}
	tt := t.ParseByLayout(value, format2layout(format), timezone...)
	if tt.HasError() {
		tt.Error = invalidFormatError(value, format)
	}
	return tt
}

// ParseByFormat parses time string as a Time instance by chrono format
func ParseByFormat(value, format string, timezone ...string) Time {
	return New().ParseByFormat(value, format, timezone...)
}

// ParseByLayout parses time string as a Time instance by layout
func (t Time) ParseByLayout(value, layout string, timezone ...string) Time {
	if value == "" {
		t.Error = emptyValueError()
		return t
	}
	if layout == "" {
		t.Error = emptyLayoutError()
		return t
	}
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.HasError() {
		return t
	}
	if layout == TimestampLayout {
		ts, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			t.Error = invalidTimestampError(value)
			return t
		}
		return t.FromTimestamp(ts, t.Timezone())
	}
	if layout == TimestampMilliLayout {
		ts, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			t.Error = invalidTimestampError(value)
			return t
		}
		return t.FromTimestampMilli(ts, t.Timezone())
	}
	if layout == TimestampMicroLayout {
		ts, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			t.Error = invalidTimestampError(value)
			return t
		}
		return t.FromTimestampMicro(ts, t.Timezone())
	}
	if layout == TimestampNanoLayout {
		ts, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			t.Error = invalidTimestampError(value)
			return t
		}
		return t.FromTimestampNano(ts, t.Timezone())
	}
	tt, err := time.ParseInLocation(layout, value, t.loc)
	if err != nil {
		t.Error = invalidLayoutError(value, layout)
		return t
	}
	t.time = tt
	t.layout = layout
	return t
}

// ParseByLayout parses time string as a Time instance by layout
func ParseByLayout(value, layout string, timezone ...string) Time {
	return New().ParseByLayout(value, layout, timezone...)
}

// ParseWithLayouts parses time string as a Time instance with multiple layout options
func (t Time) ParseWithLayouts(value string, layouts []string, timezone ...string) Time {
	if value == "" {
		t.Error = emptyValueError()
		return t
	}
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.HasError() {
		return t
	}
	if len(layouts) == 0 {
		return t.Parse(value, timezone...)
	}
	for _, layout := range layouts {
		if tt := t.ParseByLayout(value, layout, timezone...); tt.IsValid() {
			return tt
		}
	}
	t.Error = failedParseError(value)
	return t
}

// ParseWithLayouts parses time string as a Time instance with multiple layout options
func ParseWithLayouts(value string, layouts []string, timezone ...string) Time {
	return New().ParseWithLayouts(value, layouts, timezone...)
}

// ParseWithFormats parses time string as a Time instance with multiple format options
func (t Time) ParseWithFormats(value string, formats []string, timezone ...string) Time {
	if value == "" {
		t.Error = emptyValueError()
		return t
	}
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.HasError() {
		return t
	}
	if len(formats) == 0 {
		return t.Parse(value, timezone...)
	}
	var layouts []string
	for _, format := range formats {
		layouts = append(layouts, format2layout(format))
	}
	return t.ParseWithLayouts(value, layouts, timezone...)
}

// ParseWithFormats parses time string as a Time instance with multiple format options
func ParseWithFormats(value string, formats []string, timezone ...string) Time {
	return New().ParseWithFormats(value, formats, timezone...)
}
