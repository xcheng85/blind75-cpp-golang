package sync

import (
	"fmt"
	"strings"
)

func WordCount(sentences []string) {
	// unordered_map
	countMap := make(map[string]int)
	// unbuffered channel for sync between goroutine, like promise-future
	done := make(chan bool)
	// goroutine + anonymous function (lamda in c++)
	// std::async([](){})
	go func() {
		// capture words
		for _, s := range sentences {
			words := strings.Split(s, " ")
			for _, w := range words {
				w = strings.ToLower(w)
				// one thread write to shared memory
				countMap[w]++
			}
		}
		// set promise in c++
		done <- true
	}()
	// like future.get() in c++
	if <-done {
		for k, v := range countMap {
			fmt.Printf("%s\t(%d)\n", k, v)
		}
	}
}
