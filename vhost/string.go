package vhost

import "reflect"

const (
	NilString = "<nil>"
)

func NilEmpty(s string) string {
	if s == "" {
		return NilString
	}
	return s
}

func NilZero_UNUSED[T any](s T) string {
	var t T

	i := any(t)
	switch reflect.TypeOf(t).Kind() {
	case reflect.String:
		v, ok := i.(string)
		if ok {
			if v == "" {
				return NilString
			}
			return v
		}

	}
	return "error: type not supported"
}
