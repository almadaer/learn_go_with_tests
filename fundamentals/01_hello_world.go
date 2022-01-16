package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

type language int

const (
	english language = iota
	spanish
	french
)

func main() {
	fmt.Println(hello("Sam", english))
}

func hello(name string, l language) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix
	switch l {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}
	return prefix + name
}
