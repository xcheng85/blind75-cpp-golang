
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


## Deep/Shallow copy of array-like

c++: deep copy = 

golang: slices.Clone

## search element
c++: std::find

golang: slices.Contains

slices.ContainsFunc(playersCopy, func(p Player) bool {
		return p.Number == 34
	})

## memory shrink 
c++: shrink_to_fit

golang: slices.Clip()

## min and max
return iterator
c++: std::min     const auto [min, max] = std::minmax_element(begin(v), end(v));

golang: slices.Min()/Max()

## replace algo in a range [)


c++:     std::replace(s.begin(), s.end(), 8, 88);

golang: slices.Replace(s, 1, 2, 8, 88)