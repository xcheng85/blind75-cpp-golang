
### type conversion between int and float
float32(int)

### type conversion from string to int
go: strconv.Atoi()

c++: atoi()

### type conversion from int to string
c++: std::to_string(), itoa
go: string() or strconv.Itoa() or strConv.FormatInt(int64(n), 10)

### type conversion from []byte to string
c++:
go: string(byteslice)
