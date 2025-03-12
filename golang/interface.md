# Interface virtual base class c++ vs go for OOP



## class member function --> type method (value and pointer receiver)



## sorting custom struct
c++: rely on operator overload

go: relay on implement sort.Interface: three methods (Len, Less, Swap)

## type assertion: rtii
c++: void*
go: interface{}

c++: static_cast<T*>, dynamic_cast<T*>,
c++ smart pointer cast: static_pointer_cast, dynamic_pointer_cast

```c++ rtii mechnisum
auto derived = dynamic_cast<Derived*>(b)
if(derived) {

}

typeid(b2).name()

```

```go
x.(type)

```

## type assertion: static_cast

```c++

```

```go
Derived, ok:= base.(DerivedClass)
```

## generic map go > c++

c++: all the value in map kv pair must be the same type 
go: map[string]interface{}, heterogenous type



