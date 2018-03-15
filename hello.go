package main

import "fmt"

const helloPrefix = "Hello, "

// Hello will print out "Hello, World"
func Hello(name string) string {
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("Guy"))

}
