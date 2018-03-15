package main

import "fmt"

// Hello will print out "Hello, World"
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("Guy"))

}
