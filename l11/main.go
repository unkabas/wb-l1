package main

import "fmt"

func intersection(a, b []int) []int {
	set := make(map[int]bool)
	var result []int

	// Добавляем все элементы первого слайса в map
	for _, item := range a {
		set[item] = true
	}

	// Проверяем элементы второго слайса на наличие в map
	for _, item := range b {
		if set[item] {
			result = append(result, item)
			set[item] = false // Помечаем как обработанное, чтобы избежать дубликатов
		}
	}

	return result
}
func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	fmt.Println(intersection(a, b))
}
