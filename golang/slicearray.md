
### Array vs C++ std::array<int>
c++: std::array<T, 4>()
go: [4]string{} laterall, function array args: copy

### Slice vs C++ vector
go: []string{}, pass by reference, size and capacity

make([]string, 2)

c++: vector<string>(2)

append vs push_back
// unpack
go append variadic: append(someSlice, []int{}...)

len vs size()

2d vector c++: vector<vector<string>>(2, vector<string>{"", ""})


## delete element i from collection

c++: iterator,

go: 


## copy collection to another collection

c++: std::copy + back_inserter
go: copy

## sort collection

c++: std::sort(begin, end, [](){

})

golang: 

c++: std::sort(, , std::greater<T>())