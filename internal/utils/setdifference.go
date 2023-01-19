package utils

func SetDifference[T string | int](setA, setB []T) []T {
	m := make(map[T]bool)
	var diff []T

	for _, item := range setB {
		m[item] = true
	}

	for _, item := range setA {
		if _, in := m[item]; !in {
			diff = append(diff, item)
		}
	}
	return diff
}
