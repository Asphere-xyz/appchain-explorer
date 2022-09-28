package common

import "reflect"

func CreateEntityDiff[T interface{}](a, b *T) (r, n T) {
	a2, b2 := reflect.ValueOf(a).Elem(), reflect.ValueOf(b).Elem()
	n = *a
	r2, n2 := reflect.ValueOf(&r), reflect.ValueOf(&n)
	for i := 0; i < a2.NumField(); i++ {
		val1, val2 := a2.Field(i), b2.Field(i)
		if reflect.DeepEqual(val1.Interface(), val2.Interface()) {
			continue
		}
		r2.Elem().Field(i).Set(val2)
		n2.Elem().Field(i).Set(val2)
	}
	return r, n
}
