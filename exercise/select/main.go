package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/**
	gets an int64, converts it to a type time.Duration, then uses it in time.Sleep.
	int63n returns an int64
	type Duration's underlying type is int64.
	time.Sleep takes type Duration
	time.MIllisecond is a constant in the time Package
	https://pkg.go.dev/time#pkg-constants
	**/

	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	// setting up a channel to hold values of type integers
	channel1 := make(chan int)
	channel2 := make(chan int)

	go func() {
		time.Sleep(d1 * time.Millisecond)
		channel1 <- 34
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		channel2 <- 20
	}()

	// the select case... Similar to switch cases but specifically for concurrency pattterns
	select {
	case v1 := <-channel1:
		fmt.Printf("v1: %d \n", v1)

	case v2 := <-channel2:
		fmt.Printf("v2: %d \n", v2)
	}

}
