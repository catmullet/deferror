package main

import (
	"fmt"
	"github.com/catmullet/deferror"
)

// returns "error coming from defer"
func main() {
	if err := yourLogicHere(); err != nil {
		fmt.Println(err)
	}
}

func deferWithError() error {
	return fmt.Errorf("error coming from defer")
}

// requires a named variable of interface error
func yourLogicHere() (err error) {
	defer deferror.As(deferWithError, &err)
	return
}
