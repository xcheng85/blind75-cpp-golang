# SOLID in go 

## SRP 

design small interfaces

### Type and Operation 


## OCP

Close: stable layer, existed interface 

Open: new struct which fulfill the interface

### Smelly code that broke OCP

hardcoded branch (if-else, switch)

### go to implement OCP
composition in struct or function level
embeded class