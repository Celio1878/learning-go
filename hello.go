package main

import "fmt"

func main() {
	fmt.Println(Hello("", "brazilian"))
}

func Hello(name string, language string) string {

	if name == "" {
		name = "Anybody"
	}

	return salution_prefix(language) + name
}

func salution_prefix(language string) (prefix string) {

	switch language {
	case "brazilian":
		prefix = "Ola "

	case "france":
		prefix = "Bonjour "

	default:
		prefix = "Hello "
	}

	return
}
