# Adapter Pattern

## OCP: which Solid principle it fulfills

OCP 

closeness: existing old interface
open: a new interface incompatible with old 

create a way to use the old functionality with a new signature

## Ways to implement 

1. Object-based Adaptor through inheritance 
2. Class-based Adaptor
3. Functor Adaptor: ADL (specific to c++)

## c++ impl

// object adapter

// function adapter
// ADL: correct overloaded function are called, even though they are
// overloaded in user-specific namespace

## go

## Adaptor any data structure to stack interface
