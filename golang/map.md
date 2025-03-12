#

## c++ 

std::unordered_map<int, int> hashmap

### c++ literal
std::unordered_map<int, string> hashmap{
    {},
    {}
}


## golang

map[int]string

### golang literal
m := map[int]string{
    1: "" 
    2: ""
}

### make
m := make(map[int]string)

## key numbers

c++: .size()
go: len()

## deletion
c++: erase(k)
go: delete(m, k)

## check key exists
c++: .count() > 0, find() != enditr

go: v, ok := m[k]

## iteration 
c++: for(const auto& [k, v] : m)

go: for k,v := range m {
    
}

## erase kv pair in map 

c++: 

```c++

  const auto count = std::erase_if(data, [](const auto& item)
    {
        auto const& [key, value] = item;
        return (key & 1) == 1;
    });
 
``

golang: 

```go
    	maps.DeleteFunc(playerJerseyMap, func(name string, jersey int) bool {
		return name == "Giannis"
	})
```

## copy map

separate memory must use

maps.Clone()

same memory:

:=

## insert kv to map

c++: 

   insert + make_pair
   []

go:
   []


## copy_if algorithm c++ vs go