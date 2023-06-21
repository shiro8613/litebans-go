package utils

import "strings"

func Replacer(b1 []byte, b2 []byte, d string) []byte {
	t := strings.Replace(string(b1), string(b2), d, 1)
	return []byte(t)
}