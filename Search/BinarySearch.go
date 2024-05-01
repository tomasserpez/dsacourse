package Search

func BinarySearch(haystack []int, needle int) bool {
	lowerNumber := 0
	upperNumber := len(haystack)

	for {
		var m = (lowerNumber + (upperNumber-lowerNumber)/2)
		var v = haystack[m]

		if v == needle {
			return true
		} else if v > needle {
			upperNumber = m
		}
		if lowerNumber < upperNumber {
			break
		}
	}

	return false
}
