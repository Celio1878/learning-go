package main

import "fmt"

func main() {
	fmt.Println(Hello(""))
}

const prefix = "Hello "

func Hello(name string) string {
	if name == "" {
		name = "Anybody"
	}

	return prefix + name
}
