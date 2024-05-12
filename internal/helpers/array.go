package helpers

func InArray[T comparable](needle T, haystack []T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func RemoveFromArray[T comparable](needle T, haystack []T) []T {
	for i, v := range haystack {
		if v == needle {
			return append(haystack[:i], haystack[i+1:]...)
		}
	}
	return haystack
}
