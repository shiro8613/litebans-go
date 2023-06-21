package utils

import "strings"

func ContainSwitch(s string, m map[string]func()) {
	for k,v := range m {
		if strings.Contains(s, k) {
			v()
			return
		}
	}
}