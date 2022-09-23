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

func IsPointer(i any) bool {
	if i == nil {
		return false
	}
	if reflect.TypeOf(i).Kind() != reflect.Pointer {
		return false
	}
	return true
}

// IsNil - determine if the interface{} holds is nil
func IsNil(i any) bool {
	if i == nil {
		return true
	}
	if !IsPointer(i) {
		return false
	}
	return reflect.ValueOf(i).IsNil()
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
