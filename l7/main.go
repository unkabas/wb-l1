package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)

	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			m[i] = i * 10
		}(i)
	}
	wg.Wait()
	fmt.Printf("%v\n", m)
}
