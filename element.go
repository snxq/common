package common

import "reflect"

// IsZero 检查是否为零值
func IsZero(x interface{}) bool {
	if x == nil {
		return true
	}
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

// AllF 所有元素判断为真时 返回真, 否则为假
func AllF(f func(interface{}) bool, x ...interface{}) bool {
	for _, e := range x {
		if !f(e) {
			return false
		}
	}
	return true
}

// AllZero 所有元素为零值时, 返回真, 否则为假
func AllZero(x ...interface{}) bool {
	return AllF(IsZero, x...)
}

// AnyF 存在一个元素为真 返回真, 否则为假
func AnyF(f func(interface{}) bool, x ...interface{}) bool {
	for _, e := range x {
		if f(e) {
			return true
		}
	}
	return false
}

// AnyZero 存在一个元素为零值, 返回真, 否则为假
func AnyZero(x ...interface{}) bool {
	return AnyF(IsZero, x...)
}
