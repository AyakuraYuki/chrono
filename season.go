package chrono

import (
	"strings"
)

var seasons = []struct{ month, index int }{
	{3, 0},  // spring
	{4, 0},  // spring
	{5, 0},  // spring
	{6, 1},  // summer
	{7, 1},  // summer
	{8, 1},  // summer
	{9, 2},  // autumn
	{10, 2}, // autumn
	{11, 2}, // autumn
	{12, 3}, // winter
	{1, 3},  // winter
	{2, 3},  // winter
}

func (t Time) Season() string {
	if t.IsInvalid() {
		return ""
	}

	if len(t.lang.resources) == 0 {
		t.lang.SetLocale(defaultLocale)
	}

	month, index := t.Month(), -1
	for i := 0; i < len(seasons); i++ {
		if month == seasons[i].month {
			index = seasons[i].index
			break
		}
	}

	t.lang.rw.Lock()
	defer t.lang.rw.Unlock()

	if resources, ok := t.lang.resources["seasons"]; ok {
		words := strings.Split(resources, "|")
		if len(words) == QuartersPerYear {
			return words[index]
		}
	}

	return ""
}

func (t Time) StartOfSeason() Time {
	if t.IsInvalid() {
		return t
	}
	year, month, _ := t.Date()
	if month == 1 || month == 2 {
		return t.create(year-1, 12, 1, 0, 0, 0, 0)
	}
	return t.create(year, month/3*3, 1, 0, 0, 0, 0)
}

func (t Time) EndOfSeason() Time {
	if t.IsInvalid() {
		return t
	}
	year, month, _ := t.Date()
	if month == 1 || month == 2 {
		return t.create(year, 3, 0, 23, 59, 59, 999999999)
	}
	if month == 12 {
		return t.create(year+1, 3, 0, 23, 59, 59, 999999999)
	}
	return t.create(year, month/3*3+3, 0, 23, 59, 59, 999999999)
}

func (t Time) IsSpring() bool {
	if t.IsInvalid() {
		return false
	}
	month := t.Month()
	return month == 3 || month == 4 || month == 5
}

func (t Time) IsSummer() bool {
	if t.IsInvalid() {
		return false
	}
	month := t.Month()
	return month == 6 || month == 7 || month == 8
}

func (t Time) IsAutumn() bool {
	if t.IsInvalid() {
		return false
	}
	month := t.Month()
	return month == 9 || month == 10 || month == 11
}

func (t Time) IsWinter() bool {
	if t.IsInvalid() {
		return false
	}
	month := t.Month()
	return month == 12 || month == 1 || month == 2
}
