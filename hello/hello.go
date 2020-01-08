package main

import "fmt"

func hello(name string) string {
	return "Hello GOrld, " + name
}

func main() {
	fmt.Println(hello("Rian"))
}
