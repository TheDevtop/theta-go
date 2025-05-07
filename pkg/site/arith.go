package site

/*
	Theta list processor
	Specialized arithmetic
*/

type number = interface{ int32 | float32 }

func arithAdd[T number](set ...T) T {
	var acc T = 0
	for _, v := range set {
		acc += v
	}
	return acc
}

func arithMul[T number](set ...T) T {
	var acc T = 1
	for _, v := range set {
		acc *= v
	}
	return acc
}

func arithSub[T number](set ...T) T {
	var acc T = set[0]
	for _, v := range set[1:] {
		acc -= v
	}
	return acc
}

func arithDiv[T number](set ...T) T {
	var acc T = set[0]
	for _, v := range set[1:] {
		acc /= v
	}
	return acc
}

func arithLesser[T number](set ...T) bool {
	for _, v := range set[1:] {
		if set[0] > v {
			return false
		}
	}
	return true
}

func arithGreater[T number](set ...T) bool {
	for _, v := range set[1:] {
		if set[0] < v {
			return false
		}
	}
	return true
}
