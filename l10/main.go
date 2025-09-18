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

	var keys []int
	for k := range groups {
		keys = append(keys, k)
	}

	for _, key := range keys {
		fmt.Printf("%d: %v\n", key, groups[key])
	}

}
