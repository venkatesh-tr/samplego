package main

import (
	"fmt"
	"github.com/venkatesh-tr/samplego/ping"
)

func main() {
	fmt.Println("Hello!")
	str := ping.Ping("Hi there!")
	fmt.Println("Response from Ping Server := ", str)
}
