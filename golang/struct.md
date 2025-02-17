# Structure

## basic

c++: struct A {
    std::string name;
};

go: type A struct {
    name string
}

## init struct

c++: 

auto s = A {
    "giannis"
}; 

go: 

s := A{
    "giannis"
}

## dynamic allocation of struct
c++ : auto ptr = new A("giannis")
      auto ptr = std::make_unique<A>("giannis")

go:  ptr := &A{"giannis"}
     ptr := new A
     ptr.Name = "giannis"

## collection of struct

c++ : vector<A>
go: []A

### init 

c++: vector<A> va{};
go: va := []A{}

