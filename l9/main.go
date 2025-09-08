package main

import "fmt"

func main() {
	arr := []int{1, 2, 14, 22, 31, 48, 43, 49, 42, 45, 33, 3, 6, 7, 3, 200, 552, 32}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for _, num := range arr {
			ch1 <- num
		}
		close(ch1)
	}()

	go func() {
		for num := range ch1 {
			ch2 <- num * 2
		}
		close(ch2)
	}()

	for result := range ch2 {
		fmt.Printf("%d ", result)
	}

}
