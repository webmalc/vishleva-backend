package services

// StringInSlice check if the string in the slice.
func StringInSlice(val string, slice []string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}

	return -1, false
}
