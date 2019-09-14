package main

import "fmt"

func Hello(name string) string {
	if name != "" {
		return "Hello, " + name
	} else {
		return "Hello, World"
	}
}

func main() {
	fmt.Println(Hello("world"))
}
