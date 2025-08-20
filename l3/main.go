package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

const sec = 10 //кол-во секунд для записи рандомных значений в канал

func main() {
	//для запуска требуется параметр отвещающий за кол-во воркеров
	if len(os.Args) <= 1 {
		log.Fatal("Требуется параметр - количество воркеров")
	}

	//os.Args всегда строка поэтому переводим в int
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Параметр должен быть числом")
	}
	var ch = make(chan int)
	var wg = sync.WaitGroup{}

	//запускаем n кол-во горутин читающих канал
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for num := range ch {
				fmt.Printf("Worker %d: %d\n", workerID, num)
			}
			fmt.Printf("Worker %d finished\n", workerID)
		}(i)
	}

	//главная горутина заполняющая какнал рандомными числами
	go func() {
		defer close(ch)
		timeout := time.After(sec * time.Second)
		for {
			select {
			case <-timeout:
				return
			default:
				ch <- rand.Intn(500)
			}
		}
	}()
	wg.Wait()
}
