package chrono

import "strings"

var constellations = []struct{ startMonth, startDay, endMonth, endDay int }{
	{3, 21, 4, 19},   // Aries
	{4, 20, 5, 20},   // Taurus
	{5, 21, 6, 21},   // Gemini
	{6, 22, 7, 22},   // Cancer
	{7, 23, 8, 22},   // Leo
	{8, 23, 9, 22},   // Virgo
	{9, 23, 10, 23},  // Libra
	{10, 24, 11, 22}, // Scorpio
	{11, 23, 12, 21}, // Sagittarius
	{12, 22, 1, 19},  // Capricorn
	{1, 20, 2, 18},   // Aquarius
	{2, 19, 3, 20},   // Pisces
}

var constellationsMapping = map[string]struct{ startMonth, startDay, endMonth, endDay int }{
	Aries:       {3, 21, 4, 19},
	Taurus:      {4, 20, 5, 20},
	Gemini:      {5, 21, 6, 21},
	Cancer:      {6, 22, 7, 22},
	Leo:         {7, 23, 8, 22},
	Virgo:       {8, 23, 9, 22},
	Libra:       {9, 23, 10, 23},
	Scorpio:     {10, 24, 11, 22},
	Sagittarius: {11, 23, 12, 21},
	Capricorn:   {12, 22, 1, 19},
	Aquarius:    {1, 20, 2, 18},
	Pisces:      {2, 19, 3, 20},
}

func (t Time) Constellation() string {
	if t.IsInvalid() {
		return ""
	}

	if len(t.lang.resources) == 0 {
		t.lang.SetLocale(defaultLocale)
	}

	_, month, day := t.Date()
	index := -1
	for i := 0; i < len(constellations); i++ {
		if month == constellations[i].startMonth && day >= constellations[i].startDay {
			index = i
		}
		if month == constellations[i].endMonth && day <= constellations[i].endDay {
			index = i
		}
	}

	t.lang.rw.Lock()
	defer t.lang.rw.Unlock()

	if resources, ok := t.lang.resources["constellations"]; ok {
		words := strings.Split(resources, "|")
		if len(words) == MonthsPerYear {
			return words[index]
		}
	}

	return ""
}

func (t Time) isConstellation(constellation string) bool {
	if t.IsInvalid() {
		return false
	}
	dateRange, ok := constellationsMapping[constellation]
	if !ok {
		return false
	}
	_, month, day := t.Date()
	if month == dateRange.startMonth && day >= dateRange.startDay {
		return true
	}
	if month == dateRange.endMonth && day <= dateRange.endDay {
		return true
	}
	return false
}

func (t Time) IsAries() bool       { return t.isConstellation(Aries) }
func (t Time) IsTaurus() bool      { return t.isConstellation(Taurus) }
func (t Time) IsGemini() bool      { return t.isConstellation(Gemini) }
func (t Time) IsCancer() bool      { return t.isConstellation(Cancer) }
func (t Time) IsLeo() bool         { return t.isConstellation(Leo) }
func (t Time) IsVirgo() bool       { return t.isConstellation(Virgo) }
func (t Time) IsLibra() bool       { return t.isConstellation(Libra) }
func (t Time) IsScorpio() bool     { return t.isConstellation(Scorpio) }
func (t Time) IsSagittarius() bool { return t.isConstellation(Sagittarius) }
func (t Time) IsCapricorn() bool   { return t.isConstellation(Capricorn) }
func (t Time) IsAquarius() bool    { return t.isConstellation(Aquarius) }
func (t Time) IsPisces() bool      { return t.isConstellation(Pisces) }
