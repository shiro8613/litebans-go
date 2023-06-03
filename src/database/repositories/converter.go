package repositories

import "reflect"

func Converter[T any, V any](a T) V {

	b := new(V)

	vb := reflect.ValueOf(b)
	vs := reflect.ValueOf(&a)

	if vs.Elem().NumField() != vb.Elem().NumField() {
		return *new(V)
	}

	for i := 0; i < vs.Elem().NumField(); i++ {
		if vs.Elem().Field(i).Type().String() == "int" {
			vb.Elem().Field(i).SetInt(vs.Elem().Field(i).Int())
		} else if vs.Elem().Field(i).Type().String() == "[]uint8" {
			if len(vs.Elem().Field(i).Bytes()) == 1 {
				if vs.Elem().Field(i).Bytes()[0] == 1 {
					vb.Elem().Field(i).SetString("true")
				} else {
					vb.Elem().Field(i).SetString("false")
				}
			} else {
				if reflect.DeepEqual(vs.Elem().Field(i).Bytes(), nil) {
					vb.Elem().Field(i).SetString("")
				} else {
					vb.Elem().Field(i).SetString(string(vs.Elem().Field(i).Bytes()))
				}
			}
		}
	}

	return *b

}
