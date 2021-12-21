package main

import "fmt"

func main() {
	fmt.Printf(Hello("Davide"))
}

func Hello(name string) string {
	return "Hello, " + name
}
