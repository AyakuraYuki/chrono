package chrono

var (
	defaultLayout       = DateTimeLayout // default layout: 2006-01-02 15:04:05
	defaultLocale       = "en"           // default language locale: en
	defaultTimezone     = UTC            // default timezone: UTC
	defaultWeekStartsAt = Sunday         // default week starts at Sunday
)

type Default struct {
	Layout       string
	Locale       string
	Timezone     string
	WeekStartsAt string
}

// SetDefault sets default
func SetDefault(d Default) {
	if d.Layout != "" {
		defaultLayout = d.Layout
	}
	if d.Locale != "" {
		defaultLocale = d.Locale
	}
	if d.Timezone != "" {
		defaultTimezone = d.Timezone
	}
	if d.WeekStartsAt != "" {
		defaultWeekStartsAt = d.WeekStartsAt
	}
}

// ResetDefault resets default settings
func ResetDefault() {
	defaultLayout = DateTimeLayout
	defaultLocale = "en"
	defaultTimezone = UTC
	defaultWeekStartsAt = Sunday
}
