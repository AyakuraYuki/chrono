package chrono

import (
	"errors"
	"fmt"
)

func emptyLocaleError() error {
	return errors.New("locale is required")
}

func invalidLocaleError(locale string) error {
	return fmt.Errorf("invalid locale file %q", locale)
}

func invalidResourcesError() error {
	return errors.New("invalid resources")
}

func emptyTimezoneError() error {
	return errors.New("timezone is required")
}

func invalidTimezoneError(timezone string) error {
	return fmt.Errorf(`invalid timezone %q, please see the file "$GOROOT/lib/time/zoneinfo.zip" for all valid timezones`, timezone)
}

func emptyDurationError() error {
	return errors.New("duration is required")
}

func invalidDurationError(duration string) error {
	return fmt.Errorf("invalid duration %q", duration)
}

func invalidLocationError() error {
	return errors.New("invalid location, please make sure the location is valid")
}

func nilLanguageError() error {
	return errors.New("language is required")
}

func emptyLayoutError() error {
	return errors.New("layout is required")
}

func emptyFormatError() error {
	return errors.New("format is required")
}

func emptyWeekStartsAtError() error {
	return fmt.Errorf("week start day is required")
}

func invalidWeekStartsAtError(day string) error {
	return fmt.Errorf("invalid week starts at day %s, acceptable options are [Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday]", day)
}
