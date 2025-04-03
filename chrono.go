package chrono

import "time"

type Time struct {
	time         time.Time
	layout       string
	loc          *time.Location
	lang         *Language
	weekStartsAt time.Weekday
	Error        error
}

// New chrono Time instance
func New(time ...time.Time) Time {
	t := Time{lang: newLanguage()}
	t.loc, t.Error = getLocationByTimezone(defaultTimezone)
	if weekday, ok := weekdays[defaultWeekStartsAt]; ok {
		t.weekStartsAt = weekday
	}
	if len(time) > 0 {
		t.time = time[0]
		t.loc = t.time.Location()
	}
	t.layout = defaultLayout
	return t
}
