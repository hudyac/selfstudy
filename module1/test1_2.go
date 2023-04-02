package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	queue := make(chan int, 10)
	done := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for i := 0; ; i++{
			select {
			case <-done:
				return
			case <-ticker.C:
				select {
				case queue <- i:
					fmt.Println("Produced:", i)
				default:
					fmt.Println("Producer blocked")
				}
			}
		}
	}()

	go func() {
		defer wg.Done()
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				select {
				case item := <-queue:
					fmt.Println("Consumed:", item)
				default:
					fmt.Println("Consumer blocked")
				}
			}
		}
	}()

	time.Sleep(10 * time.Second)
	close(done)
	wg.Wait()
}
