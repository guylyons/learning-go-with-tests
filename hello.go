package main

import "fmt"

const helloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "

// Hello will print out "Hello, World"
func Hello(name string, language string) string {
	if name == "" {
		return helloPrefix + "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == french {
		return frenchHelloPrefix + name
	}

	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("Guy", ""))
}
