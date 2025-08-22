package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	n := 10 //длительность
	fmt.Printf("Service will work %d seconds\n", n)

	ch := make(chan int) //главный канал

	var wg sync.WaitGroup
	done := make(chan struct{}) //канал для завершения

	wg.Add(2)
	go Writer(ch, done, &wg)
	go Reader(ch, done, &wg)

	time.Sleep(time.Duration(n) * time.Second) //спим пока горутины работают
	close(done)                                //сигнализируем о конце
	wg.Wait()
}

func Writer(ch chan int, done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)
	counter := 1
	for {
		select {
		case ch <- counter:
			fmt.Printf("Write: %d\n", counter)
			counter++
			time.Sleep(500 * time.Millisecond)
		case <-done:
			fmt.Println("Writer closed")
			return
		}
	}
}

func Reader(ch chan int, done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case value := <-ch:
			fmt.Printf("Read : %v\n", value)

		case <-done:
			fmt.Println("Reader closed")
			return
		}
	}

}
