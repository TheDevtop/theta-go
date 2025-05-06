package site

/*
	Theta virtual machine
	Arithmetic functions
*/

// // Cast list to typed list (best effort)
// func genCast[T any](list types.List) []T {
// 	var res = make([]T, 0, len(list))
// 	for _, exp := range list {
// 		if ce, ok := exp.(T); ok {
// 			res = append(res, ce)
// 		}
// 	}
// 	return res
// }

// func genAdd[T int32 | float32](set []T) T {
// 	var acc T = 0
// 	for _, v := range set {
// 		acc += v
// 	}
// 	return acc
// }

// func genMul[T int32 | float32](set []T) T {
// 	var acc T = 1
// 	for _, v := range set {
// 		acc *= v
// 	}
// 	return acc
// }

// func genSub[T int32 | float32](set []T) T {
// 	var acc T = set[0]
// 	for _, v := range set[1:] {
// 		acc -= v
// 	}
// 	return acc
// }

// func genDiv[T int32 | float32](set []T) T {
// 	var acc T = set[0]
// 	for _, v := range set[1:] {
// 		acc /= v
// 	}
// 	return acc
// }
