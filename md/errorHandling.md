Error Handling
---
#### Living Without Exceptions
  - Go does not have exceptions. Uses the error interface.
  - Because error is an interface, creating custom error types is easy, with as much or as little detail as needed. 
  - Because functions can return multiple values, at least one of the returned values are for errors.
  - Must do something with the error returned, using a variable or an underscore to intentionally ignore the error.
  - For example: `val, err = sqrt(-2)` would return a default value to val, and an error to err.
  - Without an error, err will be equal to nil. So we would check for errors by comparing the returned err to nil.
``` 
  val, err := sqrt(-2) 
  if err != nil {
       // do error handling
  }
```
  - Can create custom error messages calling errors New method. 
```
  func sqrt(f float64) (float64, err){
    if f < 0 {
      return 0, errors.New("Input value is below zero.")
    }
  }
```
 - Some users may be frustrated by constantly error checking with Go, but that just means you must in some way handle it which requires thought and leads to solid, stable code.
#### Handle It... Later
  - The defer statement pushes a function call onto a deferred stack. The list of saved calls is then executed at the end of the surrounding function.

  - Commonly this is used to simplify function clean up.

  - Deffered functions can read and assign to the functions return values.

#### Panic! 
  - Built in function stops the flow of the program and ... Panics.

  - If a function calls F, F stops, any deferred calls run, and the process goes up the call stack until all routines have returned and the program crashes.

#### Aaand Recover.

  - Recovers from a panicking function.
  
  - Is only useful in deferred functions. 

  - Returns nil in non-panicking situations. 

