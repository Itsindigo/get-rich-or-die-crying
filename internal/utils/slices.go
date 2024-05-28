package utils

func StringInSlice(s1 string, list []string) bool {
	for _, s2 := range list {
		if s2 == s1 {
			return true
		}
	}
	return false
}
