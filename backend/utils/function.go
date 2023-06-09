package utils

import "reflect"

// Find: return copy of item by pointer,
// return nil, -1 if found nothing
func Find[T any](arr []T, conditionCb func(index int, item T) bool) (*T, int) {
	var found T
	for index, item := range arr {
		if conditionCb(index, item) {
			found = item
			return &found, index
		}
	}
	return nil, -1
}

// IsEmpty: check is zero value of comparable type
// return true if is empty
func IsEmpty[T any](value T) bool {
	return reflect.ValueOf(value).IsZero()
}

// // Loop array and use callback perform mapper
// // T: original type, K: new type
// // return array of new type
// func ForEachTransform[T any, K any](arr []T, cb func(i int, j T) *K) []*K {
// 	newArr := []*K{}
// 	for i, j := range arr {
// 		newArr = append(newArr, cb(i, j))
// 	}
// 	return newArr
// }

// func Filter[T any](arr []T, conditionCb func(index int, item T) bool) []T {
// 	newArr := []T{}
// 	for index, item := range arr {
// 		if conditionCb(index, item) {
// 			newArr = append(newArr, item)
// 		}
// 	}
// 	return newArr
// }
