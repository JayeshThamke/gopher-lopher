package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	isLeapYear := false
	if year%4 == 0 {
		isLeapYear = true
		if year%100 == 0 {
			if year%400 == 0 {
				isLeapYear = true
			} else {
				isLeapYear = false
			}
		}
	}
	return isLeapYear
}
