# generic programming paradigm

## c++ print array-like container

```c++

template<typename T>
void print(const std::vector<T>& container) {

}

```

## go

```go
func print[T any](container []T) {
    for _, v := range container {
        fmt.Println(v)
    }
}
```


## constraints

c++: concept , requires requires

```c++
template <typename TG>
concept IsThreadGroup = requires(TG x) {
    x.thread_rank();
    x.size();
};
```

go: comparable

create cpncept vs create constraints in go

``` go

type NumericConstraint interface {
    int | int8 
}


```

## generic tree/linkedlist node

```c++

template<typename T> 
struct Node {
    T value;
    Node* next;
};

```

```go

type Node[T any] struct {
    value T
    next *Node[T]
}


```