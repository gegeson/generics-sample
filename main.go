package main

import (
	"fmt"
)

func main() {
	ptr := Ptr[bool](true)
	fmt.Println("ptr: ", *ptr)

	var sum int
	Apply([]int{10, 20}, func(i, v int) {
		sum += v
	})
	fmt.Println("apply: ", sum)

	ns := []int{1, 2, 3, 4}
	ms := Filter(ns, func(i, v int) bool {
		return v%2 == 0
	})
	fmt.Println("filter: ", ms)

	var ss []string = Map([]int{10, 20}, func(v int) string {
		return fmt.Sprintf("0x%x", v)
	})
	fmt.Println("map: ", ss)

	var t *Tuple[int, string] = New(10, "hello")
	fmt.Println("tuple: ", t)
}

func Ptr[T any](v bool) *bool {
	return &v
}

func Apply[T any](s []T, f func(i int, v T)) {
	for i, v := range s {
		f(i, v)
	}
}

func Filter[T any](target []T, f func(i int, v T) bool) []T {
	var result []T
	for i, v := range target {
		if f(i, v) {
			result = append(result, v)
		}
	}
	return result
}

func Map[T1, T2 any](s []T1, f func(i T1) T2) []T2 {
	result := make([]T2, 0, len(s))
	for _, v := range s {
		result = append(result, f(v))
	}
	return result
}

type Tuple[T1, T2 any] struct {
	X T1
	Y T2
}

func New[T1, T2 any](x T1, y T2) *Tuple[T1, T2] {
	return &Tuple[T1, T2]{X: x, Y: y}
}
