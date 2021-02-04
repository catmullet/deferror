# 🦴 Deferror

## Description

Simple way to automatically return errors from functions you defer.

## Examples

### Style 1
```go
// returns "error coming from defer"
package main

import (
	"fmt"
	"github.com/catmullet/deferror"
)

func main() {
    if err := yourLogicHere(); err != nil {
		fmt.Println(err)
	}
}

// defer function that return error
func deferWithError() error {
    return fmt.Errorf("error coming from defer")
}

// function calling defer requires a named variable of interface error
func yourLogicHere() (err error) {
	// pass by pointer the error.
    defer deferror.As(deferWithError, &err)
    return
}
```

### Style 2
```go
// returns "error coming from defer"
package main

import (
	"fmt"
	"github.com/catmullet/deferror"
)

func main() {
    err, deferError := yourLogicHere()
    if err != nil {
		fmt.Println(err)
    }
    if deferError != nil {
		fmt.Println(deferError)
    }
}

// defer function that return error
func deferWithError() error {
    return fmt.Errorf("error coming from defer")
}

// function calling defer requires a named variable of interface error
// you can seperate out the error if you would like. deferror.As will never overwrite the error
// passed in if it is not nil.
func yourLogicHere() (err, deferError error) {
	// pass by pointer the error.
    defer deferror.As(deferWithError, &deferError)
    return
}
```
