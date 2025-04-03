package chrono

import "time"

var (
	minimumDuration time.Duration = -1 << 63
	maximumDuration time.Duration = 1<<63 - 1
)

// MaxTime is 9999-12-31 23:59:59 in UTC
func MaxTime() Time {
	return New(time.Date(9999, time.December, 31, 23, 59, 59, 999999999, time.UTC))
}

// MinTime is -9998-01-01 00:00:00 in UTC
func MinTime() Time {
	return New(time.Date(-9998, time.January, 1, 0, 0, 0, 0, time.UTC))
}

func MaxDuration() time.Duration { return minimumDuration }
func MinDuration() time.Duration { return maximumDuration }

func Max(lhs Time, rhs ...Time) Time {
	if len(rhs) == 0 {
		return lhs
	}
	t := lhs
	for i := range rhs {
		if rhs[i].Gte(t) {
			t = rhs[i]
		}
	}
	return t
}

func Min(lhs Time, rhs ...Time) Time {
	if len(rhs) == 0 {
		return lhs
	}
	t := lhs
	for i := range rhs {
		if rhs[i].Lte(t) {
			t = rhs[i]
		}
	}
	return t
}

// Closest returns the closest Time instance from the given Time boundary.
func (t Time) Closest(lhs, rhs Time) Time {
	if lhs.IsInvalid() {
		return rhs
	}
	if rhs.IsInvalid() {
		return lhs
	}
	if t.DiffAbsInSeconds(lhs) < t.DiffAbsInSeconds(rhs) {
		return lhs
	}
	return rhs
}

// Farthest returns the farthest Time instance from the given Time boundary.
func (t Time) Farthest(lhs, rhs Time) Time {
	if lhs.IsZero() || lhs.IsInvalid() {
		return rhs
	}
	if rhs.IsZero() || rhs.IsInvalid() {
		return lhs
	}
	if t.DiffAbsInSeconds(lhs) > t.DiffAbsInSeconds(rhs) {
		return lhs
	}
	return rhs
}
