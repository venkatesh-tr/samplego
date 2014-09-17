package ping

import (
	"fmt"
)

// Ping package is a sample package to demo the capabilities of a package in Golang

// Ping method just takes a single parameter and returns the same value back to the client
func Ping(str string) string {
	fmt.Println("Call to Ping Service : ", str)
	return str
}
