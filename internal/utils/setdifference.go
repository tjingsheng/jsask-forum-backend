package utils

func SetDifference(setA, setB []string) []string {
	setA = RemoveDuplicateStr(setA)
	setB = RemoveDuplicateStr(setB)

	m := make(map[string]bool)
	var diff []string

	for _, item := range setA {
		m[item] = true
	}

	for _, item := range setB {
		if _, in := m[item]; !in {
			diff = append(diff, item)
		}
	}
	return diff
}
