#  Type conversion in c++ and go

### type conversion between int and float: 
go: float32(int)
c++: (float)int, static_cast<float>(int)

### type conversion from string to int
go: strconv.Atoi()
c++: stoi(i)
c++: atoi()

### type conversion from int to string
c++: std::to_string(), itoa
go: string() or strconv.Itoa() or strConv.FormatInt(int64(n), 10)

### type conversion from []byte to string
c++: .data()
go: string(byteslice)
