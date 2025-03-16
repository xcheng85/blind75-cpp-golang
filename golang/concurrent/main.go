package main

import (
	"fmt"

	"github.com/xcheng85/blind75-cpp-golang/golang/concurrent/sync"
)

func main() {
	sync.WordCount([]string{
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
	createSessionStreamChReadonly := sync.PublishAsync(&createSessionStreamCh, []string{
		"Dame Lillard",
		"Damian Lillard",
		"Giannis",
		"Gianni",
		"Mike Jordan",
		"Michael Jordan",
	})

	for task := range createSessionStreamChReadonly {
		fmt.Println(task)
	}
}
