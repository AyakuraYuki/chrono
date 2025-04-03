package chrono

// MaxTime is 9999-12-31 23:59:59 in UTC
func MaxTime() Time {
	return New().create(9999, 12, 31, 23, 59, 59, 999999999, UTC)
}

// MinTime is -9998-01-01 00:00:00 in UTC
func MinTime() Time {
	return New().create(-9998, 1, 1, 0, 0, 0, 0, UTC)
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
