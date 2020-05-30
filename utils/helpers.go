package utils

func Contains(source []string, element string) bool {
	for _, x := range source {
		if x == element {
			return true
		}
	}

	return false
}
