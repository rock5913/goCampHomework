package week1

import (
	"fmt"
)

func SliceDelVal[T any](src []T, index int, shrink bool) ([]T, T, error) {
	if index < 0 || index >= len(src) {
		var zero T
		return nil, zero, errIndexOutOfRange(len(src), index)
	}
	res := src[index]
	for i := index; i < len(src)-1; i++ {
		src[i] = src[i+1]
	}
	src = src[:len(src)-1]
	if shrink {
		src = sliceShrink(src)
	}
	return src, res, nil
}

func errIndexOutOfRange(length, index int) error {
	return fmt.Errorf("index out of range, length: %d, index: %d", length, index)
}

func sliceShrink[T any](src []T) []T {
	if cap(src)-len(src) <= 64 {
		return src
	}
	var capacity int
	if cap(src) > 2048 && cap(src)/len(src) >= 2 {
		capacity = int(0.625 * float64(cap(src)))
	} else if cap(src) <= 2048 && cap(src)/len(src) >= 4 {
		capacity = cap(src) / 2
	} else {
		return src
	}
	res := make([]T, len(src), capacity)
	copy(res, src)
	return res
}
