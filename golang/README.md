# Notes for golang

## Brief summary

### Adv: 

simple concurrency model: goroutines and channels
static linked

mimic ood: polymorphism, encapsulation and composition

### limitation: 
working with large chunk of memory and avoid fragmatation, bad for finiancial programming

## flow

### switch 

switch a, no () as compared with c++

### while style loop

using for and break to mimic while loop in c++

### for loop (slices or array)

for range 

c++: for(const auto& a : collections)

### type vs use (c++)
c++: use sometype = std::vector<std::string>
go: type Digit int

### enum class in go

c++: 

enum TTT:int {
    ZERO = 0,
    ONE,
    TWO,
}

go: 

const (
    ZERO = iota
    One
    Two
)

## Data Type
int
int(n)

uint
uint(n)
float(32/64)
no double
complex

### numeric_limits
c++: std::numeric_limits template
golang: math.MaxInt, math.MinInt, math.MaxFloat32, 


### global var
var Global int = 1

### constant var
const 

### ASCII char and unicode char
ASCII char: 1 byte
unicode char: multi bytes 
rune (int32): represent unicode char

c++: std::string
go: string, []byte("")

### string manip
1. length of strings

c++: .size()
go: len()

### buffer for io 

c++: vector<uchar8_t>
go: []byte 

## Composite Data Type

## Generic Programming

## Interface

## Reflection

## Packages and Modules

## Concurrency

## TestingUnix