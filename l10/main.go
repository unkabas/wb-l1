package main

import "fmt"

func main() {
	slice := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	//находим ключи целые десятки
	for _, num := range slice {
		key := int(num/10) * 10
		groups[key] = append(groups[key], num)
	}
	//собираем все ключи
	var keys []int
	for k := range groups {
		keys = append(keys, k)
	}
	//сортируем пузырем
	n := len(keys)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if keys[j] > keys[j+1] {
				keys[j], keys[j+1] = keys[j+1], keys[j]
			}
		}
	}
	for _, key := range keys {
		fmt.Printf("%d: %.1f\n", key, groups[key])
	}
}
