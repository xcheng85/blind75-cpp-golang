package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/xcheng85/blind75-cpp-golang/golang/concurrent/streaming"
	sync2 "github.com/xcheng85/blind75-cpp-golang/golang/concurrent/sync"
)

func main() {
	sync2.WordCount([]string{
		"Dame Lillard",
		"Damian Lillard",
		"Giannis",
		"Gianni",
		"Mike Jordan",
		"Michael Jordan",
	})

	// create consumers
	createSessionStreamCh := make(chan string)
	// no-blocking due to a internal goroutine
	createSessionStreamChReadonly := streaming.PublishAsync(&createSessionStreamCh, []string{
		"Dame Lillard",
		"Damian Lillard",
		"Giannis",
		"Gianni",
		"Mike Jordan",
		"Michael Jordan",
	})

	var wg sync.WaitGroup
	wg.Add(6)
	// create two consumers within one consumer group
	consumerTaskFunc := func(consumerId string){
		defer wg.Done()
		for task := range createSessionStreamChReadonly {
			fmt.Println(task)
			break
		}
		fmt.Println(consumerId)
	}
	// execute order is random every time, up to goroutine scheduler
	go consumerTaskFunc("1")
	go consumerTaskFunc("2")
	go consumerTaskFunc("3")
	go consumerTaskFunc("4")
	go consumerTaskFunc("5")
	go consumerTaskFunc("6")

	wg.Wait()

	// promise.any in go
	{
		stopCh := make(chan struct{})
		workerThread := func(sleep time.Duration, ch *chan struct{}){
			time.Sleep(sleep)
			close(*ch)
		}
		go workerThread(time.Second, &stopCh)
		go workerThread(3 * time.Second, &stopCh)
		
		select {
		case <-stopCh:
			return
		}
	}
}
