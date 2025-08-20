package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, m := range arr {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("%d ", m*m)
		}()
	}
	wg.Wait()

}
