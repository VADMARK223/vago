package main

import (
	"fmt"
	"unicode/utf8"
)

func Filter[T any](in []T, f func(T) bool) []T {
	var result []T
	for _, v := range in {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

type Box[T any] struct {
	Value T
}

func main() {
	// Отфильтровываем числа больше либо равно 10
	inInt := []int{1, 2, 30, 3, 60, 4, 5, 10}
	lessThanTen := func(value int) bool { return value < 10 }
	fmt.Println(Filter(inInt, lessThanTen)) // [1 2 3 4 5]

	// Отфильтровываем строки с длинной больше либо равно 3
	inStr := []string{"as", "ячс", "z"}
	lenLessThanThree := func(value string) bool {
		return utf8.RuneCountInString(value) < 3
	}
	fmt.Println(Filter(inStr, lenLessThanThree)) // [as z]

	b1 := Box[int]{Value: 1}
	b2 := Box[string]{Value: "Sac"}
}
