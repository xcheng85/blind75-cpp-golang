package main

// generator number thread

// channel could be replaced by promise-async
// isp pattern in cpp

// read-only for next stages
const StreamCap int = 4

func srcStream() <-chan int {
	outCh := make(chan int, StreamCap)
	go func() {
		for i := 1; i <= StreamCap; i++ {
			outCh <- i
		}
		close(outCh)
	}()
	return outCh
}

// stage1: map (power 2 for example)
func power2Map(in <-chan int) <-chan int {
	outCh := make(chan int, StreamCap)

	go func() {
		// streaming read from previous stage
		for v := range in {
			outCh <- v * v
		}
		close(outCh)
	}()

	return outCh
}

// stage2: reduce
func reduce(in <-chan int) <-chan int {
	outCh := make(chan int, StreamCap)

	go func() {
		// streaming read from previous stage
		var sum int
		for v := range in {
			sum += v
		}
		outCh <- sum
		close(outCh)
	}()

	return outCh
}

func AsyncLaunchMapReduce() int {
	res := <-reduce(power2Map(srcStream()))
	return res
}
