package main

import (
	"fmt"
	"strings"
)

// strings1 Напиши функцию, которая разворачивает строку. func reverse(s string) string
func strings1() {
	fmt.Println("result:", reverse("вадим"))
}

func reverse(s string) string {
	runes := []rune(s)
	fmt.Println("Кол-во байтов в строке: ", len(s))     // 10
	fmt.Println("Кол-во rune (символов): ", len(runes)) // 5

	// Понятный вариант:
	//var result = make([]rune, 0, len(runes))
	//for i := len(runes) - 1; i >= 0; i-- {
	//	result = append(result, runes[i])
	//}

	// Можно короче:
	var result = make([]rune, len(runes))
	for i := range runes {
		result[len(runes)-1-i] = runes[i]
	}

	return string(result)
}

// Удалить все повторяющиеся символы из строки, сохранив порядок.
func strings2() {
	fmt.Println("result:", unique("вваааддиииммм"))
}

func unique(s string) string {
	runes := []rune(s)
	seen := make(map[rune]struct{}, len(runes))
	var result strings.Builder
	result.Grow(len(runes)) // Необязательная оптимизация

	for _, r := range runes { // Можно идти и по s
		_, ok := seen[r]
		if !ok {
			seen[r] = struct{}{}
			result.WriteRune(r)
		}
	}

	return result.String()
}
