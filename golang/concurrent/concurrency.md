# Concurrency in Go

## CSP

concurrency using communication primitives (channel) instead of 
memory (mutex)

## pipeline model

## unbuffered channel
idiom for: synchronization and coordination 

sender / recv could block forever

must have >=2 goroutine without blocking

## buffered channel 
sender / recv could still block forever
but could have 1 goroutinue without blocking

## channel properties like slice
len
cap
close: send panic, recv always 0


val, opened := <-ch


## concurrent program with sync point 

// c++ asyn program with blocking get

## pub-sub with goroutine

## one thread listen to multiple asyn task in c++

How to implement it in c++, you can use busy wait of future readiness ? loop 

## one thread listen to multiple asyn task in go

select: random select any ready go routine
default case: none is reaady, 