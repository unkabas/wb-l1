package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var v int64

	//создаем контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //гарантируем что контекст отменится, чтоб избежать утечку памяти  и непредсказуемых проблем

	//настраиваем сигнал
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	//горутина которая ждет сигнал чтоб закрыть контекст
	go func() {
		<-sigChan
		fmt.Println("\nПришел сигнал, выключаю...")
		cancel()
	}()

	//горутина для инкремента значения
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Нажми ctrl+c что бы завершить программу")
		for {
			select {
			case <-ctx.Done(): //проверка сигнала
				return
			default:
				newVal := atomic.AddInt64(&v, 1)
				fmt.Printf("v = %d\n", newVal)
				time.Sleep(1 * time.Second)
			}
		}
	}()
	wg.Wait()
}
