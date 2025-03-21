package main

import (
	"fmt"
	"math/rand"
	"time"
)

const timeoutMs int = 5000

type TaskInput string

type TaskResp struct {
	Err  error
	Resp string
}

func randSeconds(from, to int) time.Duration {
	source := rand.NewSource(time.Now().UnixNano())
	randomNumGenerator := rand.New(source)
	seconds := randomNumGenerator.Intn(to - from) + from
	// no implicit ctor in go
	return time.Duration(seconds * int(time.Second))
}

// internal
// out like promise
func taskRunner(task TaskInput, out chan <- TaskResp) {
	// sleep to simulate
	time.Sleep(randSeconds(1,10))

	res := TaskResp{
		nil, 
		fmt.Sprintf("taskRunner completed for task: %s !", task),
	}
	out <- res
}

func AsynRunTasksAndWait(tasks ...TaskInput) []TaskResp{
	numConcurrentTasks := len(tasks)
	// ownership of futures. like vector<std::future>
	// non block sender for all the threads
	bufferedCh := make(chan TaskResp, numConcurrentTasks)
	// like vector, reserve
	responses := make([]TaskResp, numConcurrentTasks)
	// fire async and cache future into stl container alike
	for _, task :=range tasks {
		// all the threads will write to the bufferedCh
		go taskRunner(task, bufferedCh)
	}
	// non block for both sides of the channel
	for i := 0; i < numConcurrentTasks; i++ {
		// read bufferedCh, out of order
		responses[i] = <-bufferedCh
	}
	return responses
}
