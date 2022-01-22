package util

func ContainsString(values []string, search string) bool {
	for _, value := range values {
		if search == value {
			return true
		}
	}
	return false
}
