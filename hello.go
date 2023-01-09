package main

import "fmt"

const portuguese = "pt"
const spanish = "es"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case portuguese:
		prefix = "Olá "
	case spanish:
		prefix = "Hola "
	default:
		prefix = "Hello "
	}

	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
