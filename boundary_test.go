package chrono

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTime_StartOfCentury(t *testing.T) {
	t.Run("zero time", func(t *testing.T) {
		assert.Equal(t, "0000-01-01 00:00:00 +0000 UTC", New().StartOfCentury().ToString())
	})
}

func TestTime_EndOfCentury(t *testing.T) {

}

func TestTime_StartOfDecade(t *testing.T) {

}

func TestTime_EndOfDecade(t *testing.T) {

}

func TestTime_StartOfYear(t *testing.T) {

}

func TestTime_EndOfYear(t *testing.T) {

}

func TestTime_StartOfQuarter(t *testing.T) {

}

func TestTime_EndOfQuarter(t *testing.T) {

}

func TestTime_StartOfMonth(t *testing.T) {

}

func TestTime_EndOfMonth(t *testing.T) {

}

func TestTime_StartOfWeek(t *testing.T) {

}

func TestTime_EndOfWeek(t *testing.T) {

}

func TestTime_StartOfDay(t *testing.T) {

}

func TestTime_EndOfDay(t *testing.T) {

}

func TestTime_StartOfHour(t *testing.T) {

}

func TestTime_EndOfHour(t *testing.T) {

}

func TestTime_StartOfMinute(t *testing.T) {

}

func TestTime_EndOfMinute(t *testing.T) {

}

func TestTime_StartOfSecond(t *testing.T) {

}

func TestTime_EndOfSecond(t *testing.T) {

}
