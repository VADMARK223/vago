package main

import (
	"fmt"
	"math/rand/v2"
)

// RunSlices1 написать программу, которая возвращает слайс с заданным количеством случайных уникальных чисел
func RunSlices1() {
	fmt.Println(uniqRandom(10, 11))
}

func uniqRandom(n, max int) []int {
	if n > max {
		panic("n > max будет бесконечный цикл")
	}
	// Результирующий срез (резервируем емкость, чтобы избежать аллокаций базового массива)
	result := make([]int, 0, n)
	// Срез для кеширования чисел
	memory := make(map[int]struct{}, n)

	for len(result) < n {
		val := rand.IntN(max)
		if _, ok := memory[val]; ok {
			continue
		} else {
			memory[val] = struct{}{}
			result = append(result, val)
		}
	}

	return result
}
