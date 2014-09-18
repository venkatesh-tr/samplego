package main

import (
	"fmt"
	"github.com/venkatesh-tr/samplego/inheritdemo"
	"github.com/venkatesh-tr/samplego/ping"
)

func main() {
	fmt.Println("Hello!")
	str := ping.Ping("Hi there!")
	fmt.Println("Response from Ping Server := ", str)
	fmt.Println("Calling Base Function!")
	var base inheritdemo.Baser
	str1, _ := base.DoIt("Hello")
	fmt.Println(" Response from Base : = ", str1)
	var derived inheritdemo.Derived
	str2, _ := derived.DoIt("Hello")
	fmt.Println(" Response from Derived : = ", str2)
}
