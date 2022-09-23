package util

import "reflect"

// Copy - return a copy of an interface{}, this will dereference a pointer.
func Copy(i any) any {
	if i == nil {
		return nil
	}
	t := reflect.Indirect(reflect.ValueOf(i))
	return t.Interface()
}

func IsNillable(a any) bool {
	return IsPointer(a) || IsPointerType(a)
}

func IsPointer(a any) bool {
	if a == nil {
		return false
	}
	if reflect.TypeOf(a).Kind() != reflect.Pointer {
		return false
	}
	return true
}

func IsPointerType(a any) bool {
	if a == nil {
		return false
	}
	switch reflect.ValueOf(a).Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map:
		return true
	}
	return false
}

// IsNil - determine if the interface{} holds is nil
func IsNil(a any) bool {
	if a == nil {
		return true
	}
	if !IsNillable(a) {
		return false
	}
	return reflect.ValueOf(a).IsNil()
}

// IsSerializable - determine if any can be serialized
func IsSerializable(i any) bool {
	if IsNil(i) {
		return false
	}
	if _, ok := i.([]byte); ok {
		return false
	}
	if _, ok := i.(string); ok {
		return false
	}
	return true
}
