## GO CLI
- go build - Complies a bunch of go source code fiels
- go run - Compiles and executes one or two files
- go fmt - Formats all the code in each file in the current directory
- go install - Compiles and installs a package
- go get - Downloads the raw source code of someone else's package
- go test - Runs any tests associated with the current project

## Basic data type
- bool
- string
- int
- float64

## Basic Go Types
- string
- integer
- float
- array
- map

## Type of list
- Array - fixed length list of things
- Slice 
  - An array that can grow and shrink
  - Every element in a slice must be of same type

## Test
- To make a test, create a new file ending in _test.go

## Default Zero Value
- string = ""
- int = 0
- float = 0
- bool = false

## Pointer
- &variable - Give me the memory address of the value this variable is pointing at
- *pointer - Give me the value this memory address is pointing at

## Types
### Value Types - use pointers to change these things in a function
- int
- float
- string
- bool
- structs

### Reference Types
- slices
- maps
- channels
- pointers
- functions

## FAQ
1. What does 'package main' mean?
Package == Project == Workspace
2. Type of Packages
- Executable - Generates a file that we can run
- Reusable - Code used as 'helpers'. Good place to put reusable logic
3. All standard packages in Go
- https://golang.org/pkg/