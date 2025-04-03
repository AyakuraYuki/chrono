package chrono

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func (t Time) String() string {
	return t.Layout(t.layout, t.Timezone())
}

func (t Time) GoString() string {
	if t.IsInvalid() {
		return ""
	}
	return t.StdTime().GoString()
}

func (t Time) ToString(timezone ...string) string {
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.IsInvalid() {
		return ""
	}
	return t.StdTime().String()
}

func (t Time) ToMonthString(timezone ...string) string {
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.IsInvalid() {
		return ""
	}
	if len(t.lang.resources) == 0 {
		t.lang.SetLocale(defaultLocale)
	}
	t.lang.rw.Lock()
	defer t.lang.rw.Unlock()
	if resources, ok := t.lang.resources["months"]; ok {
		words := strings.Split(resources, "|")
		if len(words) == MonthsPerYear {
			return words[t.Month()-1]
		}
	}
	return ""
}

func (t Time) ToShortMonthString(timezone ...string) string {
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.IsInvalid() {
		return ""
	}
	if len(t.lang.resources) == 0 {
		t.lang.SetLocale(defaultLocale)
	}
	t.lang.rw.Lock()
	defer t.lang.rw.Unlock()
	if resources, ok := t.lang.resources["short_months"]; ok {
		words := strings.Split(resources, "|")
		if len(words) == MonthsPerYear {
			return words[t.Month()-1]
		}
	}
	return ""
}

func (t Time) ToWeekString(timezone ...string) string {
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.IsInvalid() {
		return ""
	}
	if len(t.lang.resources) == 0 {
		t.lang.SetLocale(defaultLocale)
	}
	t.lang.rw.Lock()
	defer t.lang.rw.Unlock()
	if resources, ok := t.lang.resources["weeks"]; ok {
		words := strings.Split(resources, "|")
		if len(words) == DaysPerWeek {
			return words[t.DayOfWeek()%DaysPerWeek]
		}
	}
	return ""
}

func (t Time) ToShortWeekString(timezone ...string) string {
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.IsInvalid() {
		return ""
	}
	if len(t.lang.resources) == 0 {
		t.lang.SetLocale(defaultLocale)
	}
	t.lang.rw.Lock()
	defer t.lang.rw.Unlock()
	if resources, ok := t.lang.resources["short_weeks"]; ok {
		words := strings.Split(resources, "|")
		if len(words) == DaysPerWeek {
			return words[t.DayOfWeek()%DaysPerWeek]
		}
	}
	return ""
}

// ----------------------------------------------------------------------------------------------------

func (t Time) stdTimeFormatLayout(layout string, timezone ...string) string {
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.IsInvalid() {
		return ""
	}
	return t.StdTime().Format(layout)
}

func (t Time) ToDayDateTimeString(timezone ...string) string {
	return t.stdTimeFormatLayout(DayDateTimeLayout, timezone...)
}

func (t Time) ToDateTimeString(timezone ...string) string {
	return t.stdTimeFormatLayout(DateTimeLayout, timezone...)
}

func (t Time) ToDateTimeMilliString(timezone ...string) string {
	return t.stdTimeFormatLayout(DateTimeMilliLayout, timezone...)
}

func (t Time) ToDateTimeMicroString(timezone ...string) string {
	return t.stdTimeFormatLayout(DateTimeMicroLayout, timezone...)
}

func (t Time) ToDateTimeNanoString(timezone ...string) string {
	return t.stdTimeFormatLayout(DateTimeNanoLayout, timezone...)
}

func (t Time) ToShortDateTimeString(timezone ...string) string {
	return t.stdTimeFormatLayout(ShortDateTimeLayout, timezone...)
}

func (t Time) ToShortDateTimeMilliString(timezone ...string) string {
	return t.stdTimeFormatLayout(ShortDateTimeMilliLayout, timezone...)
}

func (t Time) ToShortDateTimeMicroString(timezone ...string) string {
	return t.stdTimeFormatLayout(ShortDateTimeMicroLayout, timezone...)
}

func (t Time) ToShortDateTimeNanoString(timezone ...string) string {
	return t.stdTimeFormatLayout(ShortDateTimeNanoLayout, timezone...)
}

func (t Time) ToDateString(timezone ...string) string {
	return t.stdTimeFormatLayout(DateLayout, timezone...)
}

func (t Time) ToDateMilliString(timezone ...string) string {
	return t.stdTimeFormatLayout(DateMilliLayout, timezone...)
}

func (t Time) ToDateMicroString(timezone ...string) string {
	return t.stdTimeFormatLayout(DateMicroLayout, timezone...)
}

func (t Time) ToDateNanoString(timezone ...string) string {
	return t.stdTimeFormatLayout(DateNanoLayout, timezone...)
}

func (t Time) ToShortDateString(timezone ...string) string {
	return t.stdTimeFormatLayout(ShortDateLayout, timezone...)
}

func (t Time) ToShortDateMilliString(timezone ...string) string {
	return t.stdTimeFormatLayout(ShortDateMilliLayout, timezone...)
}

func (t Time) ToShortDateMicroString(timezone ...string) string {
	return t.stdTimeFormatLayout(ShortDateMicroLayout, timezone...)
}

func (t Time) ToShortDateNanoString(timezone ...string) string {
	return t.stdTimeFormatLayout(ShortDateNanoLayout, timezone...)
}

func (t Time) ToTimeString(timezone ...string) string {
	return t.stdTimeFormatLayout(TimeLayout, timezone...)
}

func (t Time) ToTimeMilliString(timezone ...string) string {
	return t.stdTimeFormatLayout(TimeMilliLayout, timezone...)
}

func (t Time) ToTimeMicroString(timezone ...string) string {
	return t.stdTimeFormatLayout(TimeMicroLayout, timezone...)
}

func (t Time) ToTimeNanoString(timezone ...string) string {
	return t.stdTimeFormatLayout(TimeNanoLayout, timezone...)
}

func (t Time) ToShortTimeString(timezone ...string) string {
	return t.stdTimeFormatLayout(ShortTimeLayout, timezone...)
}

func (t Time) ToShortTimeMilliString(timezone ...string) string {
	return t.stdTimeFormatLayout(ShortTimeMilliLayout, timezone...)
}

func (t Time) ToShortTimeMicroString(timezone ...string) string {
	return t.stdTimeFormatLayout(ShortTimeMicroLayout, timezone...)
}

func (t Time) ToShortTimeNanoString(timezone ...string) string {
	return t.stdTimeFormatLayout(ShortTimeNanoLayout, timezone...)
}

func (t Time) ToAtomString(timezone ...string) string {
	return t.stdTimeFormatLayout(AtomLayout, timezone...)
}

func (t Time) ToANSICString(timezone ...string) string {
	return t.stdTimeFormatLayout(ANSICLayout, timezone...)
}

func (t Time) ToCookieString(timezone ...string) string {
	return t.stdTimeFormatLayout(CookieLayout, timezone...)
}

func (t Time) ToKitchenString(timezone ...string) string {
	return t.stdTimeFormatLayout(KitchenLayout, timezone...)
}

func (t Time) ToRssString(timezone ...string) string {
	return t.stdTimeFormatLayout(RssLayout, timezone...)
}

func (t Time) ToRubyDateString(timezone ...string) string {
	return t.stdTimeFormatLayout(RubyDateLayout, timezone...)
}

func (t Time) ToUnixDateString(timezone ...string) string {
	return t.stdTimeFormatLayout(UnixDateLayout, timezone...)
}

func (t Time) ToW3cString(timezone ...string) string {
	return t.stdTimeFormatLayout(W3cLayout, timezone...)
}

func (t Time) ToISO8601String(timezone ...string) string {
	return t.stdTimeFormatLayout(ISO8601Layout, timezone...)
}

func (t Time) ToISO8601MilliString(timezone ...string) string {
	return t.stdTimeFormatLayout(ISO8601MilliLayout, timezone...)
}

func (t Time) ToISO8601MicroString(timezone ...string) string {
	return t.stdTimeFormatLayout(ISO8601MicroLayout, timezone...)
}

func (t Time) ToISO8601NanoString(timezone ...string) string {
	return t.stdTimeFormatLayout(ISO8601NanoLayout, timezone...)
}

func (t Time) ToISO8601ZuluString(timezone ...string) string {
	return t.stdTimeFormatLayout(ISO8601ZuluLayout, timezone...)
}

func (t Time) ToISO8601ZuluMilliString(timezone ...string) string {
	return t.stdTimeFormatLayout(ISO8601ZuluMilliLayout, timezone...)
}

func (t Time) ToISO8601ZuluMicroString(timezone ...string) string {
	return t.stdTimeFormatLayout(ISO8601ZuluMicroLayout, timezone...)
}

func (t Time) ToISO8601ZuluNanoString(timezone ...string) string {
	return t.stdTimeFormatLayout(ISO8601ZuluNanoLayout, timezone...)
}

func (t Time) ToRFC1036String(timezone ...string) string {
	return t.stdTimeFormatLayout(RFC1036Layout, timezone...)
}

func (t Time) ToRFC1123String(timezone ...string) string {
	return t.stdTimeFormatLayout(RFC1123Layout, timezone...)
}

func (t Time) ToRFC1123ZString(timezone ...string) string {
	return t.stdTimeFormatLayout(RFC1123ZLayout, timezone...)
}

func (t Time) ToRFC2822String(timezone ...string) string {
	return t.stdTimeFormatLayout(RFC2822Layout, timezone...)
}

func (t Time) ToRFC3339String(timezone ...string) string {
	return t.stdTimeFormatLayout(RFC3339Layout, timezone...)
}

func (t Time) ToRFC3339MilliString(timezone ...string) string {
	return t.stdTimeFormatLayout(RFC3339MilliLayout, timezone...)
}

func (t Time) ToRFC3339MicroString(timezone ...string) string {
	return t.stdTimeFormatLayout(RFC3339MicroLayout, timezone...)
}

func (t Time) ToRFC3339NanoString(timezone ...string) string {
	return t.stdTimeFormatLayout(RFC3339NanoLayout, timezone...)
}

func (t Time) ToRFC7231String(timezone ...string) string {
	return t.stdTimeFormatLayout(RFC7231Layout, timezone...)
}

func (t Time) ToRFC822String(timezone ...string) string {
	return t.stdTimeFormatLayout(RFC822Layout, timezone...)
}

func (t Time) ToRFC822ZString(timezone ...string) string {
	return t.stdTimeFormatLayout(RFC822ZLayout, timezone...)
}

func (t Time) ToRFC850String(timezone ...string) string {
	return t.stdTimeFormatLayout(RFC850Layout, timezone...)
}

func (t Time) ToFormattedDateString(timezone ...string) string {
	return t.stdTimeFormatLayout(FormattedDateLayout, timezone...)
}

func (t Time) ToFormattedDayDateString(timezone ...string) string {
	return t.stdTimeFormatLayout(FormattedDayDateLayout, timezone...)
}

// ----------------------------------------------------------------------------------------------------

func (t Time) Layout(layout string, timezone ...string) string {
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.IsInvalid() {
		return ""
	}
	return t.StdTime().Format(layout)
}

func (t Time) Format(format string, timezone ...string) string {
	if len(timezone) > 0 {
		t.loc, t.Error = getLocationByTimezone(timezone[0])
	}
	if t.IsInvalid() {
		return ""
	}

	buffer := bytes.NewBuffer(nil)

	for i := 0; i < len(format); i++ {

		if layout, ok := formatMap[format[i]]; ok {

			switch format[i] {
			case 'D': // short week (Mon)
				buffer.WriteString(t.ToShortWeekString())
			case 'F': // month (January)
				buffer.WriteString(t.ToMonthString())
			case 'M': // short month (Jan)
				buffer.WriteString(t.ToShortMonthString())
			case 'S': // unix timestamp (1743465600)
				buffer.WriteString(strconv.FormatInt(t.Timestamp(), 10))
			case 'U': // timestamp with millisecond (1743465600000)
				buffer.WriteString(strconv.FormatInt(t.TimestampMilli(), 10))
			case 'V': // timestamp with microsecond (1743465600000000)
				buffer.WriteString(strconv.FormatInt(t.TimestampMicro(), 10))
			case 'X': // timestamp with nanosecond (1743465600000000000)
				buffer.WriteString(strconv.FormatInt(t.TimestampNano(), 10))
			default: // standard layout
				buffer.WriteString(t.StdTime().Format(layout))
			}

		} else {

			switch format[i] {
			case '\\': // no parse
				buffer.WriteByte(format[i+1])
				i++
				continue
			case 'W': // week number of the year in ISO-8601 format, ranging from 01-52
				week := fmt.Sprintf("%02d", t.WeekOfYear())
				buffer.WriteString(week)
			case 'N': // day of the week as a number in ISO-8601 format, ranging from 01-7
				week := fmt.Sprintf("%02d", t.DayOfWeek())
				buffer.WriteString(week)
			case 'K': // abbreviated suffix for the day of the month, such as st, nd, rd, th
				suffix := "th"
				switch t.Day() {
				case 1, 21, 31:
					suffix = "st"
				case 2, 22:
					suffix = "nd"
				case 3, 23:
					suffix = "rd"
				}
				buffer.WriteString(suffix)
			case 'L': // whether it is a leap year, if it is a leap year, it is 1, otherwise it is 0
				if t.IsLeapYear() {
					buffer.WriteString("1")
				} else {
					buffer.WriteString("0")
				}
			case 'G': // 24-hour format, no padding, ranging from 0-23
				buffer.WriteString(strconv.Itoa(t.Hour()))
			case 'w': // day of the week represented by the number, ranging from 0-6
				buffer.WriteString(strconv.Itoa(t.DayOfWeek() - 1))
			case 't': // number of days in the month, ranging from 28-31
				buffer.WriteString(strconv.Itoa(t.DaysInMonth()))
			case 'z': // current zone location, such as Asia/Tokyo
				buffer.WriteString(t.ZoneName())
			case 'o': // current zone offset, such as 28800
				buffer.WriteString(strconv.Itoa(t.ZoneOffset()))
			case 'q': // current quarter, ranging from 1-4
				buffer.WriteString(strconv.Itoa(t.Quarter()))
			case 'c': // current century, ranging from 0-99
				buffer.WriteString(strconv.Itoa(t.Century()))
			default:
				buffer.WriteByte(format[i])
			}

		}

	}

	return buffer.String()
}
