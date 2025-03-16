package sync

import "time"


// add task to a stream channel
// readonly channel  <-chan
func PublishAsync(ch *chan string, tasks []string) <-chan string {
	// async, two goroutine
	go func() {
		defer close(*ch)
		for _, task := range tasks {
			time.Sleep(1 * time.Second)
			*ch <- task
		}
	}()
	return *ch
}
