package utils

import "reflect"

func SliceContains[T any](slice []T, str T) bool {
	strl := reflect.ValueOf(&str)
	sl := reflect.ValueOf(&slice)
	for i := 0; i <= sl.Elem().NumField(); i++ {
	 	inter := sl.Elem().Field(i).Interface();
		if inter == strl.Elem().Field(0).Interface() {
			return true
		}
	}

	return false
}