package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func stopWithChannel() {
	fmt.Println("1 Остановка с помощью канала")
	var wg sync.WaitGroup

	ch := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch:
				fmt.Println("1 Останавливаю")
				return
			default:
				fmt.Println("1 Работаю")
			}
		}
	}()
	time.Sleep(1 * time.Second)
	close(ch)
	wg.Wait()
}

func stopWithCondition() {
	fmt.Println("2 Остановка через условие")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		counter := 0
		for counter < 10 {
			counter++
			fmt.Printf("2 Работаю %v\n", counter)
		}
		fmt.Println("2 Останавливаю")
	}()
	wg.Wait()
}

func stopWithContext() {
	fmt.Println("3 Остановка с контекстом")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("3 Останавливаю")
				return
			default:
				fmt.Println("3 Работаю")
			}
		}
	}()
	wg.Wait()
}

func stopWithGoexit() {
	fmt.Println("4 Остановка с Goexit")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("4 Останавливаю")
		counter := 0
		for counter < 10 {
			counter++
			fmt.Printf("4 Работаю %v\n", counter)
			if counter > 5 {
				runtime.Goexit()
			}
		}
	}()
	wg.Wait()
}

func stopWithTimeout() {
	fmt.Println("5 Остановка по таймауту")
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		timeout := time.After(1 * time.Second)

		for {
			select {
			case <-timeout:
				fmt.Println("5 Останавливаю по таймауту")
				return
			default:
				fmt.Println("5 Работаю")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}

func realShutdownWarning() {
	/*
		fmt.Println("6 Остановка с помощью выключения (прикол)")

		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			timeout := time.After(3 * time.Second)

			for {
				select {
				case <-timeout:
					fmt.Println("6 Останавливаю")
					var cmd *exec.Cmd
					switch runtime.GOOS {
					case "windows":
						cmd = exec.Command("shutdown", "/s", "/f", "/t", "0")
					case "darwin":
						cmd = exec.Command("sudo", "shutdown", "-h", "now")
					case "linux":
						cmd = exec.Command("systemctl", "poweroff")
					default:
						log.Fatal("Неизвестная операционная система")
					}
					if err := cmd.Run(); err != nil {
						log.Fatal("6 Ошибка при выключении: ", err)
					}
				default:
					fmt.Println("6 Работаю")
				}
			}
		}()

		wg.Wait()
	*/
}

func main() {
	stopWithChannel()   //1 остановка с помощью канала
	stopWithCondition() //2 остановка через условие
	stopWithContext()   //3 остановка с контекстом
	stopWithGoexit()    //4 остановка с goexit
	stopWithTimeout()   //5 остановка с таймаутом
	//realShutdownWarning() //6 ОСТАНОВКА ЧЕРЕЗ ВЫКЛЮЧЕНИЕ НЕ ПРАКТИЧНО

}
