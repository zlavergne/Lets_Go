Data Types and Syntax
---

#### bool
  - true or false

#### Integer
  - int: Size is implementation-defined, it is either 32 or 64 bits.
  - int8
  - int16
  - int32
  - int64

#### Unsigned Integer
  - uint: Same size as int
  - uint8
  - uint16
  - uint32
  - uint64
  - uintptr

#### Float
  - float32
  - float64

#### complex
  - complex64: A float32 real and float32 imaginary
  - complex128: A float64 real and float64 imaginary

#### byte
  - Alias for uint8

#### rune
  - Alias for int32
  - Also represents a Unicode point

#### string
  - Immutable
  - Arbitrary slice of bytes (Not necessarily UTF-8 or ASCII)

#### Arrays
  - Fixed number of contiguous elements of a single type
   - `[20]int`

#### Slices
  - Describe an underlying array
  - Do not require a specific number of elements
   - `[]int`

#### Functions
  - Can define multiple return values
  - Parameters are always passed by value
  - Anonymous functions are allowed.
  - Variadic functions are possible
   - `func(a string, b ...int)` Any number of integers can be passed
     as the final parameter.

#### Structs
  - Similar to classes in C++.
  - Anonymous structs are allowed.
  - Public/Protected/Private keywords are not used for visibility as in C++

#### Interfaces
  - Specifies a set of methods that a class must define
  - `interface{}` matches all types since all types implement
    0 or more methods.

#### Maps
  - Unordered key:value grouping

#### Channels
  - Mechanism for communication between functions executing concurrently.
  - Defines a direction of information and the type of messages being sent
```
chan   int    // Integers can be sent or received
chan<- int    // Integers can only be sent
<-chan int    // Integers can only be received
```	


#### References
https://golang.org/ref/spec
https://blog.golang.org/strings
