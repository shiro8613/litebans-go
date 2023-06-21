package pages

import "strings"

func urlDecoder(str string, pattern string) map[string]string {
	if !strings.Contains(str, "/") || str == "" || !strings.Contains(pattern, "/") || pattern == "" {
		return map[string]string{}
	}

	splited_str := strings.Split(str, "/")
	splited_pattern := strings.Split(pattern, "/")

	if len(splited_str) < 0 || len(splited_pattern) < 0 || len(splited_str) != len(splited_pattern) {
		return map[string]string{}
	}

	return_map := make(map[string]string)

	for i := 0; len(splited_str) > i; i++ {
		for k := 0; len(splited_pattern) > k; k++ {
			if i == k {
				before := strings.Split(splited_pattern[k], "{")
				if len(before) > 1 {
					after := strings.Split(before[1], "}")
					if len(after) > 1 {
						name := after[0]
						return_map[name] = splited_str[i]
					}
				}
			}
		}
	}
	return return_map

}