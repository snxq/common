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

// ContainsStringF checks if a given slice of strings contains the provided string.
func ContainsStringF(slice []string, s string, f func(string, string) bool) bool {
	for _, item := range slice {
		if f(item, s) {
			return true
		}
	}
	return false
}

// ContainsString checks if a given slice of strings contains the provided string.
func ContainsString(slice []string, s string) bool {
	return ContainsStringF(slice, s, func(a, b string) bool { return a == b })
}

// ContainsUint32F check if a given slice of uint32 contains the provided uint32.
func ContainsUint32F(slice []uint32, n uint32, f func(uint32, uint32) bool) bool {
	for _, item := range slice {
		if f(item, n) {
			return true
		}
	}
	return false
}

// ContainsUint32 check if a given slice of uint32 contains the provided uint32.
func ContainsUint32(slice []uint32, n uint32) bool {
	return ContainsUint32F(slice, n, func(a, b uint32) bool { return a == b })
}
