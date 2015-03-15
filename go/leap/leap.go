package leap

func IsLeapYear(testYear int) bool {
	if testYear%4 == 0 {
		if testYear%100 == 0 && testYear%400 != 0 {
			return false
		}
		return true
	}
	return false
}
