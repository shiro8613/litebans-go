package utils

func String2Bool(s string) bool {
	if !SliceContains([]string{"true", "false"}, s) {
		return false
	} else if s == "true" {
		return true
	} else {
		return false
	}
}
